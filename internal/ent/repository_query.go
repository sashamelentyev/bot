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
	"github.com/go-faster/bot/internal/ent/gitcommit"
	"github.com/go-faster/bot/internal/ent/organization"
	"github.com/go-faster/bot/internal/ent/predicate"
	"github.com/go-faster/bot/internal/ent/repository"
)

// RepositoryQuery is the builder for querying Repository entities.
type RepositoryQuery struct {
	config
	ctx              *QueryContext
	order            []repository.OrderOption
	inters           []Interceptor
	predicates       []predicate.Repository
	withOrganization *OrganizationQuery
	withCommits      *GitCommitQuery
	withFKs          bool
	withNamedCommits map[string]*GitCommitQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RepositoryQuery builder.
func (rq *RepositoryQuery) Where(ps ...predicate.Repository) *RepositoryQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *RepositoryQuery) Limit(limit int) *RepositoryQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *RepositoryQuery) Offset(offset int) *RepositoryQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RepositoryQuery) Unique(unique bool) *RepositoryQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RepositoryQuery) Order(o ...repository.OrderOption) *RepositoryQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryOrganization chains the current query on the "organization" edge.
func (rq *RepositoryQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(repository.Table, repository.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, repository.OrganizationTable, repository.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCommits chains the current query on the "commits" edge.
func (rq *RepositoryQuery) QueryCommits() *GitCommitQuery {
	query := (&GitCommitClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(repository.Table, repository.FieldID, selector),
			sqlgraph.To(gitcommit.Table, gitcommit.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, repository.CommitsTable, repository.CommitsColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Repository entity from the query.
// Returns a *NotFoundError when no Repository was found.
func (rq *RepositoryQuery) First(ctx context.Context) (*Repository, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{repository.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RepositoryQuery) FirstX(ctx context.Context) *Repository {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Repository ID from the query.
// Returns a *NotFoundError when no Repository ID was found.
func (rq *RepositoryQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{repository.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RepositoryQuery) FirstIDX(ctx context.Context) int64 {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Repository entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Repository entity is found.
// Returns a *NotFoundError when no Repository entities are found.
func (rq *RepositoryQuery) Only(ctx context.Context) (*Repository, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{repository.Label}
	default:
		return nil, &NotSingularError{repository.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RepositoryQuery) OnlyX(ctx context.Context) *Repository {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Repository ID in the query.
// Returns a *NotSingularError when more than one Repository ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RepositoryQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{repository.Label}
	default:
		err = &NotSingularError{repository.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RepositoryQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Repositories.
func (rq *RepositoryQuery) All(ctx context.Context) ([]*Repository, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Repository, *RepositoryQuery]()
	return withInterceptors[[]*Repository](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *RepositoryQuery) AllX(ctx context.Context) []*Repository {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Repository IDs.
func (rq *RepositoryQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(repository.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RepositoryQuery) IDsX(ctx context.Context) []int64 {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RepositoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RepositoryQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RepositoryQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RepositoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RepositoryQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RepositoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RepositoryQuery) Clone() *RepositoryQuery {
	if rq == nil {
		return nil
	}
	return &RepositoryQuery{
		config:           rq.config,
		ctx:              rq.ctx.Clone(),
		order:            append([]repository.OrderOption{}, rq.order...),
		inters:           append([]Interceptor{}, rq.inters...),
		predicates:       append([]predicate.Repository{}, rq.predicates...),
		withOrganization: rq.withOrganization.Clone(),
		withCommits:      rq.withCommits.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RepositoryQuery) WithOrganization(opts ...func(*OrganizationQuery)) *RepositoryQuery {
	query := (&OrganizationClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withOrganization = query
	return rq
}

// WithCommits tells the query-builder to eager-load the nodes that are connected to
// the "commits" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RepositoryQuery) WithCommits(opts ...func(*GitCommitQuery)) *RepositoryQuery {
	query := (&GitCommitClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withCommits = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Repository.Query().
//		GroupBy(repository.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *RepositoryQuery) GroupBy(field string, fields ...string) *RepositoryGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RepositoryGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = repository.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Repository.Query().
//		Select(repository.FieldName).
//		Scan(ctx, &v)
func (rq *RepositoryQuery) Select(fields ...string) *RepositorySelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &RepositorySelect{RepositoryQuery: rq}
	sbuild.label = repository.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RepositorySelect configured with the given aggregations.
func (rq *RepositoryQuery) Aggregate(fns ...AggregateFunc) *RepositorySelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RepositoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !repository.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RepositoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Repository, error) {
	var (
		nodes       = []*Repository{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [2]bool{
			rq.withOrganization != nil,
			rq.withCommits != nil,
		}
	)
	if rq.withOrganization != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, repository.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Repository).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Repository{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withOrganization; query != nil {
		if err := rq.loadOrganization(ctx, query, nodes, nil,
			func(n *Repository, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withCommits; query != nil {
		if err := rq.loadCommits(ctx, query, nodes,
			func(n *Repository) { n.Edges.Commits = []*GitCommit{} },
			func(n *Repository, e *GitCommit) { n.Edges.Commits = append(n.Edges.Commits, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range rq.withNamedCommits {
		if err := rq.loadCommits(ctx, query, nodes,
			func(n *Repository) { n.appendNamedCommits(name) },
			func(n *Repository, e *GitCommit) { n.appendNamedCommits(name, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RepositoryQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*Repository, init func(*Repository), assign func(*Repository, *Organization)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Repository)
	for i := range nodes {
		if nodes[i].organization_repositories == nil {
			continue
		}
		fk := *nodes[i].organization_repositories
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "organization_repositories" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *RepositoryQuery) loadCommits(ctx context.Context, query *GitCommitQuery, nodes []*Repository, init func(*Repository), assign func(*Repository, *GitCommit)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Repository)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.GitCommit(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(repository.CommitsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.repository_commits
		if fk == nil {
			return fmt.Errorf(`foreign-key "repository_commits" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "repository_commits" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rq *RepositoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RepositoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(repository.Table, repository.Columns, sqlgraph.NewFieldSpec(repository.FieldID, field.TypeInt64))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, repository.FieldID)
		for i := range fields {
			if fields[i] != repository.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RepositoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(repository.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = repository.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedCommits tells the query-builder to eager-load the nodes that are connected to the "commits"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rq *RepositoryQuery) WithNamedCommits(name string, opts ...func(*GitCommitQuery)) *RepositoryQuery {
	query := (&GitCommitClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if rq.withNamedCommits == nil {
		rq.withNamedCommits = make(map[string]*GitCommitQuery)
	}
	rq.withNamedCommits[name] = query
	return rq
}

// RepositoryGroupBy is the group-by builder for Repository entities.
type RepositoryGroupBy struct {
	selector
	build *RepositoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RepositoryGroupBy) Aggregate(fns ...AggregateFunc) *RepositoryGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RepositoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RepositoryQuery, *RepositoryGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RepositoryGroupBy) sqlScan(ctx context.Context, root *RepositoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RepositorySelect is the builder for selecting fields of Repository entities.
type RepositorySelect struct {
	*RepositoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RepositorySelect) Aggregate(fns ...AggregateFunc) *RepositorySelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RepositorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RepositoryQuery, *RepositorySelect](ctx, rs.RepositoryQuery, rs, rs.inters, v)
}

func (rs *RepositorySelect) sqlScan(ctx context.Context, root *RepositoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
