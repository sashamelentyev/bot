// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-faster/bot/internal/ent/organization"
	"github.com/go-faster/bot/internal/ent/predicate"
	"github.com/go-faster/bot/internal/ent/repository"
)

// RepositoryUpdate is the builder for updating Repository entities.
type RepositoryUpdate struct {
	config
	hooks    []Hook
	mutation *RepositoryMutation
}

// Where appends a list predicates to the RepositoryUpdate builder.
func (ru *RepositoryUpdate) Where(ps ...predicate.Repository) *RepositoryUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RepositoryUpdate) SetName(s string) *RepositoryUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetFullName sets the "full_name" field.
func (ru *RepositoryUpdate) SetFullName(s string) *RepositoryUpdate {
	ru.mutation.SetFullName(s)
	return ru
}

// SetHTMLURL sets the "html_url" field.
func (ru *RepositoryUpdate) SetHTMLURL(s string) *RepositoryUpdate {
	ru.mutation.SetHTMLURL(s)
	return ru
}

// SetNillableHTMLURL sets the "html_url" field if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableHTMLURL(s *string) *RepositoryUpdate {
	if s != nil {
		ru.SetHTMLURL(*s)
	}
	return ru
}

// ClearHTMLURL clears the value of the "html_url" field.
func (ru *RepositoryUpdate) ClearHTMLURL() *RepositoryUpdate {
	ru.mutation.ClearHTMLURL()
	return ru
}

// SetDescription sets the "description" field.
func (ru *RepositoryUpdate) SetDescription(s string) *RepositoryUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableDescription(s *string) *RepositoryUpdate {
	if s != nil {
		ru.SetDescription(*s)
	}
	return ru
}

// SetLastPushedAt sets the "last_pushed_at" field.
func (ru *RepositoryUpdate) SetLastPushedAt(t time.Time) *RepositoryUpdate {
	ru.mutation.SetLastPushedAt(t)
	return ru
}

// SetNillableLastPushedAt sets the "last_pushed_at" field if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableLastPushedAt(t *time.Time) *RepositoryUpdate {
	if t != nil {
		ru.SetLastPushedAt(*t)
	}
	return ru
}

// ClearLastPushedAt clears the value of the "last_pushed_at" field.
func (ru *RepositoryUpdate) ClearLastPushedAt() *RepositoryUpdate {
	ru.mutation.ClearLastPushedAt()
	return ru
}

// SetLastEventAt sets the "last_event_at" field.
func (ru *RepositoryUpdate) SetLastEventAt(t time.Time) *RepositoryUpdate {
	ru.mutation.SetLastEventAt(t)
	return ru
}

// SetNillableLastEventAt sets the "last_event_at" field if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableLastEventAt(t *time.Time) *RepositoryUpdate {
	if t != nil {
		ru.SetLastEventAt(*t)
	}
	return ru
}

