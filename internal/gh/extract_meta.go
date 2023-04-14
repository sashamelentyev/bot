package gh

import (
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/attribute"
)

type eventMeta struct {
	Organization   string // go-faster
	OrganizationID int64  // 93744681

	Repository   string // bot
	RepositoryID int64  // 512150878
}

func (m eventMeta) Attributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("org.name", m.Organization),
		attribute.Int64("org.id", m.OrganizationID),
		attribute.String("repo.name", m.Repository),
		attribute.Int64("repo.id", m.RepositoryID),
	}
}

func extractEventMeta(raw []byte) (*eventMeta, error) {
	d := jx.GetDecoder()
	defer jx.PutDecoder(d)

	d.ResetBytes(raw)

	var m eventMeta
	parseOrg := func(d *jx.Decoder, key []byte) error {
		switch string(key) {
		case "login":
			v, err := d.Str()
			if err != nil {
				return err
			}
			m.Organization = v
			return nil
		case "id":
			v, err := d.Int64()
			if err != nil {
				return err
			}
			m.OrganizationID = v
			return nil
		default:
			return d.Skip()
		}
	}
	parseRepo := func(d *jx.Decoder, key []byte) error {
		switch string(key) {
		case "name":
			v, err := d.Str()
			if err != nil {
				return err
			}
			m.Repository = v
			return nil
		case "id":
			v, err := d.Int64()
			if err != nil {
				return err
			}
			m.RepositoryID = v
			return nil
		default:
			return d.Skip()
		}
	}
	if err := d.ObjBytes(func(d *jx.Decoder, key []byte) error {
		switch string(key) {
		case "organization":
			return d.ObjBytes(parseOrg)
		case "repository":
			return d.ObjBytes(parseRepo)
		default:
			return d.Skip()
		}
	}); err != nil {
		return nil, err
	}

	return &m, nil
}
