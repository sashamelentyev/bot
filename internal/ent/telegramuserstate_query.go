// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-faster/bot/internal/ent/predicate"
	"github.com/go-faster/bot/internal/ent/telegramchannelstate"
	"github.com/go-faster/bot/internal/ent/telegramuserstate"
)

// TelegramUserStateQuery is the builder for querying TelegramUserState entities.
type TelegramUserStateQuery struct {
	config
	ctx               *QueryContext
	order             []telegramuserstate.OrderOption
	inters            []Interceptor
	predicates        []predicate.TelegramUserState
	withChannels      *TelegramChannelStateQuery
	withNamedChannels map[string]*TelegramChannelStateQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TelegramUserStateQuery builder.
func (tusq *TelegramUserStateQuery) Where(ps ...predicate.TelegramUserState) *TelegramUserStateQuery {
	tusq.predicates = append(tusq.predicates, ps...)
	return tusq
}

// Limit the number of records to be returned by this query.
func (tusq *TelegramUserStateQuery) Limit(limit int) *TelegramUserStateQuery {
	tusq.ctx.Limit = &limit
	return tusq
}

// Offset to start from.
func (tusq *TelegramUserStateQuery) Offset(offset int) *TelegramUserStateQuery {
	tusq.ctx.Offset = &offset
	return tusq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tusq *TelegramUserStateQuery) Unique(unique bool) *TelegramUserStateQuery {
	tusq.ctx.Unique = &unique
	return tusq
}

// Order specifies how the records should be ordered.
func (tusq *TelegramUserStateQuery) Order(o ...telegramuserstate.OrderOption) *TelegramUserStateQuery {
	tusq.order = append(tusq.order, o...)
	return tusq
}

// QueryChannels chains the current query on the "channels" edge.
func (tusq *TelegramUserStateQuery) QueryChannels() *TelegramChannelStateQuery {
	query := (&TelegramChannelStateClient{config: tusq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tusq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tusq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(telegramuserstate.Table, telegramuserstate.FieldID, selector),
			sqlgraph.To(telegramchannelstate.Table, telegramchannelstate.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, telegramuserstate.ChannelsTable, telegramuserstate.ChannelsColumn),
		)
		fromU = sqlgraph.SetNeighbors(tusq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TelegramUserState entity from the query.
// Returns a *NotFoundError when no TelegramUserState was found.
func (tusq *TelegramUserStateQuery) First(ctx context.Context) (*TelegramUserState, error) {
	nodes, err := tusq.Limit(1).All(setContextOp(ctx, tusq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{telegramuserstate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tusq *TelegramUserStateQuery) FirstX(ctx context.Context) *TelegramUserState {
	node, err := tusq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TelegramUserState ID from the query.
// Returns a *NotFoundError when no TelegramUserState ID was found.
func (tusq *TelegramUserStateQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = tusq.Limit(1).IDs(setContextOp(ctx, tusq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{telegramuserstate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tusq *TelegramUserStateQuery) FirstIDX(ctx context.Context) int64 {
	id, err := tusq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TelegramUserState entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TelegramUserState entity is found.
// Returns a *NotFoundError when no TelegramUserState entities are found.
func (tusq *TelegramUserStateQuery) Only(ctx context.Context) (*TelegramUserState, error) {
	nodes, err := tusq.Limit(2).All(setContextOp(ctx, tusq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{telegramuserstate.Label}
	default:
		return nil, &NotSingularError{telegramuserstate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tusq *TelegramUserStateQuery) OnlyX(ctx context.Context) *TelegramUserState {
	node, err := tusq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TelegramUserState ID in the query.
// Returns a *NotSingularError when more than one TelegramUserState ID is found.
// Returns a *NotFoundError when no entities are found.
func (tusq *TelegramUserStateQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = tusq.Limit(2).IDs(setContextOp(ctx, tusq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{telegramuserstate.Label}
	default:
		err = &NotSingularError{telegramuserstate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tusq *TelegramUserStateQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := tusq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TelegramUserStates.
func (tusq *TelegramUserStateQuery) All(ctx context.Context) ([]*TelegramUserState, error) {
	ctx = setContextOp(ctx, tusq.ctx, "All")
	if err := tusq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TelegramUserState, *TelegramUserStateQuery]()
	return withInterceptors[[]*TelegramUserState](ctx, tusq, qr, tusq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tusq *TelegramUserStateQuery) AllX(ctx context.Context) []*TelegramUserState {
	nodes, err := tusq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TelegramUserState IDs.
func (tusq *TelegramUserStateQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if tusq.ctx.Unique == nil && tusq.path != nil {
		tusq.Unique(true)
	}
	ctx = setContextOp(ctx, tusq.ctx, "IDs")
	if err = tusq.Select(telegramuserstate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tusq *TelegramUserStateQuery) IDsX(ctx context.Context) []int64 {
	ids, err := tusq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tusq *TelegramUserStateQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tusq.ctx, "Count")
	if err := tusq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tusq, querierCount[*TelegramUserStateQuery](), tusq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tusq *TelegramUserStateQuery) CountX(ctx context.Context) int {
	count, err := tusq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tusq *TelegramUserStateQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tusq.ctx, "Exist")
	switch _, err := tusq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tusq *TelegramUserStateQuery) ExistX(ctx context.Context) bool {
	exist, err := tusq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TelegramUserStateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tusq *TelegramUserStateQuery) Clone() *TelegramUserStateQuery {
	if tusq == nil {
		return nil
	}
	return &TelegramUserStateQuery{
		config:       tusq.config,
		ctx:          tusq.ctx.Clone(),
		order:        append([]telegramuserstate.OrderOption{}, tusq.order...),
		inters:       append([]Interceptor{}, tusq.inters...),
		predicates:   append([]predicate.TelegramUserState{}, tusq.predicates...),
		withChannels: tusq.withChannels.Clone(),
		// clone intermediate query.
		sql:  tusq.sql.Clone(),
		path: tusq.path,
	}
}

// WithChannels tells the query-builder to eager-load the nodes that are connected to
// the "channels" edge. The optional arguments are used to configure the query builder of the edge.
func (tusq *TelegramUserStateQuery) WithChannels(opts ...func(*TelegramChannelStateQuery)) *TelegramUserStateQuery {
	query := (&TelegramChannelStateClient{config: tusq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tusq.withChannels = query
	return tusq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Qts int `json:"qts,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TelegramUserState.Query().
//		GroupBy(telegramuserstate.FieldQts).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tusq *TelegramUserStateQuery) GroupBy(field string, fields ...string) *TelegramUserStateGroupBy {
	tusq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TelegramUserStateGroupBy{build: tusq}
	grbuild.flds = &tusq.ctx.Fields
	grbuild.label = telegramuserstate.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Qts int `json:"qts,omitempty"`
//	}
//
//	client.TelegramUserState.Query().
//		Select(telegramuserstate.FieldQts).
//		Scan(ctx, &v)
func (tusq *TelegramUserStateQuery) Select(fields ...string) *TelegramUserStateSelect {
	tusq.ctx.Fields = append(tusq.ctx.Fields, fields...)
	sbuild := &TelegramUserStateSelect{TelegramUserStateQuery: tusq}
	sbuild.label = telegramuserstate.Label
	sbuild.flds, sbuild.scan = &tusq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TelegramUserStateSelect configured with the given aggregations.
func (tusq *TelegramUserStateQuery) Aggregate(fns ...AggregateFunc) *TelegramUserStateSelect {
	return tusq.Select().Aggregate(fns...)
}

func (tusq *TelegramUserStateQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tusq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tusq); err != nil {
				return err
			}
		}
	}
	for _, f := range tusq.ctx.Fields {
		if !telegramuserstate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tusq.path != nil {
		prev, err := tusq.path(ctx)
		if err != nil {
			return err
		}
		tusq.sql = prev
	}
	return nil
}

func (tusq *TelegramUserStateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TelegramUserState, error) {
	var (
		nodes       = []*TelegramUserState{}
		_spec       = tusq.querySpec()
		loadedTypes = [1]bool{
			tusq.withChannels != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TelegramUserState).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TelegramUserState{config: tusq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tusq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tusq.withChannels; query != nil {
		if err := tusq.loadChannels(ctx, query, nodes,
			func(n *TelegramUserState) { n.Edges.Channels = []*TelegramChannelState{} },
			func(n *TelegramUserState, e *TelegramChannelState) { n.Edges.Channels = append(n.Edges.Channels, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range tusq.withNamedChannels {
		if err := tusq.loadChannels(ctx, query, nodes,
			func(n *TelegramUserState) { n.appendNamedChannels(name) },
			func(n *TelegramUserState, e *TelegramChannelState) { n.appendNamedChannels(name, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tusq *TelegramUserStateQuery) loadChannels(ctx context.Context, query *TelegramChannelStateQuery, nodes []*TelegramUserState, init func(*TelegramUserState), assign func(*TelegramUserState, *TelegramChannelState)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*TelegramUserState)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(telegramchannelstate.FieldUserID)
	}
	query.Where(predicate.TelegramChannelState(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(telegramuserstate.ChannelsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UserID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tusq *TelegramUserStateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tusq.querySpec()
	_spec.Node.Columns = tusq.ctx.Fields
	if len(tusq.ctx.Fields) > 0 {
		_spec.Unique = tusq.ctx.Unique != nil && *tusq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tusq.driver, _spec)
}

func (tusq *TelegramUserStateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(telegramuserstate.Table, telegramuserstate.Columns, sqlgraph.NewFieldSpec(telegramuserstate.FieldID, field.TypeInt64))
	_spec.From = tusq.sql
	if unique := tusq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tusq.path != nil {
		_spec.Unique = true
	}
	if fields := tusq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, telegramuserstate.FieldID)
		for i := range fields {
			if fields[i] != telegramuserstate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tusq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tusq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tusq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tusq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tusq *TelegramUserStateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tusq.driver.Dialect())
	t1 := builder.Table(telegramuserstate.Table)
	columns := tusq.ctx.Fields
	if len(columns) == 0 {
		columns = telegramuserstate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tusq.sql != nil {
		selector = tusq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tusq.ctx.Unique != nil && *tusq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tusq.predicates {
		p(selector)
	}
	for _, p := range tusq.order {
		p(selector)
	}
	if offset := tusq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tusq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedChannels tells the query-builder to eager-load the nodes that are connected to the "channels"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (tusq *TelegramUserStateQuery) WithNamedChannels(name string, opts ...func(*TelegramChannelStateQuery)) *TelegramUserStateQuery {
	query := (&TelegramChannelStateClient{config: tusq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if tusq.withNamedChannels == nil {
		tusq.withNamedChannels = make(map[string]*TelegramChannelStateQuery)
	}
	tusq.withNamedChannels[name] = query
	return tusq
}

// TelegramUserStateGroupBy is the group-by builder for TelegramUserState entities.
type TelegramUserStateGroupBy struct {
	selector
	build *TelegramUserStateQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tusgb *TelegramUserStateGroupBy) Aggregate(fns ...AggregateFunc) *TelegramUserStateGroupBy {
	tusgb.fns = append(tusgb.fns, fns...)
	return tusgb
}

// Scan applies the selector query and scans the result into the given value.
func (tusgb *TelegramUserStateGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tusgb.build.ctx, "GroupBy")
	if err := tusgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TelegramUserStateQuery, *TelegramUserStateGroupBy](ctx, tusgb.build, tusgb, tusgb.build.inters, v)
}

func (tusgb *TelegramUserStateGroupBy) sqlScan(ctx context.Context, root *TelegramUserStateQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tusgb.fns))
	for _, fn := range tusgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tusgb.flds)+len(tusgb.fns))
		for _, f := range *tusgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tusgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tusgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TelegramUserStateSelect is the builder for selecting fields of TelegramUserState entities.
type TelegramUserStateSelect struct {
	*TelegramUserStateQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tuss *TelegramUserStateSelect) Aggregate(fns ...AggregateFunc) *TelegramUserStateSelect {
	tuss.fns = append(tuss.fns, fns...)
	return tuss
}

// Scan applies the selector query and scans the result into the given value.
func (tuss *TelegramUserStateSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tuss.ctx, "Select")
	if err := tuss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TelegramUserStateQuery, *TelegramUserStateSelect](ctx, tuss.TelegramUserStateQuery, tuss, tuss.inters, v)
}

func (tuss *TelegramUserStateSelect) sqlScan(ctx context.Context, root *TelegramUserStateQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tuss.fns))
	for _, fn := range tuss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tuss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tuss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
