package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Comments holds the schema definition for the Comments entity.
type Comments struct {
	ent.Schema
}

// Fields of the Comments.
func (Comments) Fields() []ent.Field {
	return nil
}

// Edges of the Comments.
func (Comments) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("videoId", Videos.Type).Ref("commentId"),
		edge.From("userId", UserSec.Type).Ref("commentId"),
	}
}
