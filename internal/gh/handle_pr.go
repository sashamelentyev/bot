package gh

import (
	"context"
	"fmt"
	"net/url"
	"path"

	"github.com/go-faster/errors"
	"github.com/go-faster/sdk/zctx"
	"github.com/google/go-github/v52/github"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/gotd/td/telegram/message/entity"
	"github.com/gotd/td/telegram/message/markup"
	"github.com/gotd/td/telegram/message/styling"
	"github.com/gotd/td/telegram/message/unpack"
	"github.com/gotd/td/tg"

	"github.com/go-faster/bot/internal/action"
)

type PullRequestUpdate struct {
	// Possible values: pr_update, check_update
	Event string
	// Possible values for Event == "update": opened, merged
	Action string

	Repo   *github.Repository
	PR     *github.PullRequest
	Checks []Check
}

func (u PullRequestUpdate) ActionIn(actions ...string) bool {
	for _, a := range actions {
		if u.Action == a {
			return true
		}
	}
	return false
}

func generateChecksStatus(checks []Check) string {
	var (
		failed  int
		pending int
		success int
	)
	for _, check := range checks {
		switch check.Status {
		case "created", "rerequested":
			pending++
		case "completed":
			switch check.Conclusion {
			case "failure", "timed_out", "cancelled":
				failed++
			case "success":
				success++
			case "action_required", "neutral", "skipped":
				// Ignore
			}
		case "requested_action":
			// Ignore
		}
	}

	switch {
	case failed > 0:
		return "Checks❌"
	case pending > 0:
		return "Checks⏳"
	case success > 0:
		return "Checks✅"
	default:
		return "Checks▶"
	}
}

func getPullRequestURL(repo *github.Repository, pr *github.PullRequest) styling.StyledTextOption {
	urlName := fmt.Sprintf("%s#%d",
		repo.GetFullName(),
		pr.GetNumber(),
	)

	return styling.TextURL(urlName, pr.GetHTMLURL())
}

func githubUserLink(u *github.User) styling.StyledTextOption {
	return styling.TextURL(u.GetLogin(), u.GetHTMLURL())
}

