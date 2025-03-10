package gh

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/go-github/v52/github"
	"go.opentelemetry.io/otel/trace"
)

type Event struct {
	Type     string
	RepoName string
	RepoID   int64
}

func htmlURL(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		return s
	}
	u.Host = "github.com"
	secondSlash := strings.Index(u.Path[1:], "/")
	if secondSlash == -1 {
		return s
	}
	u.Path = u.Path[secondSlash+1:]
	return u.String()
}

func (h *Webhook) GetRepoByID(ctx context.Context, id int64) (*github.Repository, error) {
	ctx, span := h.tracer.Start(ctx, "wh.GetRepoByID",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer span.End()

	client, err := h.Client(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "client")
	}

	key := fmt.Sprintf("github:repo:%d", id)
	if raw, err := h.cache.Get(ctx, key).Result(); err == nil {
		var repo github.Repository
		if err := json.Unmarshal([]byte(raw), &repo); err != nil {
			return nil, errors.Wrap(err, "unmarshal")
		}
		return &repo, nil
	}

	repo, _, err := client.Repositories.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "get by id")
	}

	raw, err := json.Marshal(repo)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}
	if _, err := h.cache.Set(ctx, key, string(raw), time.Hour).Result(); err != nil {
		return nil, errors.Wrap(err, "set")
	}

	return repo, nil
}

func (h *Webhook) Transform(ctx context.Context, d *jx.Decoder, e *jx.Encoder) (*Event, error) {
	ctx, span := h.tracer.Start(ctx, "wh.Transform",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer span.End()

	var (
		repoID       int64
		fullRepoName string
		repoURL      string
		evType       string
	)
	e.ObjStart()
	if err := d.ObjBytes(func(d *jx.Decoder, key []byte) error {
		var err error
		switch string(key) {
		case "actor":
			e.FieldStart("sender")
			e.ObjStart()
			if err := d.ObjBytes(func(d *jx.Decoder, key []byte) error {
				switch string(key) {
				case "url":
					s, err := d.Str()
					if err != nil {
						return errors.Wrap(err, "url")
					}
					e.Field("url", func(e *jx.Encoder) {
						e.Str(s)
					})
					e.Field("html_url", func(e *jx.Encoder) {
						e.Str(htmlURL(s))
					})
					return nil
				default:
					v, err := d.Raw()
					if err != nil {
						return errors.Wrap(err, "actor")
					}
					e.Field(string(key), func(e *jx.Encoder) {
						e.Raw(v)
					})
					return nil
				}
			}); err != nil {
				return err
			}
			e.ObjEnd()
			return nil
		case "payload":
			return d.ObjBytes(func(d *jx.Decoder, key []byte) error {
				v, err := d.Raw()
				if err != nil {
					return errors.Wrap(err, "payload")
				}
				e.Field(string(key), func(e *jx.Encoder) {
					e.Raw(v)
				})
				return nil
			})
		case "type":
			if evType, err = d.Str(); err != nil {
				return errors.Wrap(err, "type")
			}
			return nil
		case "repo":
			return d.ObjBytes(func(d *jx.Decoder, key []byte) error {
				switch string(key) {
				case "id":
					if repoID, err = d.Int64(); err != nil {
						return errors.Wrap(err, "id")
					}
					return nil
				case "name":
					if fullRepoName, err = d.Str(); err != nil {
						return errors.Wrap(err, "name")
					}
					return nil
				case "url":
					if repoURL, err = d.Str(); err != nil {
						return errors.Wrap(err, "url")
					}
					return nil
				default:
					return d.Skip()
				}
			})
		default:
			return d.Skip()
		}
	}); err != nil {
		return nil, errors.Wrap(err, "decode")
	}

	var (
		repoName  string
		ownerName string
		ownerID   int64
	)

	// Fetch owner.
	if h.cache != nil {
		repo, err := h.GetRepoByID(ctx, repoID)
		if err != nil {
			return nil, errors.Wrap(err, "get repo by id")
		}
		ownerID = repo.GetOwner().GetID()
		ownerName = repo.GetOwner().GetLogin()
		repoName = repo.GetName()
	}
	if ownerName == "" {
		owner, name, ok := strings.Cut(fullRepoName, "/")
		if ok {
			ownerName = owner
			repoName = name
		}
	}
	e.Field("repository", func(e *jx.Encoder) {
		e.Obj(func(e *jx.Encoder) {
			e.Field("id", func(e *jx.Encoder) {
				e.Int64(repoID)
			})
			e.Field("full_name", func(e *jx.Encoder) {
				e.Str(fullRepoName)
			})
			e.Field("url", func(e *jx.Encoder) {
				e.Str(repoURL)
			})
			e.Field("html_url", func(e *jx.Encoder) {
				e.Str(htmlURL(repoURL))
			})
			if repoName != "" {
				e.Field("name", func(e *jx.Encoder) {
					e.Str(repoName)
				})
			}
			if ownerName != "" {
				e.Field("owner", func(e *jx.Encoder) {
					e.Obj(func(e *jx.Encoder) {
						e.Field("login", func(e *jx.Encoder) {
							e.Str(ownerName)
						})
						if ownerID != 0 {
							e.Field("id", func(e *jx.Encoder) {
								e.Int64(ownerID)
							})
						}
					})
				})
			}
		})
	})
	e.ObjEnd()
	d.ResetBytes(e.Bytes())

	return &Event{
		Type:     evType,
		RepoName: fullRepoName,
		RepoID:   repoID,
	}, nil
}
