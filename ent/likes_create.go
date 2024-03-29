// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"server04/ent/likes"
	"server04/ent/usersec"
	"server04/ent/videos"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LikesCreate is the builder for creating a Likes entity.
type LikesCreate struct {
	config
	mutation *LikesMutation
	hooks    []Hook
}

// AddVideoIDs adds the "videos" edge to the Videos entity by IDs.
func (lc *LikesCreate) AddVideoIDs(ids ...int) *LikesCreate {
	lc.mutation.AddVideoIDs(ids...)
	return lc
}

// AddVideos adds the "videos" edges to the Videos entity.
func (lc *LikesCreate) AddVideos(v ...*Videos) *LikesCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return lc.AddVideoIDs(ids...)
}

// AddUserIDs adds the "user" edge to the UserSec entity by IDs.
func (lc *LikesCreate) AddUserIDs(ids ...int) *LikesCreate {
	lc.mutation.AddUserIDs(ids...)
	return lc
}

// AddUser adds the "user" edges to the UserSec entity.
func (lc *LikesCreate) AddUser(u ...*UserSec) *LikesCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return lc.AddUserIDs(ids...)
}

// Mutation returns the LikesMutation object of the builder.
func (lc *LikesCreate) Mutation() *LikesMutation {
	return lc.mutation
}

// Save creates the Likes in the database.
func (lc *LikesCreate) Save(ctx context.Context) (*Likes, error) {
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LikesCreate) SaveX(ctx context.Context) *Likes {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LikesCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LikesCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LikesCreate) check() error {
	return nil
}

func (lc *LikesCreate) sqlSave(ctx context.Context) (*Likes, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LikesCreate) createSpec() (*Likes, *sqlgraph.CreateSpec) {
	var (
		_node = &Likes{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(likes.Table, sqlgraph.NewFieldSpec(likes.FieldID, field.TypeInt))
	)
	if nodes := lc.mutation.VideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   likes.VideosTable,
			Columns: likes.VideosPrimaryKey,
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
	if nodes := lc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   likes.UserTable,
			Columns: likes.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersec.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LikesCreateBulk is the builder for creating many Likes entities in bulk.
type LikesCreateBulk struct {
	config
	err      error
	builders []*LikesCreate
}

// Save creates the Likes entities in the database.
func (lcb *LikesCreateBulk) Save(ctx context.Context) ([]*Likes, error) {
	if lcb.err != nil {
		return nil, lcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Likes, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LikesMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LikesCreateBulk) SaveX(ctx context.Context) []*Likes {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LikesCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LikesCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
