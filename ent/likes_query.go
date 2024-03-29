// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"server04/ent/likes"
	"server04/ent/predicate"
	"server04/ent/usersec"
	"server04/ent/videos"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LikesQuery is the builder for querying Likes entities.
type LikesQuery struct {
	config
	ctx        *QueryContext
	order      []likes.OrderOption
	inters     []Interceptor
	predicates []predicate.Likes
	withVideos *VideosQuery
	withUser   *UserSecQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LikesQuery builder.
func (lq *LikesQuery) Where(ps ...predicate.Likes) *LikesQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit the number of records to be returned by this query.
func (lq *LikesQuery) Limit(limit int) *LikesQuery {
	lq.ctx.Limit = &limit
	return lq
}

// Offset to start from.
func (lq *LikesQuery) Offset(offset int) *LikesQuery {
	lq.ctx.Offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LikesQuery) Unique(unique bool) *LikesQuery {
	lq.ctx.Unique = &unique
	return lq
}

// Order specifies how the records should be ordered.
func (lq *LikesQuery) Order(o ...likes.OrderOption) *LikesQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// QueryVideos chains the current query on the "videos" edge.
func (lq *LikesQuery) QueryVideos() *VideosQuery {
	query := (&VideosClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(likes.Table, likes.FieldID, selector),
			sqlgraph.To(videos.Table, videos.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, likes.VideosTable, likes.VideosPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (lq *LikesQuery) QueryUser() *UserSecQuery {
	query := (&UserSecClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(likes.Table, likes.FieldID, selector),
			sqlgraph.To(usersec.Table, usersec.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, likes.UserTable, likes.UserPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Likes entity from the query.
// Returns a *NotFoundError when no Likes was found.
func (lq *LikesQuery) First(ctx context.Context) (*Likes, error) {
	nodes, err := lq.Limit(1).All(setContextOp(ctx, lq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{likes.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LikesQuery) FirstX(ctx context.Context) *Likes {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Likes ID from the query.
// Returns a *NotFoundError when no Likes ID was found.
func (lq *LikesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(1).IDs(setContextOp(ctx, lq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{likes.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LikesQuery) FirstIDX(ctx context.Context) int {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Likes entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Likes entity is found.
// Returns a *NotFoundError when no Likes entities are found.
func (lq *LikesQuery) Only(ctx context.Context) (*Likes, error) {
	nodes, err := lq.Limit(2).All(setContextOp(ctx, lq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{likes.Label}
	default:
		return nil, &NotSingularError{likes.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LikesQuery) OnlyX(ctx context.Context) *Likes {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Likes ID in the query.
// Returns a *NotSingularError when more than one Likes ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LikesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(2).IDs(setContextOp(ctx, lq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{likes.Label}
	default:
		err = &NotSingularError{likes.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LikesQuery) OnlyIDX(ctx context.Context) int {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LikesSlice.
func (lq *LikesQuery) All(ctx context.Context) ([]*Likes, error) {
	ctx = setContextOp(ctx, lq.ctx, "All")
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Likes, *LikesQuery]()
	return withInterceptors[[]*Likes](ctx, lq, qr, lq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lq *LikesQuery) AllX(ctx context.Context) []*Likes {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Likes IDs.
func (lq *LikesQuery) IDs(ctx context.Context) (ids []int, err error) {
	if lq.ctx.Unique == nil && lq.path != nil {
		lq.Unique(true)
	}
	ctx = setContextOp(ctx, lq.ctx, "IDs")
	if err = lq.Select(likes.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LikesQuery) IDsX(ctx context.Context) []int {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LikesQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lq.ctx, "Count")
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lq, querierCount[*LikesQuery](), lq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LikesQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LikesQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lq.ctx, "Exist")
	switch _, err := lq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LikesQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LikesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LikesQuery) Clone() *LikesQuery {
	if lq == nil {
		return nil
	}
	return &LikesQuery{
		config:     lq.config,
		ctx:        lq.ctx.Clone(),
		order:      append([]likes.OrderOption{}, lq.order...),
		inters:     append([]Interceptor{}, lq.inters...),
		predicates: append([]predicate.Likes{}, lq.predicates...),
		withVideos: lq.withVideos.Clone(),
		withUser:   lq.withUser.Clone(),
		// clone intermediate query.
		sql:  lq.sql.Clone(),
		path: lq.path,
	}
}

// WithVideos tells the query-builder to eager-load the nodes that are connected to
// the "videos" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LikesQuery) WithVideos(opts ...func(*VideosQuery)) *LikesQuery {
	query := (&VideosClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withVideos = query
	return lq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LikesQuery) WithUser(opts ...func(*UserSecQuery)) *LikesQuery {
	query := (&UserSecClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withUser = query
	return lq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (lq *LikesQuery) GroupBy(field string, fields ...string) *LikesGroupBy {
	lq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LikesGroupBy{build: lq}
	grbuild.flds = &lq.ctx.Fields
	grbuild.label = likes.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (lq *LikesQuery) Select(fields ...string) *LikesSelect {
	lq.ctx.Fields = append(lq.ctx.Fields, fields...)
	sbuild := &LikesSelect{LikesQuery: lq}
	sbuild.label = likes.Label
	sbuild.flds, sbuild.scan = &lq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LikesSelect configured with the given aggregations.
func (lq *LikesQuery) Aggregate(fns ...AggregateFunc) *LikesSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LikesQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lq); err != nil {
				return err
			}
		}
	}
	for _, f := range lq.ctx.Fields {
		if !likes.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LikesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Likes, error) {
	var (
		nodes       = []*Likes{}
		_spec       = lq.querySpec()
		loadedTypes = [2]bool{
			lq.withVideos != nil,
			lq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Likes).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Likes{config: lq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lq.withVideos; query != nil {
		if err := lq.loadVideos(ctx, query, nodes,
			func(n *Likes) { n.Edges.Videos = []*Videos{} },
			func(n *Likes, e *Videos) { n.Edges.Videos = append(n.Edges.Videos, e) }); err != nil {
			return nil, err
		}
	}
	if query := lq.withUser; query != nil {
		if err := lq.loadUser(ctx, query, nodes,
			func(n *Likes) { n.Edges.User = []*UserSec{} },
			func(n *Likes, e *UserSec) { n.Edges.User = append(n.Edges.User, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lq *LikesQuery) loadVideos(ctx context.Context, query *VideosQuery, nodes []*Likes, init func(*Likes), assign func(*Likes, *Videos)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Likes)
	nids := make(map[int]map[*Likes]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(likes.VideosTable)
		s.Join(joinT).On(s.C(videos.FieldID), joinT.C(likes.VideosPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(likes.VideosPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(likes.VideosPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Likes]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Videos](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "videos" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (lq *LikesQuery) loadUser(ctx context.Context, query *UserSecQuery, nodes []*Likes, init func(*Likes), assign func(*Likes, *UserSec)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Likes)
	nids := make(map[int]map[*Likes]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(likes.UserTable)
		s.Join(joinT).On(s.C(usersec.FieldID), joinT.C(likes.UserPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(likes.UserPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(likes.UserPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Likes]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*UserSec](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "user" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (lq *LikesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	_spec.Node.Columns = lq.ctx.Fields
	if len(lq.ctx.Fields) > 0 {
		_spec.Unique = lq.ctx.Unique != nil && *lq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LikesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(likes.Table, likes.Columns, sqlgraph.NewFieldSpec(likes.FieldID, field.TypeInt))
	_spec.From = lq.sql
	if unique := lq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lq.path != nil {
		_spec.Unique = true
	}
	if fields := lq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, likes.FieldID)
		for i := range fields {
			if fields[i] != likes.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LikesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(likes.Table)
	columns := lq.ctx.Fields
	if len(columns) == 0 {
		columns = likes.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.ctx.Unique != nil && *lq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LikesGroupBy is the group-by builder for Likes entities.
type LikesGroupBy struct {
	selector
	build *LikesQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LikesGroupBy) Aggregate(fns ...AggregateFunc) *LikesGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgb *LikesGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgb.build.ctx, "GroupBy")
	if err := lgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LikesQuery, *LikesGroupBy](ctx, lgb.build, lgb, lgb.build.inters, v)
}

func (lgb *LikesGroupBy) sqlScan(ctx context.Context, root *LikesQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lgb.flds)+len(lgb.fns))
		for _, f := range *lgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LikesSelect is the builder for selecting fields of Likes entities.
type LikesSelect struct {
	*LikesQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LikesSelect) Aggregate(fns ...AggregateFunc) *LikesSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LikesSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ls.ctx, "Select")
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LikesQuery, *LikesSelect](ctx, ls.LikesQuery, ls, ls.inters, v)
}

func (ls *LikesSelect) sqlScan(ctx context.Context, root *LikesQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
