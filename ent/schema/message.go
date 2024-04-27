package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Immutable(),
		field.String("text"),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("from", User.Type).Unique().Immutable().Required(),
		edge.To("match", Match.Type).Unique().Immutable().Required(),
	}
}

func (Message) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("match", "from").Fields("created_at"),
	}
}
