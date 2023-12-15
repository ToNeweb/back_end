// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"server04/ent/comments"
	"server04/ent/likes"
	"server04/ent/predicate"
	"server04/ent/usersec"
	"server04/ent/videos"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VideosQuery is the builder for querying Videos entities.
type VideosQuery struct {
	config
	ctx           *QueryContext
	order         []videos.OrderOption
	inters        []Interceptor
	predicates    []predicate.Videos
	withUser      *UserSecQuery
	withLikeId    *LikesQuery
	withCommentId *CommentsQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VideosQuery builder.
func (vq *VideosQuery) Where(ps ...predicate.Videos) *VideosQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit the number of records to be returned by this query.
func (vq *VideosQuery) Limit(limit int) *VideosQuery {
	vq.ctx.Limit = &limit
	return vq
}

// Offset to start from.
func (vq *VideosQuery) Offset(offset int) *VideosQuery {
	vq.ctx.Offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VideosQuery) Unique(unique bool) *VideosQuery {
	vq.ctx.Unique = &unique
	return vq
}

// Order specifies how the records should be ordered.
func (vq *VideosQuery) Order(o ...videos.OrderOption) *VideosQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryUser chains the current query on the "user" edge.
func (vq *VideosQuery) QueryUser() *UserSecQuery {
	query := (&UserSecClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(videos.Table, videos.FieldID, selector),
			sqlgraph.To(usersec.Table, usersec.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, videos.UserTable, videos.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLikeId chains the current query on the "likeId" edge.
func (vq *VideosQuery) QueryLikeId() *LikesQuery {
	query := (&LikesClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(videos.Table, videos.FieldID, selector),
			sqlgraph.To(likes.Table, likes.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, videos.LikeIdTable, videos.LikeIdPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCommentId chains the current query on the "commentId" edge.
func (vq *VideosQuery) QueryCommentId() *CommentsQuery {
	query := (&CommentsClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(videos.Table, videos.FieldID, selector),
			sqlgraph.To(comments.Table, comments.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, videos.CommentIdTable, videos.CommentIdPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Videos entity from the query.
// Returns a *NotFoundError when no Videos was found.
func (vq *VideosQuery) First(ctx context.Context) (*Videos, error) {
	nodes, err := vq.Limit(1).All(setContextOp(ctx, vq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{videos.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VideosQuery) FirstX(ctx context.Context) *Videos {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Videos ID from the query.
// Returns a *NotFoundError when no Videos ID was found.
func (vq *VideosQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vq.Limit(1).IDs(setContextOp(ctx, vq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{videos.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VideosQuery) FirstIDX(ctx context.Context) int {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Videos entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Videos entity is found.
// Returns a *NotFoundError when no Videos entities are found.
func (vq *VideosQuery) Only(ctx context.Context) (*Videos, error) {
	nodes, err := vq.Limit(2).All(setContextOp(ctx, vq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{videos.Label}
	default:
		return nil, &NotSingularError{videos.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VideosQuery) OnlyX(ctx context.Context) *Videos {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Videos ID in the query.
// Returns a *NotSingularError when more than one Videos ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VideosQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vq.Limit(2).IDs(setContextOp(ctx, vq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{videos.Label}
	default:
		err = &NotSingularError{videos.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VideosQuery) OnlyIDX(ctx context.Context) int {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VideosSlice.
func (vq *VideosQuery) All(ctx context.Context) ([]*Videos, error) {
	ctx = setContextOp(ctx, vq.ctx, "All")
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Videos, *VideosQuery]()
	return withInterceptors[[]*Videos](ctx, vq, qr, vq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vq *VideosQuery) AllX(ctx context.Context) []*Videos {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Videos IDs.
func (vq *VideosQuery) IDs(ctx context.Context) (ids []int, err error) {
	if vq.ctx.Unique == nil && vq.path != nil {
		vq.Unique(true)
	}
	ctx = setContextOp(ctx, vq.ctx, "IDs")
	if err = vq.Select(videos.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VideosQuery) IDsX(ctx context.Context) []int {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VideosQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vq.ctx, "Count")
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vq, querierCount[*VideosQuery](), vq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VideosQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VideosQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vq.ctx, "Exist")
	switch _, err := vq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VideosQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VideosQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VideosQuery) Clone() *VideosQuery {
	if vq == nil {
		return nil
	}
	return &VideosQuery{
		config:        vq.config,
		ctx:           vq.ctx.Clone(),
		order:         append([]videos.OrderOption{}, vq.order...),
		inters:        append([]Interceptor{}, vq.inters...),
		predicates:    append([]predicate.Videos{}, vq.predicates...),
		withUser:      vq.withUser.Clone(),
		withLikeId:    vq.withLikeId.Clone(),
		withCommentId: vq.withCommentId.Clone(),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VideosQuery) WithUser(opts ...func(*UserSecQuery)) *VideosQuery {
	query := (&UserSecClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withUser = query
	return vq
}

// WithLikeId tells the query-builder to eager-load the nodes that are connected to
// the "likeId" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VideosQuery) WithLikeId(opts ...func(*LikesQuery)) *VideosQuery {
	query := (&LikesClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withLikeId = query
	return vq
}

// WithCommentId tells the query-builder to eager-load the nodes that are connected to
// the "commentId" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VideosQuery) WithCommentId(opts ...func(*CommentsQuery)) *VideosQuery {
	query := (&CommentsClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withCommentId = query
	return vq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Desc string `json:"Desc,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Videos.Query().
//		GroupBy(videos.FieldDesc).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vq *VideosQuery) GroupBy(field string, fields ...string) *VideosGroupBy {
	vq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VideosGroupBy{build: vq}
	grbuild.flds = &vq.ctx.Fields
	grbuild.label = videos.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Desc string `json:"Desc,omitempty"`
//	}
//
//	client.Videos.Query().
//		Select(videos.FieldDesc).
//		Scan(ctx, &v)
func (vq *VideosQuery) Select(fields ...string) *VideosSelect {
	vq.ctx.Fields = append(vq.ctx.Fields, fields...)
	sbuild := &VideosSelect{VideosQuery: vq}
	sbuild.label = videos.Label
	sbuild.flds, sbuild.scan = &vq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VideosSelect configured with the given aggregations.
func (vq *VideosQuery) Aggregate(fns ...AggregateFunc) *VideosSelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *VideosQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vq); err != nil {
				return err
			}
		}
	}
	for _, f := range vq.ctx.Fields {
		if !videos.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VideosQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Videos, error) {
	var (
		nodes       = []*Videos{}
		withFKs     = vq.withFKs
		_spec       = vq.querySpec()
		loadedTypes = [3]bool{
			vq.withUser != nil,
			vq.withLikeId != nil,
			vq.withCommentId != nil,
		}
	)
	if vq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, videos.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Videos).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Videos{config: vq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vq.withUser; query != nil {
		if err := vq.loadUser(ctx, query, nodes, nil,
			func(n *Videos, e *UserSec) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := vq.withLikeId; query != nil {
		if err := vq.loadLikeId(ctx, query, nodes,
			func(n *Videos) { n.Edges.LikeId = []*Likes{} },
			func(n *Videos, e *Likes) { n.Edges.LikeId = append(n.Edges.LikeId, e) }); err != nil {
			return nil, err
		}
	}
	if query := vq.withCommentId; query != nil {
		if err := vq.loadCommentId(ctx, query, nodes,
			func(n *Videos) { n.Edges.CommentId = []*Comments{} },
			func(n *Videos, e *Comments) { n.Edges.CommentId = append(n.Edges.CommentId, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vq *VideosQuery) loadUser(ctx context.Context, query *UserSecQuery, nodes []*Videos, init func(*Videos), assign func(*Videos, *UserSec)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Videos)
	for i := range nodes {
		if nodes[i].user_sec_video_id == nil {
			continue
		}
		fk := *nodes[i].user_sec_video_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(usersec.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_sec_video_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (vq *VideosQuery) loadLikeId(ctx context.Context, query *LikesQuery, nodes []*Videos, init func(*Videos), assign func(*Videos, *Likes)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Videos)
	nids := make(map[int]map[*Videos]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(videos.LikeIdTable)
		s.Join(joinT).On(s.C(likes.FieldID), joinT.C(videos.LikeIdPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(videos.LikeIdPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(videos.LikeIdPrimaryKey[0]))
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
					nids[inValue] = map[*Videos]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Likes](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "likeId" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (vq *VideosQuery) loadCommentId(ctx context.Context, query *CommentsQuery, nodes []*Videos, init func(*Videos), assign func(*Videos, *Comments)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Videos)
	nids := make(map[int]map[*Videos]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(videos.CommentIdTable)
		s.Join(joinT).On(s.C(comments.FieldID), joinT.C(videos.CommentIdPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(videos.CommentIdPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(videos.CommentIdPrimaryKey[0]))
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
					nids[inValue] = map[*Videos]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Comments](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "commentId" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (vq *VideosQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	_spec.Node.Columns = vq.ctx.Fields
	if len(vq.ctx.Fields) > 0 {
		_spec.Unique = vq.ctx.Unique != nil && *vq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VideosQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(videos.Table, videos.Columns, sqlgraph.NewFieldSpec(videos.FieldID, field.TypeInt))
	_spec.From = vq.sql
	if unique := vq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vq.path != nil {
		_spec.Unique = true
	}
	if fields := vq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, videos.FieldID)
		for i := range fields {
			if fields[i] != videos.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VideosQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(videos.Table)
	columns := vq.ctx.Fields
	if len(columns) == 0 {
		columns = videos.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.ctx.Unique != nil && *vq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VideosGroupBy is the group-by builder for Videos entities.
type VideosGroupBy struct {
	selector
	build *VideosQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VideosGroupBy) Aggregate(fns ...AggregateFunc) *VideosGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the selector query and scans the result into the given value.
func (vgb *VideosGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vgb.build.ctx, "GroupBy")
	if err := vgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideosQuery, *VideosGroupBy](ctx, vgb.build, vgb, vgb.build.inters, v)
}

func (vgb *VideosGroupBy) sqlScan(ctx context.Context, root *VideosQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vgb.flds)+len(vgb.fns))
		for _, f := range *vgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VideosSelect is the builder for selecting fields of Videos entities.
type VideosSelect struct {
	*VideosQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *VideosSelect) Aggregate(fns ...AggregateFunc) *VideosSelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VideosSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vs.ctx, "Select")
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideosQuery, *VideosSelect](ctx, vs.VideosQuery, vs, vs.inters, v)
}

func (vs *VideosSelect) sqlScan(ctx context.Context, root *VideosQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
