package gh

import (
	"context"
	"sync"
	"time"

	"github.com/go-faster/bot/internal/ent"
	"github.com/go-faster/errors"
	"github.com/go-faster/sdk/zctx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type prKey struct {
	// repo is full repo name.
	repo string
	// number is pr number.
	number int
}

func (k prKey) MarshalLogObject(e zapcore.ObjectEncoder) error {
	e.AddString("repo", k.repo)
	e.AddInt("pr", k.number)
	return nil
}

type queuedUpdate struct {
	Update PullRequestUpdate
	Tries  int
}

type updater struct {
	w *Webhook

	tick       time.Duration
	updates    map[prKey]*queuedUpdate
	updatesMux sync.Mutex
}

func newUpdater(w *Webhook, tick time.Duration) *updater {
	return &updater{
		w:    w,
		tick: tick,
		// TODO(tdakkota): store queue in DB?
		updates: map[prKey]*queuedUpdate{},
	}
}

func (u *updater) updateOne(ctx context.Context, update PullRequestUpdate) error {
	if update.Event == "check_update" {
		switch err := fillPRState(ctx, u.w.db.PRNotification, update.Repo, update.PR); err != nil {
		case err == nil:
		case ent.IsNotFound(err):
			// No message sent yet.
			return nil
		default:
			return errors.Wrap(err, "query cached pr fields")
		}
	}

	// Do not query checks if PR was merged: we won't send status anyway.
	if update.Action != "merged" && update.Checks == nil {
		checks, err := u.w.queryChecks(ctx, update.Repo, update.PR)
		if err != nil {
			return errors.Wrap(err, "query checks")
		}
		update.Checks = checks
	}

	return u.w.updatePR(ctx, update)
}

func (u *updater) doUpdate(ctx context.Context) {
	u.updatesMux.Lock()
	defer u.updatesMux.Unlock()

	for key, qu := range u.updates {
		ctx := zctx.With(ctx, zap.Inline(key))

		if err := u.updateOne(ctx, qu.Update); err != nil {
			zctx.From(ctx).Error("PR Update failed", zap.Error(err))
			if qu.Tries < 5 {
				qu.Tries++
				continue
			}
		}
		delete(u.updates, key)
	}
}

// Emit enqueues update.
func (u *updater) Emit(update PullRequestUpdate) {
	u.updatesMux.Lock()
	defer u.updatesMux.Unlock()

	key := prKey{
		update.Repo.GetFullName(),
		update.PR.GetNumber(),
	}
	u.updates[key] = &queuedUpdate{
		Update: update,
	}
}

// Run setups update worker.
func (u *updater) Run(ctx context.Context) error {
	t := time.NewTicker(u.tick)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			u.doUpdate(ctx)
		}
	}
}