func (h *Webhook) updatePR(ctx context.Context, state PullRequestUpdate) error {
	ctx, span := h.tracer.Start(ctx, "UpdatePR",
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	lg := zctx.From(ctx)

	p, err := h.notifyPeer(ctx)
	if err != nil {
		return errors.Wrap(err, "peer")
	}

	ch, ok := p.(*tg.InputPeerChannel)
	if !ok {
		return errors.Errorf("unexpected peer %T", p)
	}

	var (
		repo = state.Repo
		pr   = state.PR
	)
	existingMsgID, lastMsgID, err := h.findPRNotification(ctx, repo, pr, ch.ChannelID)
	if err != nil {
		return errors.Wrap(err, "query message state")
	}

	if existingMsgID == 0 && state.Event == "check_update" {
		// Prevent a possible race between handlers: pr event handler is only who sends.
		lg.Warn("Update pull request check status: there is no notification yet")
		return nil
	}

	// Send a new message if
	//
	// 1) There is no existing message notification
	// 2) There is less than 10 messages between this notification and last messsge.
	// 3) State update is a PR state update, not a check run or something like that.
	//
	gonnaSendNewMessage := (existingMsgID == 0 || lastMsgID-existingMsgID > 10) && state.Event == "pr_update"

	r := h.sender.To(p).NoWebpage()
	// Setup buttons.
	checksStatus := generateChecksStatus(state.Checks)
	if u, _ := url.ParseRequestURI(pr.GetHTMLURL()); u != nil && !state.ActionIn("merged", "closed") {
		files, checks := *u, *u
		files.Path = path.Join(files.Path, "files")
		checks.Path = path.Join(checks.Path, "checks")

		mergeAction := action.Action{
			Type:         action.Merge,
			ID:           pr.GetNumber(),
			RepositoryID: repo.GetID(),
			Entity:       action.PullRequest,
		}
		r = r.Row(
			markup.URL(checksStatus, checks.String()),
			markup.Callback("Merge", action.Marshal(mergeAction)),
		)
	}

	var text []styling.StyledTextOption
	// TODO: handle more states?
	switch state.Action {
	case "", "opened":
		readiness := " opened"
		if pr.GetDraft() {
			readiness = " drafted"
		}

		text = append(text,
			styling.Plain("New pull request "),
			getPullRequestURL(repo, pr),
			styling.Plain(readiness),
			styling.Plain(" by "),
			githubUserLink(pr.GetUser()),
		)

	case "closed":
		text = append(text,
			styling.Plain("Pull request "),
			getPullRequestURL(repo, pr),
			styling.Plain(" "),
		)

		if !gonnaSendNewMessage {
			text = append(text,
				styling.Strike("opened by "),
				styling.Custom(func(eb *entity.Builder) error {
					u := pr.GetUser()
					eb.Format(
						u.GetLogin(),
						entity.Strike(),
						entity.TextURL(u.GetHTMLURL()),
					)
					return nil
				}),
			)
		}

		msg := " closed by "
		merger := pr.GetMergedBy()

		text = append(text,
			styling.Plain(msg),
			githubUserLink(merger),
		)

	case "merged":
		// We can get here only if PR was merged.
		text = append(text,
			styling.Plain("Pull request "),
			getPullRequestURL(repo, pr),
			styling.Plain(" "),
		)

		if !gonnaSendNewMessage {
			text = append(text,
				styling.Strike("opened by "),
				styling.Custom(func(eb *entity.Builder) error {
					u := pr.GetUser()
					eb.Format(
						u.GetLogin(),
						entity.Strike(),
						entity.TextURL(u.GetHTMLURL()),
					)
					return nil
				}),
			)
		}

		msg := " merged by "
		merger := pr.GetMergedBy()
		if am := pr.AutoMerge; am != nil {
			msg = " auto-merged by "
			if u := am.EnabledBy; u != nil {
				merger = u
			}
		}

		text = append(text,
			styling.Plain(msg),
			githubUserLink(merger),
		)
	}

	text = append(text,
		styling.Plain("\n\n"),
		styling.Italic(pr.GetTitle()),
	)

	lg.Debug("Updating PR",
		zap.String("updater.event", state.Event),
		zap.String("updater.action", state.Action),
		zap.Bool("new_message", gonnaSendNewMessage),
		zap.String("checks", checksStatus),
	)
	if !gonnaSendNewMessage {
		if _, err := r.Edit(existingMsgID).StyledText(ctx, text...); err != nil && !tg.IsMessageNotModified(err) {
			return errors.Wrap(err, "edit message")
		}
		return nil
	}

	if existingMsgID > 0 {
		r = r.Reply(existingMsgID)
	}
	newMsgID, err := unpack.MessageID(r.StyledText(ctx, text...))
	if err != nil {
		return errors.Wrap(err, "send message")
	}

	return multierr.Append(
		h.updateLastMsgID(ctx, ch.ChannelID, newMsgID),
		h.setPRNotification(ctx, repo, pr, newMsgID),
	)
}

func (h *Webhook) handlePRClosed(ctx context.Context, e *github.PullRequestEvent) error {
	var (
		repo = e.GetRepo()
		pr   = e.GetPullRequest()
	)
	if !pr.GetMerged() {
		return h.updater.Emit(PullRequestUpdate{
			Event:  "pr_update",
			Action: "closed",
			Repo:   repo,
			PR:     pr,
			Checks: nil,
		})
	}

	return h.updater.Emit(PullRequestUpdate{
		Event:  "pr_update",
		Action: "merged",
		Repo:   repo,
		PR:     pr,
		Checks: nil,
	})
}

func (h *Webhook) handlePROpened(ctx context.Context, e *github.PullRequestEvent) error {
	return h.updater.Emit(PullRequestUpdate{
		Event:  "pr_update",
		Action: "opened",
		Repo:   e.GetRepo(),
		PR:     e.GetPullRequest(),
		Checks: nil,
	})
}

func (h *Webhook) handlePR(ctx context.Context, e *github.PullRequestEvent) error {
	switch e.GetAction() {
	case "opened":
		return h.handlePROpened(ctx, e)
	case "closed":
		return h.handlePRClosed(ctx, e)
	}
	return nil
}
