package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"itamconnect/internal/domain/valueobject"
)

type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("login").Unique(),
		field.String("hashed_password"),
		field.String("email").GoType(valueobject.Email("")),
		field.String("bio"),
		field.String("education_info"),
		field.String("phone").Optional(),
		field.String("telegram").Optional(),
		field.String("other_contacts").Optional(),
		field.Enum("type").GoType(valueobject.UserType("")),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("real_experiences", RealExperience.Type),
		edge.To("menti", Menti.Type).Unique(),
		edge.To("mentor", Mentor.Type).Unique(),
		edge.To("skills", Skill.Type).Required(),
	}
}
