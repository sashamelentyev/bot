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
	"github.com/go-faster/bot/internal/ent/gitcommit"
	"github.com/go-faster/bot/internal/ent/predicate"
	"github.com/go-faster/bot/internal/ent/repository"
)

// GitCommitUpdate is the builder for updating GitCommit entities.
type GitCommitUpdate struct {
	config
	hooks    []Hook
	mutation *GitCommitMutation
}

// Where appends a list predicates to the GitCommitUpdate builder.
func (gcu *GitCommitUpdate) Where(ps ...predicate.GitCommit) *GitCommitUpdate {
	gcu.mutation.Where(ps...)
	return gcu
}

// SetMessage sets the "message" field.
func (gcu *GitCommitUpdate) SetMessage(s string) *GitCommitUpdate {
	gcu.mutation.SetMessage(s)
	return gcu
}

// SetAuthorLogin sets the "author_login" field.
func (gcu *GitCommitUpdate) SetAuthorLogin(s string) *GitCommitUpdate {
	gcu.mutation.SetAuthorLogin(s)
	return gcu
}

// SetAuthorID sets the "author_id" field.
func (gcu *GitCommitUpdate) SetAuthorID(i int64) *GitCommitUpdate {
	gcu.mutation.ResetAuthorID()
	gcu.mutation.SetAuthorID(i)
	return gcu
}

// AddAuthorID adds i to the "author_id" field.
func (gcu *GitCommitUpdate) AddAuthorID(i int64) *GitCommitUpdate {
	gcu.mutation.AddAuthorID(i)
	return gcu
}

// SetDate sets the "date" field.
func (gcu *GitCommitUpdate) SetDate(t time.Time) *GitCommitUpdate {
	gcu.mutation.SetDate(t)
	return gcu
}

// SetRepositoryID sets the "repository" edge to the Repository entity by ID.
func (gcu *GitCommitUpdate) SetRepositoryID(id int64) *GitCommitUpdate {
	gcu.mutation.SetRepositoryID(id)
	return gcu
}

// SetNillableRepositoryID sets the "repository" edge to the Repository entity by ID if the given value is not nil.
func (gcu *GitCommitUpdate) SetNillableRepositoryID(id *int64) *GitCommitUpdate {
	if id != nil {
		gcu = gcu.SetRepositoryID(*id)
	}
	return gcu
}

// SetRepository sets the "repository" edge to the Repository entity.
func (gcu *GitCommitUpdate) SetRepository(r *Repository) *GitCommitUpdate {
	return gcu.SetRepositoryID(r.ID)
}

// Mutation returns the GitCommitMutation object of the builder.
func (gcu *GitCommitUpdate) Mutation() *GitCommitMutation {
	return gcu.mutation
}

