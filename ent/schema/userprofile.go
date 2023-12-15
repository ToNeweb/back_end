package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserProfile holds the schema definition for the UserProfile entity.
type UserProfile struct {
	ent.Schema
}

// Fields of the UserProfile.
func (UserProfile) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.String("profileImage"),
	}
}

// Edges of the UserProfile.
func (UserProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("UserSecure", UserSec.Type),
	}
}
