package gh

import (
	"context"

	"github.com/go-faster/errors"
	"github.com/go-faster/sdk/zctx"
	"github.com/google/go-github/v50/github"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func (h *Webhook) handleCheckRun(ctx context.Context, e *github.CheckRunEvent) error {
	_, span := h.tracer.Start(ctx, "handleCheckRun",
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	span.AddEvent("CheckRunEvent",
		trace.WithStackTrace(true),
		trace.WithAttributes(
			attribute.String("action", e.GetAction()),
			attribute.String("check_run.name", e.GetCheckRun().GetName()),
			attribute.String("check_run.status", e.GetCheckRun().GetStatus()),
			attribute.String("check_run.conclusion", e.GetCheckRun().GetConclusion()),
			attribute.String("check_run.head_sha", e.GetCheckRun().GetHeadSHA()),

			attribute.Int64("organization.id", e.GetOrg().GetID()),
			attribute.String("organization.login", e.GetOrg().GetLogin()),
			attribute.String("repository.full_name", e.GetRepo().GetFullName()),
			attribute.Int64("repository.id", e.GetRepo().GetID()),
		),
	)

	ctx = zctx.With(ctx,
		zap.String("action", e.GetAction()),
		zap.Int64("run_id", e.GetCheckRun().GetID()),
		zap.String("head_sha", e.GetCheckRun().GetHeadSHA()),
	)
	lg := zctx.From(ctx)

	pr, err := h.upsertCheck(ctx, e)
	if err != nil {
		return errors.Wrap(err, "upsert check")
	}
	if pr == nil {
		// No PR - no update.
		lg.Debug("Ignore event: no PR info")
		return nil
	}
	lg.Debug("Emit check update",
		zap.String("pr_head_sha", pr.GetHead().GetSHA()),
	)

	h.updater.Emit(PullRequestUpdate{
		Event:  "check_update",
		Action: "",
		Repo:   e.GetRepo(),
		PR:     pr,
		Checks: nil,
	})
	return nil
}
