// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-faster/bot/internal/ent/check"
	"github.com/go-faster/bot/internal/ent/predicate"
)

// CheckUpdate is the builder for updating Check entities.
type CheckUpdate struct {
	config
	hooks    []Hook
	mutation *CheckMutation
}

// Where appends a list predicates to the CheckUpdate builder.
func (cu *CheckUpdate) Where(ps ...predicate.Check) *CheckUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetRepoID sets the "repo_id" field.
func (cu *CheckUpdate) SetRepoID(i int64) *CheckUpdate {
	cu.mutation.ResetRepoID()
	cu.mutation.SetRepoID(i)
	return cu
}

// AddRepoID adds i to the "repo_id" field.
func (cu *CheckUpdate) AddRepoID(i int64) *CheckUpdate {
	cu.mutation.AddRepoID(i)
	return cu
}

// SetPullRequestID sets the "pull_request_id" field.
func (cu *CheckUpdate) SetPullRequestID(i int) *CheckUpdate {
	cu.mutation.ResetPullRequestID()
	cu.mutation.SetPullRequestID(i)
	return cu
}

// AddPullRequestID adds i to the "pull_request_id" field.
func (cu *CheckUpdate) AddPullRequestID(i int) *CheckUpdate {
	cu.mutation.AddPullRequestID(i)
	return cu
}

// SetName sets the "name" field.
func (cu *CheckUpdate) SetName(s string) *CheckUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetStatus sets the "status" field.
func (cu *CheckUpdate) SetStatus(s string) *CheckUpdate {
	cu.mutation.SetStatus(s)
	return cu
}

// SetConclusion sets the "conclusion" field.
func (cu *CheckUpdate) SetConclusion(s string) *CheckUpdate {
	cu.mutation.SetConclusion(s)
	return cu
}

// SetNillableConclusion sets the "conclusion" field if the given value is not nil.
func (cu *CheckUpdate) SetNillableConclusion(s *string) *CheckUpdate {
	if s != nil {
		cu.SetConclusion(*s)
	}
	return cu
}

// ClearConclusion clears the value of the "conclusion" field.
func (cu *CheckUpdate) ClearConclusion() *CheckUpdate {
	cu.mutation.ClearConclusion()
	return cu
}

// Mutation returns the CheckMutation object of the builder.
func (cu *CheckUpdate) Mutation() *CheckMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CheckUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CheckUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CheckUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CheckUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CheckUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(check.Table, check.Columns, sqlgraph.NewFieldSpec(check.FieldID, field.TypeInt64))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.RepoID(); ok {
		_spec.SetField(check.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedRepoID(); ok {
		_spec.AddField(check.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.PullRequestID(); ok {
		_spec.SetField(check.FieldPullRequestID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedPullRequestID(); ok {
		_spec.AddField(check.FieldPullRequestID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(check.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(check.FieldStatus, field.TypeString, value)
	}
	if value, ok := cu.mutation.Conclusion(); ok {
		_spec.SetField(check.FieldConclusion, field.TypeString, value)
	}
	if cu.mutation.ConclusionCleared() {
		_spec.ClearField(check.FieldConclusion, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{check.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CheckUpdateOne is the builder for updating a single Check entity.
type CheckUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CheckMutation
}

// SetRepoID sets the "repo_id" field.
func (cuo *CheckUpdateOne) SetRepoID(i int64) *CheckUpdateOne {
	cuo.mutation.ResetRepoID()
	cuo.mutation.SetRepoID(i)
	return cuo
}

// AddRepoID adds i to the "repo_id" field.
func (cuo *CheckUpdateOne) AddRepoID(i int64) *CheckUpdateOne {
	cuo.mutation.AddRepoID(i)
	return cuo
}

// SetPullRequestID sets the "pull_request_id" field.
func (cuo *CheckUpdateOne) SetPullRequestID(i int) *CheckUpdateOne {
	cuo.mutation.ResetPullRequestID()
	cuo.mutation.SetPullRequestID(i)
	return cuo
}

// AddPullRequestID adds i to the "pull_request_id" field.
func (cuo *CheckUpdateOne) AddPullRequestID(i int) *CheckUpdateOne {
	cuo.mutation.AddPullRequestID(i)
	return cuo
}

// SetName sets the "name" field.
func (cuo *CheckUpdateOne) SetName(s string) *CheckUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *CheckUpdateOne) SetStatus(s string) *CheckUpdateOne {
	cuo.mutation.SetStatus(s)
	return cuo
}

// SetConclusion sets the "conclusion" field.
func (cuo *CheckUpdateOne) SetConclusion(s string) *CheckUpdateOne {
	cuo.mutation.SetConclusion(s)
	return cuo
}

// SetNillableConclusion sets the "conclusion" field if the given value is not nil.
func (cuo *CheckUpdateOne) SetNillableConclusion(s *string) *CheckUpdateOne {
	if s != nil {
		cuo.SetConclusion(*s)
	}
	return cuo
}

// ClearConclusion clears the value of the "conclusion" field.
func (cuo *CheckUpdateOne) ClearConclusion() *CheckUpdateOne {
	cuo.mutation.ClearConclusion()
	return cuo
}

// Mutation returns the CheckMutation object of the builder.
func (cuo *CheckUpdateOne) Mutation() *CheckMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CheckUpdate builder.
func (cuo *CheckUpdateOne) Where(ps ...predicate.Check) *CheckUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CheckUpdateOne) Select(field string, fields ...string) *CheckUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Check entity.
func (cuo *CheckUpdateOne) Save(ctx context.Context) (*Check, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CheckUpdateOne) SaveX(ctx context.Context) *Check {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CheckUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CheckUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CheckUpdateOne) sqlSave(ctx context.Context) (_node *Check, err error) {
	_spec := sqlgraph.NewUpdateSpec(check.Table, check.Columns, sqlgraph.NewFieldSpec(check.FieldID, field.TypeInt64))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Check.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, check.FieldID)
		for _, f := range fields {
			if !check.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != check.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.RepoID(); ok {
		_spec.SetField(check.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedRepoID(); ok {
		_spec.AddField(check.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.PullRequestID(); ok {
		_spec.SetField(check.FieldPullRequestID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedPullRequestID(); ok {
		_spec.AddField(check.FieldPullRequestID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(check.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(check.FieldStatus, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Conclusion(); ok {
		_spec.SetField(check.FieldConclusion, field.TypeString, value)
	}
	if cuo.mutation.ConclusionCleared() {
		_spec.ClearField(check.FieldConclusion, field.TypeString)
	}
	_node = &Check{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{check.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