// ClearLastEventAt clears the value of the "last_event_at" field.
func (ru *RepositoryUpdate) ClearLastEventAt() *RepositoryUpdate {
	ru.mutation.ClearLastEventAt()
	return ru
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (ru *RepositoryUpdate) SetOrganizationID(id int64) *RepositoryUpdate {
	ru.mutation.SetOrganizationID(id)
	return ru
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableOrganizationID(id *int64) *RepositoryUpdate {
	if id != nil {
		ru = ru.SetOrganizationID(*id)
	}
	return ru
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (ru *RepositoryUpdate) SetOrganization(o *Organization) *RepositoryUpdate {
	return ru.SetOrganizationID(o.ID)
}

// Mutation returns the RepositoryMutation object of the builder.
func (ru *RepositoryUpdate) Mutation() *RepositoryMutation {
	return ru.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (ru *RepositoryUpdate) ClearOrganization() *RepositoryUpdate {
	ru.mutation.ClearOrganization()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RepositoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, RepositoryMutation](ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RepositoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RepositoryUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RepositoryUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RepositoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(repository.Table, repository.Columns, sqlgraph.NewFieldSpec(repository.FieldID, field.TypeInt64))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(repository.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.FullName(); ok {
		_spec.SetField(repository.FieldFullName, field.TypeString, value)
	}
	if value, ok := ru.mutation.HTMLURL(); ok {
		_spec.SetField(repository.FieldHTMLURL, field.TypeString, value)
	}
	if ru.mutation.HTMLURLCleared() {
		_spec.ClearField(repository.FieldHTMLURL, field.TypeString)
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.SetField(repository.FieldDescription, field.TypeString, value)
	}
	if value, ok := ru.mutation.LastPushedAt(); ok {
		_spec.SetField(repository.FieldLastPushedAt, field.TypeTime, value)
	}
	if ru.mutation.LastPushedAtCleared() {
		_spec.ClearField(repository.FieldLastPushedAt, field.TypeTime)
	}
	if value, ok := ru.mutation.LastEventAt(); ok {
		_spec.SetField(repository.FieldLastEventAt, field.TypeTime, value)
	}
	if ru.mutation.LastEventAtCleared() {
		_spec.ClearField(repository.FieldLastEventAt, field.TypeTime)
	}
	if ru.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.OrganizationTable,
			Columns: []string{repository.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.OrganizationTable,
			Columns: []string{repository.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repository.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RepositoryUpdateOne is the builder for updating a single Repository entity.
type RepositoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RepositoryMutation
}

// SetName sets the "name" field.
func (ruo *RepositoryUpdateOne) SetName(s string) *RepositoryUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetFullName sets the "full_name" field.
func (ruo *RepositoryUpdateOne) SetFullName(s string) *RepositoryUpdateOne {
	ruo.mutation.SetFullName(s)
	return ruo
}

// SetHTMLURL sets the "html_url" field.
func (ruo *RepositoryUpdateOne) SetHTMLURL(s string) *RepositoryUpdateOne {
	ruo.mutation.SetHTMLURL(s)
	return ruo
}

// SetNillableHTMLURL sets the "html_url" field if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableHTMLURL(s *string) *RepositoryUpdateOne {
	if s != nil {
		ruo.SetHTMLURL(*s)
	}
	return ruo
}

// ClearHTMLURL clears the value of the "html_url" field.
func (ruo *RepositoryUpdateOne) ClearHTMLURL() *RepositoryUpdateOne {
	ruo.mutation.ClearHTMLURL()
	return ruo
}

// SetDescription sets the "description" field.
func (ruo *RepositoryUpdateOne) SetDescription(s string) *RepositoryUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableDescription(s *string) *RepositoryUpdateOne {
	if s != nil {
		ruo.SetDescription(*s)
	}
	return ruo
}

// SetLastPushedAt sets the "last_pushed_at" field.
func (ruo *RepositoryUpdateOne) SetLastPushedAt(t time.Time) *RepositoryUpdateOne {
	ruo.mutation.SetLastPushedAt(t)
	return ruo
}

// SetNillableLastPushedAt sets the "last_pushed_at" field if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableLastPushedAt(t *time.Time) *RepositoryUpdateOne {
	if t != nil {
		ruo.SetLastPushedAt(*t)
	}
	return ruo
}

// ClearLastPushedAt clears the value of the "last_pushed_at" field.
func (ruo *RepositoryUpdateOne) ClearLastPushedAt() *RepositoryUpdateOne {
	ruo.mutation.ClearLastPushedAt()
	return ruo
}

// SetLastEventAt sets the "last_event_at" field.
func (ruo *RepositoryUpdateOne) SetLastEventAt(t time.Time) *RepositoryUpdateOne {
	ruo.mutation.SetLastEventAt(t)
	return ruo
}

// SetNillableLastEventAt sets the "last_event_at" field if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableLastEventAt(t *time.Time) *RepositoryUpdateOne {
	if t != nil {
		ruo.SetLastEventAt(*t)
	}
	return ruo
}

// ClearLastEventAt clears the value of the "last_event_at" field.
func (ruo *RepositoryUpdateOne) ClearLastEventAt() *RepositoryUpdateOne {
	ruo.mutation.ClearLastEventAt()
	return ruo
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (ruo *RepositoryUpdateOne) SetOrganizationID(id int64) *RepositoryUpdateOne {
	ruo.mutation.SetOrganizationID(id)
	return ruo
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableOrganizationID(id *int64) *RepositoryUpdateOne {
	if id != nil {
		ruo = ruo.SetOrganizationID(*id)
	}
	return ruo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (ruo *RepositoryUpdateOne) SetOrganization(o *Organization) *RepositoryUpdateOne {
	return ruo.SetOrganizationID(o.ID)
}

// Mutation returns the RepositoryMutation object of the builder.
func (ruo *RepositoryUpdateOne) Mutation() *RepositoryMutation {
	return ruo.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (ruo *RepositoryUpdateOne) ClearOrganization() *RepositoryUpdateOne {
	ruo.mutation.ClearOrganization()
	return ruo
}

// Where appends a list predicates to the RepositoryUpdate builder.
func (ruo *RepositoryUpdateOne) Where(ps ...predicate.Repository) *RepositoryUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RepositoryUpdateOne) Select(field string, fields ...string) *RepositoryUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Repository entity.
func (ruo *RepositoryUpdateOne) Save(ctx context.Context) (*Repository, error) {
	return withHooks[*Repository, RepositoryMutation](ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RepositoryUpdateOne) SaveX(ctx context.Context) *Repository {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RepositoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RepositoryUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RepositoryUpdateOne) sqlSave(ctx context.Context) (_node *Repository, err error) {
	_spec := sqlgraph.NewUpdateSpec(repository.Table, repository.Columns, sqlgraph.NewFieldSpec(repository.FieldID, field.TypeInt64))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Repository.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, repository.FieldID)
		for _, f := range fields {
			if !repository.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != repository.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(repository.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.FullName(); ok {
		_spec.SetField(repository.FieldFullName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.HTMLURL(); ok {
		_spec.SetField(repository.FieldHTMLURL, field.TypeString, value)
	}
	if ruo.mutation.HTMLURLCleared() {
		_spec.ClearField(repository.FieldHTMLURL, field.TypeString)
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.SetField(repository.FieldDescription, field.TypeString, value)
	}
	if value, ok := ruo.mutation.LastPushedAt(); ok {
		_spec.SetField(repository.FieldLastPushedAt, field.TypeTime, value)
	}
	if ruo.mutation.LastPushedAtCleared() {
		_spec.ClearField(repository.FieldLastPushedAt, field.TypeTime)
	}
	if value, ok := ruo.mutation.LastEventAt(); ok {
		_spec.SetField(repository.FieldLastEventAt, field.TypeTime, value)
	}
	if ruo.mutation.LastEventAtCleared() {
		_spec.ClearField(repository.FieldLastEventAt, field.TypeTime)
	}
	if ruo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.OrganizationTable,
			Columns: []string{repository.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.OrganizationTable,
			Columns: []string{repository.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Repository{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repository.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