// ClearRepository clears the "repository" edge to the Repository entity.
func (gcu *GitCommitUpdate) ClearRepository() *GitCommitUpdate {
	gcu.mutation.ClearRepository()
	return gcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gcu *GitCommitUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, gcu.sqlSave, gcu.mutation, gcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gcu *GitCommitUpdate) SaveX(ctx context.Context) int {
	affected, err := gcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gcu *GitCommitUpdate) Exec(ctx context.Context) error {
	_, err := gcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcu *GitCommitUpdate) ExecX(ctx context.Context) {
	if err := gcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gcu *GitCommitUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(gitcommit.Table, gitcommit.Columns, sqlgraph.NewFieldSpec(gitcommit.FieldID, field.TypeString))
	if ps := gcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gcu.mutation.Message(); ok {
		_spec.SetField(gitcommit.FieldMessage, field.TypeString, value)
	}
	if value, ok := gcu.mutation.AuthorLogin(); ok {
		_spec.SetField(gitcommit.FieldAuthorLogin, field.TypeString, value)
	}
	if value, ok := gcu.mutation.AuthorID(); ok {
		_spec.SetField(gitcommit.FieldAuthorID, field.TypeInt64, value)
	}
	if value, ok := gcu.mutation.AddedAuthorID(); ok {
		_spec.AddField(gitcommit.FieldAuthorID, field.TypeInt64, value)
	}
	if value, ok := gcu.mutation.Date(); ok {
		_spec.SetField(gitcommit.FieldDate, field.TypeTime, value)
	}
	if gcu.mutation.RepositoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gitcommit.RepositoryTable,
			Columns: []string{gitcommit.RepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(repository.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gcu.mutation.RepositoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gitcommit.RepositoryTable,
			Columns: []string{gitcommit.RepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(repository.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gitcommit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gcu.mutation.done = true
	return n, nil
}

// GitCommitUpdateOne is the builder for updating a single GitCommit entity.
type GitCommitUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GitCommitMutation
}

// SetMessage sets the "message" field.
func (gcuo *GitCommitUpdateOne) SetMessage(s string) *GitCommitUpdateOne {
	gcuo.mutation.SetMessage(s)
	return gcuo
}

// SetAuthorLogin sets the "author_login" field.
func (gcuo *GitCommitUpdateOne) SetAuthorLogin(s string) *GitCommitUpdateOne {
	gcuo.mutation.SetAuthorLogin(s)
	return gcuo
}

// SetAuthorID sets the "author_id" field.
func (gcuo *GitCommitUpdateOne) SetAuthorID(i int64) *GitCommitUpdateOne {
	gcuo.mutation.ResetAuthorID()
	gcuo.mutation.SetAuthorID(i)
	return gcuo
}

// AddAuthorID adds i to the "author_id" field.
func (gcuo *GitCommitUpdateOne) AddAuthorID(i int64) *GitCommitUpdateOne {
	gcuo.mutation.AddAuthorID(i)
	return gcuo
}

// SetDate sets the "date" field.
func (gcuo *GitCommitUpdateOne) SetDate(t time.Time) *GitCommitUpdateOne {
	gcuo.mutation.SetDate(t)
	return gcuo
}

// SetRepositoryID sets the "repository" edge to the Repository entity by ID.
func (gcuo *GitCommitUpdateOne) SetRepositoryID(id int64) *GitCommitUpdateOne {
	gcuo.mutation.SetRepositoryID(id)
	return gcuo
}

// SetNillableRepositoryID sets the "repository" edge to the Repository entity by ID if the given value is not nil.
func (gcuo *GitCommitUpdateOne) SetNillableRepositoryID(id *int64) *GitCommitUpdateOne {
	if id != nil {
		gcuo = gcuo.SetRepositoryID(*id)
	}
	return gcuo
}

// SetRepository sets the "repository" edge to the Repository entity.
func (gcuo *GitCommitUpdateOne) SetRepository(r *Repository) *GitCommitUpdateOne {
	return gcuo.SetRepositoryID(r.ID)
}

// Mutation returns the GitCommitMutation object of the builder.
func (gcuo *GitCommitUpdateOne) Mutation() *GitCommitMutation {
	return gcuo.mutation
}

// ClearRepository clears the "repository" edge to the Repository entity.
func (gcuo *GitCommitUpdateOne) ClearRepository() *GitCommitUpdateOne {
	gcuo.mutation.ClearRepository()
	return gcuo
}

// Where appends a list predicates to the GitCommitUpdate builder.
func (gcuo *GitCommitUpdateOne) Where(ps ...predicate.GitCommit) *GitCommitUpdateOne {
	gcuo.mutation.Where(ps...)
	return gcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gcuo *GitCommitUpdateOne) Select(field string, fields ...string) *GitCommitUpdateOne {
	gcuo.fields = append([]string{field}, fields...)
	return gcuo
}

// Save executes the query and returns the updated GitCommit entity.
func (gcuo *GitCommitUpdateOne) Save(ctx context.Context) (*GitCommit, error) {
	return withHooks(ctx, gcuo.sqlSave, gcuo.mutation, gcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gcuo *GitCommitUpdateOne) SaveX(ctx context.Context) *GitCommit {
	node, err := gcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gcuo *GitCommitUpdateOne) Exec(ctx context.Context) error {
	_, err := gcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcuo *GitCommitUpdateOne) ExecX(ctx context.Context) {
	if err := gcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gcuo *GitCommitUpdateOne) sqlSave(ctx context.Context) (_node *GitCommit, err error) {
	_spec := sqlgraph.NewUpdateSpec(gitcommit.Table, gitcommit.Columns, sqlgraph.NewFieldSpec(gitcommit.FieldID, field.TypeString))
	id, ok := gcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GitCommit.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gitcommit.FieldID)
		for _, f := range fields {
			if !gitcommit.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != gitcommit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gcuo.mutation.Message(); ok {
		_spec.SetField(gitcommit.FieldMessage, field.TypeString, value)
	}
	if value, ok := gcuo.mutation.AuthorLogin(); ok {
		_spec.SetField(gitcommit.FieldAuthorLogin, field.TypeString, value)
	}
	if value, ok := gcuo.mutation.AuthorID(); ok {
		_spec.SetField(gitcommit.FieldAuthorID, field.TypeInt64, value)
	}
	if value, ok := gcuo.mutation.AddedAuthorID(); ok {
		_spec.AddField(gitcommit.FieldAuthorID, field.TypeInt64, value)
	}
	if value, ok := gcuo.mutation.Date(); ok {
		_spec.SetField(gitcommit.FieldDate, field.TypeTime, value)
	}
	if gcuo.mutation.RepositoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gitcommit.RepositoryTable,
			Columns: []string{gitcommit.RepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(repository.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gcuo.mutation.RepositoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gitcommit.RepositoryTable,
			Columns: []string{gitcommit.RepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(repository.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GitCommit{config: gcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gitcommit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gcuo.mutation.done = true
	return _node, nil
}
