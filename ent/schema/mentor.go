package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Mentor holds the schema definition for the Mentor entity.
type Mentor struct {
	ent.Schema
}

// Fields of the Mentor.
func (Mentor) Fields() []ent.Field {
	return []ent.Field{
		field.String("search_for"),
	}
}

// Edges of the Mentor.
func (Mentor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("mentor").Required(),
		// Matches
		edge.To("mentis", Menti.Type).Through("matches", Match.Type),
	}
}
