// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server04/ent/comments"
	"server04/ent/likes"
	"server04/ent/userprofile"
	"server04/ent/usersec"
	"server04/ent/videos"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserSecCreate is the builder for creating a UserSec entity.
type UserSecCreate struct {
	config
	mutation *UserSecMutation
	hooks    []Hook
}

// SetPassword sets the "password" field.
func (usc *UserSecCreate) SetPassword(s string) *UserSecCreate {
	usc.mutation.SetPassword(s)
	return usc
}

// SetEmail sets the "email" field.
func (usc *UserSecCreate) SetEmail(s string) *UserSecCreate {
	usc.mutation.SetEmail(s)
	return usc
}

// SetAddress sets the "address" field.
func (usc *UserSecCreate) SetAddress(s string) *UserSecCreate {
	usc.mutation.SetAddress(s)
	return usc
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (usc *UserSecCreate) SetNillableAddress(s *string) *UserSecCreate {
	if s != nil {
		usc.SetAddress(*s)
	}
	return usc
}

// SetProfileID sets the "profile" edge to the UserProfile entity by ID.
func (usc *UserSecCreate) SetProfileID(id int) *UserSecCreate {
	usc.mutation.SetProfileID(id)
	return usc
}

// SetNillableProfileID sets the "profile" edge to the UserProfile entity by ID if the given value is not nil.
func (usc *UserSecCreate) SetNillableProfileID(id *int) *UserSecCreate {
	if id != nil {
		usc = usc.SetProfileID(*id)
	}
	return usc
}

// SetProfile sets the "profile" edge to the UserProfile entity.
func (usc *UserSecCreate) SetProfile(u *UserProfile) *UserSecCreate {
	return usc.SetProfileID(u.ID)
}

// AddVideoIdIDs adds the "videoId" edge to the Videos entity by IDs.
func (usc *UserSecCreate) AddVideoIdIDs(ids ...int) *UserSecCreate {
	usc.mutation.AddVideoIdIDs(ids...)
	return usc
}

// AddVideoId adds the "videoId" edges to the Videos entity.
func (usc *UserSecCreate) AddVideoId(v ...*Videos) *UserSecCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return usc.AddVideoIdIDs(ids...)
}

// AddCommentIdIDs adds the "commentId" edge to the Comments entity by IDs.
func (usc *UserSecCreate) AddCommentIdIDs(ids ...int) *UserSecCreate {
	usc.mutation.AddCommentIdIDs(ids...)
	return usc
}

// AddCommentId adds the "commentId" edges to the Comments entity.
func (usc *UserSecCreate) AddCommentId(c ...*Comments) *UserSecCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return usc.AddCommentIdIDs(ids...)
}

// AddLikeIdIDs adds the "likeId" edge to the Likes entity by IDs.
func (usc *UserSecCreate) AddLikeIdIDs(ids ...int) *UserSecCreate {
	usc.mutation.AddLikeIdIDs(ids...)
	return usc
}

// AddLikeId adds the "likeId" edges to the Likes entity.
func (usc *UserSecCreate) AddLikeId(l ...*Likes) *UserSecCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return usc.AddLikeIdIDs(ids...)
}

// Mutation returns the UserSecMutation object of the builder.
func (usc *UserSecCreate) Mutation() *UserSecMutation {
	return usc.mutation
}

// Save creates the UserSec in the database.
func (usc *UserSecCreate) Save(ctx context.Context) (*UserSec, error) {
	usc.defaults()
	return withHooks(ctx, usc.sqlSave, usc.mutation, usc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (usc *UserSecCreate) SaveX(ctx context.Context) *UserSec {
	v, err := usc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (usc *UserSecCreate) Exec(ctx context.Context) error {
	_, err := usc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usc *UserSecCreate) ExecX(ctx context.Context) {
	if err := usc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usc *UserSecCreate) defaults() {
	if _, ok := usc.mutation.Address(); !ok {
		v := usersec.DefaultAddress
		usc.mutation.SetAddress(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usc *UserSecCreate) check() error {
	if _, ok := usc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "UserSec.password"`)}
	}
	if _, ok := usc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "UserSec.email"`)}
	}
	if _, ok := usc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "UserSec.address"`)}
	}
	return nil
}

func (usc *UserSecCreate) sqlSave(ctx context.Context) (*UserSec, error) {
	if err := usc.check(); err != nil {
		return nil, err
	}
	_node, _spec := usc.createSpec()
	if err := sqlgraph.CreateNode(ctx, usc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	usc.mutation.id = &_node.ID
	usc.mutation.done = true
	return _node, nil
}

func (usc *UserSecCreate) createSpec() (*UserSec, *sqlgraph.CreateSpec) {
	var (
		_node = &UserSec{config: usc.config}
		_spec = sqlgraph.NewCreateSpec(usersec.Table, sqlgraph.NewFieldSpec(usersec.FieldID, field.TypeInt))
	)
	if value, ok := usc.mutation.Password(); ok {
		_spec.SetField(usersec.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := usc.mutation.Email(); ok {
		_spec.SetField(usersec.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := usc.mutation.Address(); ok {
		_spec.SetField(usersec.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if nodes := usc.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usersec.ProfileTable,
			Columns: []string{usersec.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userprofile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_profile_user_secure = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := usc.mutation.VideoIdIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usersec.VideoIdTable,
			Columns: []string{usersec.VideoIdColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(videos.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := usc.mutation.CommentIdIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   usersec.CommentIdTable,
			Columns: usersec.CommentIdPrimaryKey,
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
	if nodes := usc.mutation.LikeIdIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   usersec.LikeIdTable,
			Columns: usersec.LikeIdPrimaryKey,
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
	return _node, _spec
}

// UserSecCreateBulk is the builder for creating many UserSec entities in bulk.
type UserSecCreateBulk struct {
	config
	err      error
	builders []*UserSecCreate
}

// Save creates the UserSec entities in the database.
func (uscb *UserSecCreateBulk) Save(ctx context.Context) ([]*UserSec, error) {
	if uscb.err != nil {
		return nil, uscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uscb.builders))
	nodes := make([]*UserSec, len(uscb.builders))
	mutators := make([]Mutator, len(uscb.builders))
	for i := range uscb.builders {
		func(i int, root context.Context) {
			builder := uscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserSecMutation)
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
					_, err = mutators[i+1].Mutate(root, uscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uscb *UserSecCreateBulk) SaveX(ctx context.Context) []*UserSec {
	v, err := uscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uscb *UserSecCreateBulk) Exec(ctx context.Context) error {
	_, err := uscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uscb *UserSecCreateBulk) ExecX(ctx context.Context) {
	if err := uscb.Exec(ctx); err != nil {
		panic(err)
	}
}
