// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server04/ent/comments"
	"server04/ent/predicate"
	"server04/ent/usersec"
	"server04/ent/videos"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommentsUpdate is the builder for updating Comments entities.
type CommentsUpdate struct {
	config
	hooks    []Hook
	mutation *CommentsMutation
}

// Where appends a list predicates to the CommentsUpdate builder.
func (cu *CommentsUpdate) Where(ps ...predicate.Comments) *CommentsUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// AddVideoIdIDs adds the "videoId" edge to the Videos entity by IDs.
func (cu *CommentsUpdate) AddVideoIdIDs(ids ...int) *CommentsUpdate {
	cu.mutation.AddVideoIdIDs(ids...)
	return cu
}

// AddVideoId adds the "videoId" edges to the Videos entity.
func (cu *CommentsUpdate) AddVideoId(v ...*Videos) *CommentsUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cu.AddVideoIdIDs(ids...)
}

// AddUserIdIDs adds the "userId" edge to the UserSec entity by IDs.
func (cu *CommentsUpdate) AddUserIdIDs(ids ...int) *CommentsUpdate {
	cu.mutation.AddUserIdIDs(ids...)
	return cu
}

// AddUserId adds the "userId" edges to the UserSec entity.
func (cu *CommentsUpdate) AddUserId(u ...*UserSec) *CommentsUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.AddUserIdIDs(ids...)
}

// Mutation returns the CommentsMutation object of the builder.
func (cu *CommentsUpdate) Mutation() *CommentsMutation {
	return cu.mutation
}

// ClearVideoId clears all "videoId" edges to the Videos entity.
func (cu *CommentsUpdate) ClearVideoId() *CommentsUpdate {
	cu.mutation.ClearVideoId()
	return cu
}

// RemoveVideoIdIDs removes the "videoId" edge to Videos entities by IDs.
func (cu *CommentsUpdate) RemoveVideoIdIDs(ids ...int) *CommentsUpdate {
	cu.mutation.RemoveVideoIdIDs(ids...)
	return cu
}

// RemoveVideoId removes "videoId" edges to Videos entities.
func (cu *CommentsUpdate) RemoveVideoId(v ...*Videos) *CommentsUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cu.RemoveVideoIdIDs(ids...)
}

// ClearUserId clears all "userId" edges to the UserSec entity.
func (cu *CommentsUpdate) ClearUserId() *CommentsUpdate {
	cu.mutation.ClearUserId()
	return cu
}

// RemoveUserIdIDs removes the "userId" edge to UserSec entities by IDs.
func (cu *CommentsUpdate) RemoveUserIdIDs(ids ...int) *CommentsUpdate {
	cu.mutation.RemoveUserIdIDs(ids...)
	return cu
}

