// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-faster/bot/internal/ent/lastchannelmessage"
	"github.com/go-faster/bot/internal/ent/predicate"
)

// LastChannelMessageQuery is the builder for querying LastChannelMessage entities.
type LastChannelMessageQuery struct {
	config
	ctx        *QueryContext
	order      []lastchannelmessage.OrderOption
	inters     []Interceptor
	predicates []predicate.LastChannelMessage
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LastChannelMessageQuery builder.
func (lcmq *LastChannelMessageQuery) Where(ps ...predicate.LastChannelMessage) *LastChannelMessageQuery {
	lcmq.predicates = append(lcmq.predicates, ps...)
	return lcmq
}

// Limit the number of records to be returned by this query.
func (lcmq *LastChannelMessageQuery) Limit(limit int) *LastChannelMessageQuery {
	lcmq.ctx.Limit = &limit
	return lcmq
}

// Offset to start from.
func (lcmq *LastChannelMessageQuery) Offset(offset int) *LastChannelMessageQuery {
	lcmq.ctx.Offset = &offset
	return lcmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lcmq *LastChannelMessageQuery) Unique(unique bool) *LastChannelMessageQuery {
	lcmq.ctx.Unique = &unique
	return lcmq
}

// Order specifies how the records should be ordered.
func (lcmq *LastChannelMessageQuery) Order(o ...lastchannelmessage.OrderOption) *LastChannelMessageQuery {
	lcmq.order = append(lcmq.order, o...)
	return lcmq
}

// First returns the first LastChannelMessage entity from the query.
// Returns a *NotFoundError when no LastChannelMessage was found.
func (lcmq *LastChannelMessageQuery) First(ctx context.Context) (*LastChannelMessage, error) {
	nodes, err := lcmq.Limit(1).All(setContextOp(ctx, lcmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{lastchannelmessage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lcmq *LastChannelMessageQuery) FirstX(ctx context.Context) *LastChannelMessage {
	node, err := lcmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LastChannelMessage ID from the query.
// Returns a *NotFoundError when no LastChannelMessage ID was found.
func (lcmq *LastChannelMessageQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lcmq.Limit(1).IDs(setContextOp(ctx, lcmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{lastchannelmessage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lcmq *LastChannelMessageQuery) FirstIDX(ctx context.Context) int64 {
	id, err := lcmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LastChannelMessage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LastChannelMessage entity is found.
// Returns a *NotFoundError when no LastChannelMessage entities are found.
func (lcmq *LastChannelMessageQuery) Only(ctx context.Context) (*LastChannelMessage, error) {
	nodes, err := lcmq.Limit(2).All(setContextOp(ctx, lcmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{lastchannelmessage.Label}
	default:
		return nil, &NotSingularError{lastchannelmessage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lcmq *LastChannelMessageQuery) OnlyX(ctx context.Context) *LastChannelMessage {
	node, err := lcmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LastChannelMessage ID in the query.
// Returns a *NotSingularError when more than one LastChannelMessage ID is found.
// Returns a *NotFoundError when no entities are found.
func (lcmq *LastChannelMessageQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lcmq.Limit(2).IDs(setContextOp(ctx, lcmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{lastchannelmessage.Label}
	default:
		err = &NotSingularError{lastchannelmessage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lcmq *LastChannelMessageQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := lcmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LastChannelMessages.
func (lcmq *LastChannelMessageQuery) All(ctx context.Context) ([]*LastChannelMessage, error) {
	ctx = setContextOp(ctx, lcmq.ctx, "All")
	if err := lcmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LastChannelMessage, *LastChannelMessageQuery]()
	return withInterceptors[[]*LastChannelMessage](ctx, lcmq, qr, lcmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lcmq *LastChannelMessageQuery) AllX(ctx context.Context) []*LastChannelMessage {
	nodes, err := lcmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LastChannelMessage IDs.
func (lcmq *LastChannelMessageQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if lcmq.ctx.Unique == nil && lcmq.path != nil {
		lcmq.Unique(true)
	}
	ctx = setContextOp(ctx, lcmq.ctx, "IDs")
	if err = lcmq.Select(lastchannelmessage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lcmq *LastChannelMessageQuery) IDsX(ctx context.Context) []int64 {
	ids, err := lcmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lcmq *LastChannelMessageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lcmq.ctx, "Count")
	if err := lcmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lcmq, querierCount[*LastChannelMessageQuery](), lcmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lcmq *LastChannelMessageQuery) CountX(ctx context.Context) int {
	count, err := lcmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lcmq *LastChannelMessageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lcmq.ctx, "Exist")
	switch _, err := lcmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lcmq *LastChannelMessageQuery) ExistX(ctx context.Context) bool {
	exist, err := lcmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LastChannelMessageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lcmq *LastChannelMessageQuery) Clone() *LastChannelMessageQuery {
	if lcmq == nil {
		return nil
	}
	return &LastChannelMessageQuery{
		config:     lcmq.config,
		ctx:        lcmq.ctx.Clone(),
		order:      append([]lastchannelmessage.OrderOption{}, lcmq.order...),
		inters:     append([]Interceptor{}, lcmq.inters...),
		predicates: append([]predicate.LastChannelMessage{}, lcmq.predicates...),
		// clone intermediate query.
		sql:  lcmq.sql.Clone(),
		path: lcmq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		MessageID int `json:"message_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LastChannelMessage.Query().
//		GroupBy(lastchannelmessage.FieldMessageID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lcmq *LastChannelMessageQuery) GroupBy(field string, fields ...string) *LastChannelMessageGroupBy {
	lcmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LastChannelMessageGroupBy{build: lcmq}
	grbuild.flds = &lcmq.ctx.Fields
	grbuild.label = lastchannelmessage.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		MessageID int `json:"message_id,omitempty"`
//	}
//
//	client.LastChannelMessage.Query().
//		Select(lastchannelmessage.FieldMessageID).
//		Scan(ctx, &v)
func (lcmq *LastChannelMessageQuery) Select(fields ...string) *LastChannelMessageSelect {
	lcmq.ctx.Fields = append(lcmq.ctx.Fields, fields...)
	sbuild := &LastChannelMessageSelect{LastChannelMessageQuery: lcmq}
	sbuild.label = lastchannelmessage.Label
	sbuild.flds, sbuild.scan = &lcmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LastChannelMessageSelect configured with the given aggregations.
func (lcmq *LastChannelMessageQuery) Aggregate(fns ...AggregateFunc) *LastChannelMessageSelect {
	return lcmq.Select().Aggregate(fns...)
}

func (lcmq *LastChannelMessageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lcmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lcmq); err != nil {
				return err
			}
		}
	}
	for _, f := range lcmq.ctx.Fields {
		if !lastchannelmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lcmq.path != nil {
		prev, err := lcmq.path(ctx)
		if err != nil {
			return err
		}
		lcmq.sql = prev
	}
	return nil
}

func (lcmq *LastChannelMessageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LastChannelMessage, error) {
	var (
		nodes = []*LastChannelMessage{}
		_spec = lcmq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LastChannelMessage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LastChannelMessage{config: lcmq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lcmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (lcmq *LastChannelMessageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lcmq.querySpec()
	_spec.Node.Columns = lcmq.ctx.Fields
	if len(lcmq.ctx.Fields) > 0 {
		_spec.Unique = lcmq.ctx.Unique != nil && *lcmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lcmq.driver, _spec)
}

func (lcmq *LastChannelMessageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(lastchannelmessage.Table, lastchannelmessage.Columns, sqlgraph.NewFieldSpec(lastchannelmessage.FieldID, field.TypeInt64))
	_spec.From = lcmq.sql
	if unique := lcmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lcmq.path != nil {
		_spec.Unique = true
	}
	if fields := lcmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lastchannelmessage.FieldID)
		for i := range fields {
			if fields[i] != lastchannelmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lcmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lcmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lcmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lcmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lcmq *LastChannelMessageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lcmq.driver.Dialect())
	t1 := builder.Table(lastchannelmessage.Table)
	columns := lcmq.ctx.Fields
	if len(columns) == 0 {
		columns = lastchannelmessage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lcmq.sql != nil {
		selector = lcmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lcmq.ctx.Unique != nil && *lcmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lcmq.predicates {
		p(selector)
	}
	for _, p := range lcmq.order {
		p(selector)
	}
	if offset := lcmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lcmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LastChannelMessageGroupBy is the group-by builder for LastChannelMessage entities.
type LastChannelMessageGroupBy struct {
	selector
	build *LastChannelMessageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lcmgb *LastChannelMessageGroupBy) Aggregate(fns ...AggregateFunc) *LastChannelMessageGroupBy {
	lcmgb.fns = append(lcmgb.fns, fns...)
	return lcmgb
}

// Scan applies the selector query and scans the result into the given value.
func (lcmgb *LastChannelMessageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lcmgb.build.ctx, "GroupBy")
	if err := lcmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LastChannelMessageQuery, *LastChannelMessageGroupBy](ctx, lcmgb.build, lcmgb, lcmgb.build.inters, v)
}

func (lcmgb *LastChannelMessageGroupBy) sqlScan(ctx context.Context, root *LastChannelMessageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lcmgb.fns))
	for _, fn := range lcmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lcmgb.flds)+len(lcmgb.fns))
		for _, f := range *lcmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lcmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lcmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LastChannelMessageSelect is the builder for selecting fields of LastChannelMessage entities.
type LastChannelMessageSelect struct {
	*LastChannelMessageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lcms *LastChannelMessageSelect) Aggregate(fns ...AggregateFunc) *LastChannelMessageSelect {
	lcms.fns = append(lcms.fns, fns...)
	return lcms
}

// Scan applies the selector query and scans the result into the given value.
func (lcms *LastChannelMessageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lcms.ctx, "Select")
	if err := lcms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LastChannelMessageQuery, *LastChannelMessageSelect](ctx, lcms.LastChannelMessageQuery, lcms, lcms.inters, v)
}

func (lcms *LastChannelMessageSelect) sqlScan(ctx context.Context, root *LastChannelMessageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lcms.fns))
	for _, fn := range lcms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lcms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lcms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
