// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-faster/bot/internal/ent/predicate"
	"github.com/go-faster/bot/internal/ent/prnotification"
)

// PRNotificationDelete is the builder for deleting a PRNotification entity.
type PRNotificationDelete struct {
	config
	hooks    []Hook
	mutation *PRNotificationMutation
}

// Where appends a list predicates to the PRNotificationDelete builder.
func (pnd *PRNotificationDelete) Where(ps ...predicate.PRNotification) *PRNotificationDelete {
	pnd.mutation.Where(ps...)
	return pnd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pnd *PRNotificationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, PRNotificationMutation](ctx, pnd.sqlExec, pnd.mutation, pnd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pnd *PRNotificationDelete) ExecX(ctx context.Context) int {
	n, err := pnd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pnd *PRNotificationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(prnotification.Table, sqlgraph.NewFieldSpec(prnotification.FieldID, field.TypeInt))
	if ps := pnd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pnd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pnd.mutation.done = true
	return affected, err
}

// PRNotificationDeleteOne is the builder for deleting a single PRNotification entity.
type PRNotificationDeleteOne struct {
	pnd *PRNotificationDelete
}

// Where appends a list predicates to the PRNotificationDelete builder.
func (pndo *PRNotificationDeleteOne) Where(ps ...predicate.PRNotification) *PRNotificationDeleteOne {
	pndo.pnd.mutation.Where(ps...)
	return pndo
}

// Exec executes the deletion query.
func (pndo *PRNotificationDeleteOne) Exec(ctx context.Context) error {
	n, err := pndo.pnd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{prnotification.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pndo *PRNotificationDeleteOne) ExecX(ctx context.Context) {
	if err := pndo.Exec(ctx); err != nil {
		panic(err)
	}
}