// RemoveUserId removes "userId" edges to UserSec entities.
func (cu *CommentsUpdate) RemoveUserId(u ...*UserSec) *CommentsUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.RemoveUserIdIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommentsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentsUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentsUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentsUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CommentsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(comments.Table, comments.Columns, sqlgraph.NewFieldSpec(comments.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cu.mutation.VideoIdCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comments.VideoIdTable,
			Columns: comments.VideoIdPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(videos.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedVideoIdIDs(); len(nodes) > 0 && !cu.mutation.VideoIdCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comments.VideoIdTable,
			Columns: comments.VideoIdPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(videos.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.VideoIdIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comments.VideoIdTable,
			Columns: comments.VideoIdPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(videos.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.UserIdCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comments.UserIdTable,
			Columns: comments.UserIdPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersec.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedUserIdIDs(); len(nodes) > 0 && !cu.mutation.UserIdCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comments.UserIdTable,
			Columns: comments.UserIdPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersec.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UserIdIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comments.UserIdTable,
			Columns: comments.UserIdPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersec.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comments.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CommentsUpdateOne is the builder for updating a single Comments entity.
type CommentsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentsMutation
}

// AddVideoIdIDs adds the "videoId" edge to the Videos entity by IDs.
func (cuo *CommentsUpdateOne) AddVideoIdIDs(ids ...int) *CommentsUpdateOne {
	cuo.mutation.AddVideoIdIDs(ids...)
	return cuo
}

// AddVideoId adds the "videoId" edges to the Videos entity.
func (cuo *CommentsUpdateOne) AddVideoId(v ...*Videos) *CommentsUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cuo.AddVideoIdIDs(ids...)
}

// AddUserIdIDs adds the "userId" edge to the UserSec entity by IDs.
func (cuo *CommentsUpdateOne) AddUserIdIDs(ids ...int) *CommentsUpdateOne {
	cuo.mutation.AddUserIdIDs(ids...)
	return cuo
}

// AddUserId adds the "userId" edges to the UserSec entity.
func (cuo *CommentsUpdateOne) AddUserId(u ...*UserSec) *CommentsUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.AddUserIdIDs(ids...)
}

// Mutation returns the CommentsMutation object of the builder.
func (cuo *CommentsUpdateOne) Mutation() *CommentsMutation {
	return cuo.mutation
}

// ClearVideoId clears all "videoId" edges to the Videos entity.
func (cuo *CommentsUpdateOne) ClearVideoId() *CommentsUpdateOne {
	cuo.mutation.ClearVideoId()
	return cuo
}

// RemoveVideoIdIDs removes the "videoId" edge to Videos entities by IDs.
func (cuo *CommentsUpdateOne) RemoveVideoIdIDs(ids ...int) *CommentsUpdateOne {
	cuo.mutation.RemoveVideoIdIDs(ids...)
	return cuo
}

// RemoveVideoId removes "videoId" edges to Videos entities.
func (cuo *CommentsUpdateOne) RemoveVideoId(v ...*Videos) *CommentsUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cuo.RemoveVideoIdIDs(ids...)
}

// ClearUserId clears all "userId" edges to the UserSec entity.
func (cuo *CommentsUpdateOne) ClearUserId() *CommentsUpdateOne {
	cuo.mutation.ClearUserId()
	return cuo
}

// RemoveUserIdIDs removes the "userId" edge to UserSec entities by IDs.
func (cuo *CommentsUpdateOne) RemoveUserIdIDs(ids ...int) *CommentsUpdateOne {
	cuo.mutation.RemoveUserIdIDs(ids...)
	return cuo
}

// RemoveUserId removes "userId" edges to UserSec entities.
func (cuo *CommentsUpdateOne) RemoveUserId(u ...*UserSec) *CommentsUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.RemoveUserIdIDs(ids...)
}

// Where appends a list predicates to the CommentsUpdate builder.
func (cuo *CommentsUpdateOne) Where(ps ...predicate.Comments) *CommentsUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommentsUpdateOne) Select(field string, fields ...string) *CommentsUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comments entity.
func (cuo *CommentsUpdateOne) Save(ctx context.Context) (*Comments, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentsUpdateOne) SaveX(ctx context.Context) *Comments {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommentsUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentsUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CommentsUpdateOne) sqlSave(ctx context.Context) (_node *Comments, err error) {
	_spec := sqlgraph.NewUpdateSpec(comments.Table, comments.Columns, sqlgraph.NewFieldSpec(comments.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Comments.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comments.FieldID)
		for _, f := range fields {
			if !comments.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comments.FieldID {
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
	if cuo.mutation.VideoIdCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comments.VideoIdTable,
			Columns: comments.VideoIdPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(videos.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedVideoIdIDs(); len(nodes) > 0 && !cuo.mutation.VideoIdCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comments.VideoIdTable,
			Columns: comments.VideoIdPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(videos.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.VideoIdIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comments.VideoIdTable,
			Columns: comments.VideoIdPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(videos.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.UserIdCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comments.UserIdTable,
			Columns: comments.UserIdPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersec.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedUserIdIDs(); len(nodes) > 0 && !cuo.mutation.UserIdCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comments.UserIdTable,
			Columns: comments.UserIdPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersec.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UserIdIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comments.UserIdTable,
			Columns: comments.UserIdPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersec.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Comments{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comments.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}