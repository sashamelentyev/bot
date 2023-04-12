package gh

import (
	"context"

	"github.com/google/go-github/v50/github"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func (h Webhook) handleCheckRun(ctx context.Context, e *github.CheckRunEvent) error {
	ctx, span := h.tracer.Start(ctx, "handleCheckRun",
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

	return nil
}
