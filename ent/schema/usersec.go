package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserSec holds the schema definition for the UserSec entity.
type UserSec struct {
	ent.Schema
}

// Fields of the UserSec.
func (UserSec) Fields() []ent.Field {
	return []ent.Field{
		field.String("password"),
		field.String("email"),
		field.String("address").Default("0x0000000000000000000000000000000000000000"),
	}
}

// Edges of the UserSec.
func (UserSec) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profile", UserProfile.Type).Ref("UserSecure").
			Unique(),
		edge.To("videoId", Videos.Type),
		edge.To("commentId", Comments.Type),
		edge.To("likeId", Likes.Type),
	}
}
