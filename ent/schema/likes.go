package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Likes holds the schema definition for the Likes entity.
type Likes struct {
	ent.Schema
}

// Fields of the Likes.
func (Likes) Fields() []ent.Field {
	return []ent.Field{
		field.String("commentStr"),
	}
}

// Edges of the Likes.
func (Likes) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("videos", Videos.Type).Ref("likeId"),
		edge.From("user", UserSec.Type).
			Ref("likeId"),
	}
}
