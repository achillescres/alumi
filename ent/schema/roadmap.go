package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// RoadMap holds the schema definition for the RoadMap entity.
type RoadMap struct {
	ent.Schema
}

func (RoadMap) Mixins() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the RoadMap.
func (RoadMap) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Strings("tags"),
		field.Text("text"),
	}
}

// Edges of the RoadMap.
func (RoadMap) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("author", Mentor.Type).Unique().Required().Immutable(),
	}
}
