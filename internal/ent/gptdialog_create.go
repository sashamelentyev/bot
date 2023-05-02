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
	"github.com/go-faster/bot/internal/ent/gptdialog"
)

// GPTDialogCreate is the builder for creating a GPTDialog entity.
type GPTDialogCreate struct {
	config
	mutation *GPTDialogMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPeerID sets the "peer_id" field.
func (gdc *GPTDialogCreate) SetPeerID(s string) *GPTDialogCreate {
	gdc.mutation.SetPeerID(s)
	return gdc
}

// SetPromptMsgID sets the "prompt_msg_id" field.
func (gdc *GPTDialogCreate) SetPromptMsgID(i int) *GPTDialogCreate {
	gdc.mutation.SetPromptMsgID(i)
	return gdc
}

// SetPromptMsg sets the "prompt_msg" field.
func (gdc *GPTDialogCreate) SetPromptMsg(s string) *GPTDialogCreate {
	gdc.mutation.SetPromptMsg(s)
	return gdc
}

// SetGptMsgID sets the "gpt_msg_id" field.
func (gdc *GPTDialogCreate) SetGptMsgID(i int) *GPTDialogCreate {
	gdc.mutation.SetGptMsgID(i)
	return gdc
}

// SetGptMsg sets the "gpt_msg" field.
func (gdc *GPTDialogCreate) SetGptMsg(s string) *GPTDialogCreate {
	gdc.mutation.SetGptMsg(s)
	return gdc
}

// SetThreadTopMsgID sets the "thread_top_msg_id" field.
func (gdc *GPTDialogCreate) SetThreadTopMsgID(i int) *GPTDialogCreate {
	gdc.mutation.SetThreadTopMsgID(i)
	return gdc
}

// SetNillableThreadTopMsgID sets the "thread_top_msg_id" field if the given value is not nil.
func (gdc *GPTDialogCreate) SetNillableThreadTopMsgID(i *int) *GPTDialogCreate {
	if i != nil {
		gdc.SetThreadTopMsgID(*i)
	}
	return gdc
}

// SetCreatedAt sets the "created_at" field.
func (gdc *GPTDialogCreate) SetCreatedAt(t time.Time) *GPTDialogCreate {
	gdc.mutation.SetCreatedAt(t)
	return gdc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gdc *GPTDialogCreate) SetNillableCreatedAt(t *time.Time) *GPTDialogCreate {
	if t != nil {
		gdc.SetCreatedAt(*t)
	}
	return gdc
}

// Mutation returns the GPTDialogMutation object of the builder.
func (gdc *GPTDialogCreate) Mutation() *GPTDialogMutation {
	return gdc.mutation
}

// Save creates the GPTDialog in the database.
func (gdc *GPTDialogCreate) Save(ctx context.Context) (*GPTDialog, error) {
	gdc.defaults()
	return withHooks(ctx, gdc.sqlSave, gdc.mutation, gdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gdc *GPTDialogCreate) SaveX(ctx context.Context) *GPTDialog {
	v, err := gdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gdc *GPTDialogCreate) Exec(ctx context.Context) error {
	_, err := gdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gdc *GPTDialogCreate) ExecX(ctx context.Context) {
	if err := gdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gdc *GPTDialogCreate) defaults() {
	if _, ok := gdc.mutation.CreatedAt(); !ok {
		v := gptdialog.DefaultCreatedAt()
		gdc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gdc *GPTDialogCreate) check() error {
	if _, ok := gdc.mutation.PeerID(); !ok {
		return &ValidationError{Name: "peer_id", err: errors.New(`ent: missing required field "GPTDialog.peer_id"`)}
	}
	if _, ok := gdc.mutation.PromptMsgID(); !ok {
		return &ValidationError{Name: "prompt_msg_id", err: errors.New(`ent: missing required field "GPTDialog.prompt_msg_id"`)}
	}
	if _, ok := gdc.mutation.PromptMsg(); !ok {
		return &ValidationError{Name: "prompt_msg", err: errors.New(`ent: missing required field "GPTDialog.prompt_msg"`)}
	}
	if _, ok := gdc.mutation.GptMsgID(); !ok {
		return &ValidationError{Name: "gpt_msg_id", err: errors.New(`ent: missing required field "GPTDialog.gpt_msg_id"`)}
	}
	if _, ok := gdc.mutation.GptMsg(); !ok {
		return &ValidationError{Name: "gpt_msg", err: errors.New(`ent: missing required field "GPTDialog.gpt_msg"`)}
	}
	if _, ok := gdc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "GPTDialog.created_at"`)}
	}
	return nil
}

func (gdc *GPTDialogCreate) sqlSave(ctx context.Context) (*GPTDialog, error) {
	if err := gdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	gdc.mutation.id = &_node.ID
	gdc.mutation.done = true
	return _node, nil
}

func (gdc *GPTDialogCreate) createSpec() (*GPTDialog, *sqlgraph.CreateSpec) {
	var (
		_node = &GPTDialog{config: gdc.config}
		_spec = sqlgraph.NewCreateSpec(gptdialog.Table, sqlgraph.NewFieldSpec(gptdialog.FieldID, field.TypeInt))
	)
	_spec.OnConflict = gdc.conflict
	if value, ok := gdc.mutation.PeerID(); ok {
		_spec.SetField(gptdialog.FieldPeerID, field.TypeString, value)
		_node.PeerID = value
	}
	if value, ok := gdc.mutation.PromptMsgID(); ok {
		_spec.SetField(gptdialog.FieldPromptMsgID, field.TypeInt, value)
		_node.PromptMsgID = value
	}
	if value, ok := gdc.mutation.PromptMsg(); ok {
		_spec.SetField(gptdialog.FieldPromptMsg, field.TypeString, value)
		_node.PromptMsg = value
	}
	if value, ok := gdc.mutation.GptMsgID(); ok {
		_spec.SetField(gptdialog.FieldGptMsgID, field.TypeInt, value)
		_node.GptMsgID = value
	}
	if value, ok := gdc.mutation.GptMsg(); ok {
		_spec.SetField(gptdialog.FieldGptMsg, field.TypeString, value)
		_node.GptMsg = value
	}
	if value, ok := gdc.mutation.ThreadTopMsgID(); ok {
		_spec.SetField(gptdialog.FieldThreadTopMsgID, field.TypeInt, value)
		_node.ThreadTopMsgID = value
	}
	if value, ok := gdc.mutation.CreatedAt(); ok {
		_spec.SetField(gptdialog.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GPTDialog.Create().
//		SetPeerID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GPTDialogUpsert) {
//			SetPeerID(v+v).
//		}).
//		Exec(ctx)
func (gdc *GPTDialogCreate) OnConflict(opts ...sql.ConflictOption) *GPTDialogUpsertOne {
	gdc.conflict = opts
	return &GPTDialogUpsertOne{
		create: gdc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GPTDialog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gdc *GPTDialogCreate) OnConflictColumns(columns ...string) *GPTDialogUpsertOne {
	gdc.conflict = append(gdc.conflict, sql.ConflictColumns(columns...))
	return &GPTDialogUpsertOne{
		create: gdc,
	}
}

type (
	// GPTDialogUpsertOne is the builder for "upsert"-ing
	//  one GPTDialog node.
	GPTDialogUpsertOne struct {
		create *GPTDialogCreate
	}

	// GPTDialogUpsert is the "OnConflict" setter.
	GPTDialogUpsert struct {
		*sql.UpdateSet
	}
)

// SetPromptMsgID sets the "prompt_msg_id" field.
func (u *GPTDialogUpsert) SetPromptMsgID(v int) *GPTDialogUpsert {
	u.Set(gptdialog.FieldPromptMsgID, v)
	return u
}

// UpdatePromptMsgID sets the "prompt_msg_id" field to the value that was provided on create.
func (u *GPTDialogUpsert) UpdatePromptMsgID() *GPTDialogUpsert {
	u.SetExcluded(gptdialog.FieldPromptMsgID)
	return u
}

// AddPromptMsgID adds v to the "prompt_msg_id" field.
func (u *GPTDialogUpsert) AddPromptMsgID(v int) *GPTDialogUpsert {
	u.Add(gptdialog.FieldPromptMsgID, v)
	return u
}

// SetPromptMsg sets the "prompt_msg" field.
func (u *GPTDialogUpsert) SetPromptMsg(v string) *GPTDialogUpsert {
	u.Set(gptdialog.FieldPromptMsg, v)
	return u
}

// UpdatePromptMsg sets the "prompt_msg" field to the value that was provided on create.
func (u *GPTDialogUpsert) UpdatePromptMsg() *GPTDialogUpsert {
	u.SetExcluded(gptdialog.FieldPromptMsg)
	return u
}

// SetGptMsgID sets the "gpt_msg_id" field.
func (u *GPTDialogUpsert) SetGptMsgID(v int) *GPTDialogUpsert {
	u.Set(gptdialog.FieldGptMsgID, v)
	return u
}

// UpdateGptMsgID sets the "gpt_msg_id" field to the value that was provided on create.
func (u *GPTDialogUpsert) UpdateGptMsgID() *GPTDialogUpsert {
	u.SetExcluded(gptdialog.FieldGptMsgID)
	return u
}

// AddGptMsgID adds v to the "gpt_msg_id" field.
func (u *GPTDialogUpsert) AddGptMsgID(v int) *GPTDialogUpsert {
	u.Add(gptdialog.FieldGptMsgID, v)
	return u
}

// SetGptMsg sets the "gpt_msg" field.
func (u *GPTDialogUpsert) SetGptMsg(v string) *GPTDialogUpsert {
	u.Set(gptdialog.FieldGptMsg, v)
	return u
}

// UpdateGptMsg sets the "gpt_msg" field to the value that was provided on create.
func (u *GPTDialogUpsert) UpdateGptMsg() *GPTDialogUpsert {
	u.SetExcluded(gptdialog.FieldGptMsg)
	return u
}

// SetThreadTopMsgID sets the "thread_top_msg_id" field.
func (u *GPTDialogUpsert) SetThreadTopMsgID(v int) *GPTDialogUpsert {
	u.Set(gptdialog.FieldThreadTopMsgID, v)
	return u
}

// UpdateThreadTopMsgID sets the "thread_top_msg_id" field to the value that was provided on create.
func (u *GPTDialogUpsert) UpdateThreadTopMsgID() *GPTDialogUpsert {
	u.SetExcluded(gptdialog.FieldThreadTopMsgID)
	return u
}

// AddThreadTopMsgID adds v to the "thread_top_msg_id" field.
func (u *GPTDialogUpsert) AddThreadTopMsgID(v int) *GPTDialogUpsert {
	u.Add(gptdialog.FieldThreadTopMsgID, v)
	return u
}

// ClearThreadTopMsgID clears the value of the "thread_top_msg_id" field.
func (u *GPTDialogUpsert) ClearThreadTopMsgID() *GPTDialogUpsert {
	u.SetNull(gptdialog.FieldThreadTopMsgID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.GPTDialog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GPTDialogUpsertOne) UpdateNewValues() *GPTDialogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.PeerID(); exists {
			s.SetIgnore(gptdialog.FieldPeerID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(gptdialog.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GPTDialog.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GPTDialogUpsertOne) Ignore() *GPTDialogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GPTDialogUpsertOne) DoNothing() *GPTDialogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GPTDialogCreate.OnConflict
// documentation for more info.
func (u *GPTDialogUpsertOne) Update(set func(*GPTDialogUpsert)) *GPTDialogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GPTDialogUpsert{UpdateSet: update})
	}))
	return u
}

// SetPromptMsgID sets the "prompt_msg_id" field.
func (u *GPTDialogUpsertOne) SetPromptMsgID(v int) *GPTDialogUpsertOne {
	return u.Update(func(s *GPTDialogUpsert) {
		s.SetPromptMsgID(v)
	})
}

// AddPromptMsgID adds v to the "prompt_msg_id" field.
func (u *GPTDialogUpsertOne) AddPromptMsgID(v int) *GPTDialogUpsertOne {
	return u.Update(func(s *GPTDialogUpsert) {
		s.AddPromptMsgID(v)
	})
}

// UpdatePromptMsgID sets the "prompt_msg_id" field to the value that was provided on create.
func (u *GPTDialogUpsertOne) UpdatePromptMsgID() *GPTDialogUpsertOne {
	return u.Update(func(s *GPTDialogUpsert) {
		s.UpdatePromptMsgID()
	})
}

// SetPromptMsg sets the "prompt_msg" field.
func (u *GPTDialogUpsertOne) SetPromptMsg(v string) *GPTDialogUpsertOne {
	return u.Update(func(s *GPTDialogUpsert) {
		s.SetPromptMsg(v)
	})
}

// UpdatePromptMsg sets the "prompt_msg" field to the value that was provided on create.
func (u *GPTDialogUpsertOne) UpdatePromptMsg() *GPTDialogUpsertOne {
	return u.Update(func(s *GPTDialogUpsert) {
		s.UpdatePromptMsg()
	})
}

// SetGptMsgID sets the "gpt_msg_id" field.
func (u *GPTDialogUpsertOne) SetGptMsgID(v int) *GPTDialogUpsertOne {
	return u.Update(func(s *GPTDialogUpsert) {
		s.SetGptMsgID(v)
	})
}

// AddGptMsgID adds v to the "gpt_msg_id" field.
func (u *GPTDialogUpsertOne) AddGptMsgID(v int) *GPTDialogUpsertOne {
	return u.Update(func(s *GPTDialogUpsert) {
		s.AddGptMsgID(v)
	})
}

// UpdateGptMsgID sets the "gpt_msg_id" field to the value that was provided on create.
func (u *GPTDialogUpsertOne) UpdateGptMsgID() *GPTDialogUpsertOne {
	return u.Update(func(s *GPTDialogUpsert) {
		s.UpdateGptMsgID()
	})
}

// SetGptMsg sets the "gpt_msg" field.
func (u *GPTDialogUpsertOne) SetGptMsg(v string) *GPTDialogUpsertOne {
	return u.Update(func(s *GPTDialogUpsert) {
		s.SetGptMsg(v)
	})
}

// UpdateGptMsg sets the "gpt_msg" field to the value that was provided on create.
func (u *GPTDialogUpsertOne) UpdateGptMsg() *GPTDialogUpsertOne {
	return u.Update(func(s *GPTDialogUpsert) {
		s.UpdateGptMsg()
	})
}

// SetThreadTopMsgID sets the "thread_top_msg_id" field.
func (u *GPTDialogUpsertOne) SetThreadTopMsgID(v int) *GPTDialogUpsertOne {
	return u.Update(func(s *GPTDialogUpsert) {
		s.SetThreadTopMsgID(v)
	})
}

// AddThreadTopMsgID adds v to the "thread_top_msg_id" field.
func (u *GPTDialogUpsertOne) AddThreadTopMsgID(v int) *GPTDialogUpsertOne {
	return u.Update(func(s *GPTDialogUpsert) {
		s.AddThreadTopMsgID(v)
	})
}

// UpdateThreadTopMsgID sets the "thread_top_msg_id" field to the value that was provided on create.
func (u *GPTDialogUpsertOne) UpdateThreadTopMsgID() *GPTDialogUpsertOne {
	return u.Update(func(s *GPTDialogUpsert) {
		s.UpdateThreadTopMsgID()
	})
}

// ClearThreadTopMsgID clears the value of the "thread_top_msg_id" field.
func (u *GPTDialogUpsertOne) ClearThreadTopMsgID() *GPTDialogUpsertOne {
	return u.Update(func(s *GPTDialogUpsert) {
		s.ClearThreadTopMsgID()
	})
}

// Exec executes the query.
func (u *GPTDialogUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GPTDialogCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GPTDialogUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GPTDialogUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GPTDialogUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GPTDialogCreateBulk is the builder for creating many GPTDialog entities in bulk.
type GPTDialogCreateBulk struct {
	config
	builders []*GPTDialogCreate
	conflict []sql.ConflictOption
}

// Save creates the GPTDialog entities in the database.
func (gdcb *GPTDialogCreateBulk) Save(ctx context.Context) ([]*GPTDialog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gdcb.builders))
	nodes := make([]*GPTDialog, len(gdcb.builders))
	mutators := make([]Mutator, len(gdcb.builders))
	for i := range gdcb.builders {
		func(i int, root context.Context) {
			builder := gdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GPTDialogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gdcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gdcb *GPTDialogCreateBulk) SaveX(ctx context.Context) []*GPTDialog {
	v, err := gdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gdcb *GPTDialogCreateBulk) Exec(ctx context.Context) error {
	_, err := gdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gdcb *GPTDialogCreateBulk) ExecX(ctx context.Context) {
	if err := gdcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GPTDialog.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GPTDialogUpsert) {
//			SetPeerID(v+v).
//		}).
//		Exec(ctx)
func (gdcb *GPTDialogCreateBulk) OnConflict(opts ...sql.ConflictOption) *GPTDialogUpsertBulk {
	gdcb.conflict = opts
	return &GPTDialogUpsertBulk{
		create: gdcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GPTDialog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gdcb *GPTDialogCreateBulk) OnConflictColumns(columns ...string) *GPTDialogUpsertBulk {
	gdcb.conflict = append(gdcb.conflict, sql.ConflictColumns(columns...))
	return &GPTDialogUpsertBulk{
		create: gdcb,
	}
}

// GPTDialogUpsertBulk is the builder for "upsert"-ing
// a bulk of GPTDialog nodes.
type GPTDialogUpsertBulk struct {
	create *GPTDialogCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GPTDialog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GPTDialogUpsertBulk) UpdateNewValues() *GPTDialogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.PeerID(); exists {
				s.SetIgnore(gptdialog.FieldPeerID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(gptdialog.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GPTDialog.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GPTDialogUpsertBulk) Ignore() *GPTDialogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GPTDialogUpsertBulk) DoNothing() *GPTDialogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GPTDialogCreateBulk.OnConflict
// documentation for more info.
func (u *GPTDialogUpsertBulk) Update(set func(*GPTDialogUpsert)) *GPTDialogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GPTDialogUpsert{UpdateSet: update})
	}))
	return u
}

// SetPromptMsgID sets the "prompt_msg_id" field.
func (u *GPTDialogUpsertBulk) SetPromptMsgID(v int) *GPTDialogUpsertBulk {
	return u.Update(func(s *GPTDialogUpsert) {
		s.SetPromptMsgID(v)
	})
}

// AddPromptMsgID adds v to the "prompt_msg_id" field.
func (u *GPTDialogUpsertBulk) AddPromptMsgID(v int) *GPTDialogUpsertBulk {
	return u.Update(func(s *GPTDialogUpsert) {
		s.AddPromptMsgID(v)
	})
}

// UpdatePromptMsgID sets the "prompt_msg_id" field to the value that was provided on create.
func (u *GPTDialogUpsertBulk) UpdatePromptMsgID() *GPTDialogUpsertBulk {
	return u.Update(func(s *GPTDialogUpsert) {
		s.UpdatePromptMsgID()
	})
}

// SetPromptMsg sets the "prompt_msg" field.
func (u *GPTDialogUpsertBulk) SetPromptMsg(v string) *GPTDialogUpsertBulk {
	return u.Update(func(s *GPTDialogUpsert) {
		s.SetPromptMsg(v)
	})
}

// UpdatePromptMsg sets the "prompt_msg" field to the value that was provided on create.
func (u *GPTDialogUpsertBulk) UpdatePromptMsg() *GPTDialogUpsertBulk {
	return u.Update(func(s *GPTDialogUpsert) {
		s.UpdatePromptMsg()
	})
}

// SetGptMsgID sets the "gpt_msg_id" field.
func (u *GPTDialogUpsertBulk) SetGptMsgID(v int) *GPTDialogUpsertBulk {
	return u.Update(func(s *GPTDialogUpsert) {
		s.SetGptMsgID(v)
	})
}

// AddGptMsgID adds v to the "gpt_msg_id" field.
func (u *GPTDialogUpsertBulk) AddGptMsgID(v int) *GPTDialogUpsertBulk {
	return u.Update(func(s *GPTDialogUpsert) {
		s.AddGptMsgID(v)
	})
}

// UpdateGptMsgID sets the "gpt_msg_id" field to the value that was provided on create.
func (u *GPTDialogUpsertBulk) UpdateGptMsgID() *GPTDialogUpsertBulk {
	return u.Update(func(s *GPTDialogUpsert) {
		s.UpdateGptMsgID()
	})
}

// SetGptMsg sets the "gpt_msg" field.
func (u *GPTDialogUpsertBulk) SetGptMsg(v string) *GPTDialogUpsertBulk {
	return u.Update(func(s *GPTDialogUpsert) {
		s.SetGptMsg(v)
	})
}

// UpdateGptMsg sets the "gpt_msg" field to the value that was provided on create.
func (u *GPTDialogUpsertBulk) UpdateGptMsg() *GPTDialogUpsertBulk {
	return u.Update(func(s *GPTDialogUpsert) {
		s.UpdateGptMsg()
	})
}

// SetThreadTopMsgID sets the "thread_top_msg_id" field.
func (u *GPTDialogUpsertBulk) SetThreadTopMsgID(v int) *GPTDialogUpsertBulk {
	return u.Update(func(s *GPTDialogUpsert) {
		s.SetThreadTopMsgID(v)
	})
}

// AddThreadTopMsgID adds v to the "thread_top_msg_id" field.
func (u *GPTDialogUpsertBulk) AddThreadTopMsgID(v int) *GPTDialogUpsertBulk {
	return u.Update(func(s *GPTDialogUpsert) {
		s.AddThreadTopMsgID(v)
	})
}

// UpdateThreadTopMsgID sets the "thread_top_msg_id" field to the value that was provided on create.
func (u *GPTDialogUpsertBulk) UpdateThreadTopMsgID() *GPTDialogUpsertBulk {
	return u.Update(func(s *GPTDialogUpsert) {
		s.UpdateThreadTopMsgID()
	})
}

// ClearThreadTopMsgID clears the value of the "thread_top_msg_id" field.
func (u *GPTDialogUpsertBulk) ClearThreadTopMsgID() *GPTDialogUpsertBulk {
	return u.Update(func(s *GPTDialogUpsert) {
		s.ClearThreadTopMsgID()
	})
}

// Exec executes the query.
func (u *GPTDialogUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GPTDialogCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GPTDialogCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GPTDialogUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
