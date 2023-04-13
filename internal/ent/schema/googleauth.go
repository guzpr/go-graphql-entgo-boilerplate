package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GoogleAuth holds the schema definition for the GoogleAuth entity.
type GoogleAuth struct {
	ent.Schema
}

func (GoogleAuth) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
		KsuIDMixin{},
	}
}

// Fields of the GoogleAuth.
func (GoogleAuth) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id"),
		field.String("google_id").
			Unique(),
	}
}

// Edges of the GoogleAuth.
func (GoogleAuth) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("google_auth").
			Field("user_id").
			Unique().
			Required(),
	}
}
