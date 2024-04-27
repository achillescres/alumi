package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Menti holds the schema definition for the Menti entity.
type Menti struct {
	ent.Schema
}

// Fields of the Menti.
func (Menti) Fields() []ent.Field {
	return nil
}

// Edges of the Menti.
func (Menti) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("menti").Required(),
		// Matches
		edge.From("mentors", Mentor.Type).
			Ref("mentis").
			Through("matches", Match.Type),
	}
}
