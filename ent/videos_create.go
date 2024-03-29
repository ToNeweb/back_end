// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server04/ent/comments"
	"server04/ent/likes"
	"server04/ent/usersec"
	"server04/ent/videos"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VideosCreate is the builder for creating a Videos entity.
type VideosCreate struct {
	config
	mutation *VideosMutation
	hooks    []Hook
}

// SetDesc sets the "Desc" field.
func (vc *VideosCreate) SetDesc(s string) *VideosCreate {
	vc.mutation.SetDesc(s)
	return vc
}

// SetVideoLink sets the "videoLink" field.
func (vc *VideosCreate) SetVideoLink(s string) *VideosCreate {
	vc.mutation.SetVideoLink(s)
	return vc
}

// SetThumb sets the "thumb" field.
func (vc *VideosCreate) SetThumb(s string) *VideosCreate {
	vc.mutation.SetThumb(s)
	return vc
}

// SetLikeNum sets the "likeNum" field.
func (vc *VideosCreate) SetLikeNum(u uint64) *VideosCreate {
	vc.mutation.SetLikeNum(u)
	return vc
}

// SetCommentNum sets the "commentNum" field.
func (vc *VideosCreate) SetCommentNum(u uint64) *VideosCreate {
	vc.mutation.SetCommentNum(u)
	return vc
}

// SetUserID sets the "user" edge to the UserSec entity by ID.
func (vc *VideosCreate) SetUserID(id int) *VideosCreate {
	vc.mutation.SetUserID(id)
	return vc
}

// SetNillableUserID sets the "user" edge to the UserSec entity by ID if the given value is not nil.
func (vc *VideosCreate) SetNillableUserID(id *int) *VideosCreate {
	if id != nil {
		vc = vc.SetUserID(*id)
	}
	return vc
}

// SetUser sets the "user" edge to the UserSec entity.
func (vc *VideosCreate) SetUser(u *UserSec) *VideosCreate {
	return vc.SetUserID(u.ID)
}

// AddLikeIdIDs adds the "likeId" edge to the Likes entity by IDs.
func (vc *VideosCreate) AddLikeIdIDs(ids ...int) *VideosCreate {
	vc.mutation.AddLikeIdIDs(ids...)
	return vc
}

// AddLikeId adds the "likeId" edges to the Likes entity.
func (vc *VideosCreate) AddLikeId(l ...*Likes) *VideosCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return vc.AddLikeIdIDs(ids...)
}

// AddCommentIdIDs adds the "commentId" edge to the Comments entity by IDs.
func (vc *VideosCreate) AddCommentIdIDs(ids ...int) *VideosCreate {
	vc.mutation.AddCommentIdIDs(ids...)
	return vc
}

// AddCommentId adds the "commentId" edges to the Comments entity.
func (vc *VideosCreate) AddCommentId(c ...*Comments) *VideosCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return vc.AddCommentIdIDs(ids...)
}

// Mutation returns the VideosMutation object of the builder.
func (vc *VideosCreate) Mutation() *VideosMutation {
	return vc.mutation
}

// Save creates the Videos in the database.
func (vc *VideosCreate) Save(ctx context.Context) (*Videos, error) {
	return withHooks(ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VideosCreate) SaveX(ctx context.Context) *Videos {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VideosCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VideosCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VideosCreate) check() error {
	if _, ok := vc.mutation.Desc(); !ok {
		return &ValidationError{Name: "Desc", err: errors.New(`ent: missing required field "Videos.Desc"`)}
	}
	if _, ok := vc.mutation.VideoLink(); !ok {
		return &ValidationError{Name: "videoLink", err: errors.New(`ent: missing required field "Videos.videoLink"`)}
	}
	if _, ok := vc.mutation.Thumb(); !ok {
		return &ValidationError{Name: "thumb", err: errors.New(`ent: missing required field "Videos.thumb"`)}
	}
	if _, ok := vc.mutation.LikeNum(); !ok {
		return &ValidationError{Name: "likeNum", err: errors.New(`ent: missing required field "Videos.likeNum"`)}
	}
	if _, ok := vc.mutation.CommentNum(); !ok {
		return &ValidationError{Name: "commentNum", err: errors.New(`ent: missing required field "Videos.commentNum"`)}
	}
	return nil
}

func (vc *VideosCreate) sqlSave(ctx context.Context) (*Videos, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VideosCreate) createSpec() (*Videos, *sqlgraph.CreateSpec) {
	var (
		_node = &Videos{config: vc.config}
		_spec = sqlgraph.NewCreateSpec(videos.Table, sqlgraph.NewFieldSpec(videos.FieldID, field.TypeInt))
	)
	if value, ok := vc.mutation.Desc(); ok {
		_spec.SetField(videos.FieldDesc, field.TypeString, value)
		_node.Desc = value
	}
	if value, ok := vc.mutation.VideoLink(); ok {
		_spec.SetField(videos.FieldVideoLink, field.TypeString, value)
		_node.VideoLink = value
	}
	if value, ok := vc.mutation.Thumb(); ok {
		_spec.SetField(videos.FieldThumb, field.TypeString, value)
		_node.Thumb = value
	}
	if value, ok := vc.mutation.LikeNum(); ok {
		_spec.SetField(videos.FieldLikeNum, field.TypeUint64, value)
		_node.LikeNum = value
	}
	if value, ok := vc.mutation.CommentNum(); ok {
		_spec.SetField(videos.FieldCommentNum, field.TypeUint64, value)
		_node.CommentNum = value
	}
	if nodes := vc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   videos.UserTable,
			Columns: []string{videos.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersec.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_sec_video_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.LikeIdIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   videos.LikeIdTable,
			Columns: videos.LikeIdPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(likes.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.CommentIdIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   videos.CommentIdTable,
			Columns: videos.CommentIdPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comments.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VideosCreateBulk is the builder for creating many Videos entities in bulk.
type VideosCreateBulk struct {
	config
	err      error
	builders []*VideosCreate
}

// Save creates the Videos entities in the database.
func (vcb *VideosCreateBulk) Save(ctx context.Context) ([]*Videos, error) {
	if vcb.err != nil {
		return nil, vcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Videos, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VideosMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VideosCreateBulk) SaveX(ctx context.Context) []*Videos {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VideosCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VideosCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
