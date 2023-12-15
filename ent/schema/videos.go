package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Videos holds the schema definition for the Videos entity.
type Videos struct {
	ent.Schema
}

// Fields of the Videos.
func (Videos) Fields() []ent.Field {
	return []ent.Field{
		field.String("Desc"),
		field.String("videoLink"),
		field.String("thumb"),
		field.Uint64("likeNum"),
		field.Uint64("commentNum"),
	}
}

// Edges of the Videos.
func (Videos) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", UserSec.Type).Ref("videoId"),
		edge.To("likeId", Likes.Type),
		edge.To("commentId", Comments.Type),
	}
}
