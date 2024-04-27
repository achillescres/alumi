package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"itamconnect/internal/domain/valueobject"
	"time"
)

// Match holds the edge schema definition of the Match relationship.
type Match struct {
	ent.Schema
}

// Fields of the Match.
func (Match) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").GoType(valueobject.MatchStatus("")),
		field.Time("created_at").
			Default(time.Now),
		field.Int("mentor_id"),
		field.Int("menti_id"),
	}
}

// Edges of the Match.
func (Match) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mentor", Mentor.Type).
			Required().
			Unique().
			Field("mentor_id"),
		edge.To("menti", Menti.Type).
			Required().
			Unique().
			Field("menti_id"),
		edge.From("messages", Message.Type).Ref("match"),
	}
}
