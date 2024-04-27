package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// RealExperience holds the schema definition for the RealExperience entity.
type RealExperience struct {
	ent.Schema
}

// Fields of the RealExperience.
func (RealExperience) Fields() []ent.Field {
	return []ent.Field{
		field.String("where"),
		field.String("as"),
		field.String("how"),
		field.String("description"),
		field.Time("when_started").Optional(),
		field.Time("when_ended").Optional(),
	}
}
