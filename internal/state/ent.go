package state

import (
	"context"

	"github.com/go-faster/errors"
	"github.com/google/go-github/v50/github"

	"github.com/go-faster/bot/internal/ent"
	"github.com/go-faster/bot/internal/ent/lastchannelmessage"
	"github.com/go-faster/bot/internal/ent/prnotification"
)

type Ent struct {
	db *ent.Client
}

func (e Ent) UpdateLastMsgID(ctx context.Context, channelID int64, msgID int) error {
	if err := e.db.LastChannelMessage.Create().
		SetID(channelID).
		SetMessageID(msgID).
		OnConflict().UpdateMessageID().
		Exec(ctx); err != nil {
		return errors.Wrap(err, "upsert")
	}
	return nil
}

func (e Ent) SetPRNotification(ctx context.Context, pr *github.PullRequestEvent, msgID int) error {
	if err := e.db.PRNotification.Create().
		SetPullRequestID(pr.GetPullRequest().GetNumber()).
		SetRepoID(pr.GetRepo().GetID()).
		SetMessageID(msgID).
		OnConflict().UpdateNewValues().Exec(ctx); err != nil {
		return errors.Wrap(err, "upsert")
	}
	return nil
}

func (e Ent) FindPRNotification(ctx context.Context, channelID int64, pr *github.PullRequestEvent) (msgID, lastMsgID int, rerr error) {
	tx, err := e.db.Tx(ctx)
	if err != nil {
		return 0, 0, errors.Wrap(err, "begin tx")
	}
	defer func() {
		_ = tx.Rollback()
	}()

	{
		list, err := tx.LastChannelMessage.Query().Where(
			lastchannelmessage.IDEQ(channelID),
		).All(ctx)
		if err != nil {
			return 0, 0, errors.Wrap(err, "query last message")
		}
		for _, v := range list {
			lastMsgID = v.MessageID
		}
	}
	{
		list, err := tx.PRNotification.Query().Where(
			prnotification.PullRequestIDEQ(pr.GetPullRequest().GetNumber()),
			prnotification.RepoIDEQ(pr.GetRepo().GetID()),
		).All(ctx)
		if err != nil {
			return 0, 0, errors.Wrap(err, "query last message")
		}
		for _, v := range list {
			msgID = v.MessageID
		}
	}
	if err := tx.Commit(); err != nil {
		return 0, 0, errors.Wrap(err, "commit")
	}

	return msgID, lastMsgID, nil
}

func NewEnt(db *ent.Client) *Ent {
	return &Ent{db: db}
}

var _ Storage = (*Ent)(nil)
