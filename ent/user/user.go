// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"itamconnect/internal/domain/valueobject"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLogin holds the string denoting the login field in the database.
	FieldLogin = "login"
	// FieldHashedPassword holds the string denoting the hashed_password field in the database.
	FieldHashedPassword = "hashed_password"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldBio holds the string denoting the bio field in the database.
	FieldBio = "bio"
	// FieldEducationInfo holds the string denoting the education_info field in the database.
	FieldEducationInfo = "education_info"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldTelegram holds the string denoting the telegram field in the database.
	FieldTelegram = "telegram"
	// FieldOtherContacts holds the string denoting the other_contacts field in the database.
	FieldOtherContacts = "other_contacts"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeRealExperiences holds the string denoting the real_experiences edge name in mutations.
	EdgeRealExperiences = "real_experiences"
	// EdgeMenti holds the string denoting the menti edge name in mutations.
	EdgeMenti = "menti"
	// EdgeMentor holds the string denoting the mentor edge name in mutations.
	EdgeMentor = "mentor"
	// EdgeSkills holds the string denoting the skills edge name in mutations.
	EdgeSkills = "skills"
	// Table holds the table name of the user in the database.
	Table = "users"
	// RealExperiencesTable is the table that holds the real_experiences relation/edge.
	RealExperiencesTable = "real_experiences"
	// RealExperiencesInverseTable is the table name for the RealExperience entity.
	// It exists in this package in order to avoid circular dependency with the "realexperience" package.
	RealExperiencesInverseTable = "real_experiences"
	// RealExperiencesColumn is the table column denoting the real_experiences relation/edge.
	RealExperiencesColumn = "user_real_experiences"
	// MentiTable is the table that holds the menti relation/edge.
	MentiTable = "users"
	// MentiInverseTable is the table name for the Menti entity.
	// It exists in this package in order to avoid circular dependency with the "menti" package.
	MentiInverseTable = "mentis"
	// MentiColumn is the table column denoting the menti relation/edge.
	MentiColumn = "user_menti"
	// MentorTable is the table that holds the mentor relation/edge.
	MentorTable = "users"
	// MentorInverseTable is the table name for the Mentor entity.
	// It exists in this package in order to avoid circular dependency with the "mentor" package.
	MentorInverseTable = "mentors"
	// MentorColumn is the table column denoting the mentor relation/edge.
	MentorColumn = "user_mentor"
	// SkillsTable is the table that holds the skills relation/edge.
	SkillsTable = "skills"
	// SkillsInverseTable is the table name for the Skill entity.
	// It exists in this package in order to avoid circular dependency with the "skill" package.
	SkillsInverseTable = "skills"
	// SkillsColumn is the table column denoting the skills relation/edge.
	SkillsColumn = "user_skills"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldLogin,
	FieldHashedPassword,
	FieldEmail,
	FieldBio,
	FieldEducationInfo,
	FieldPhone,
	FieldTelegram,
	FieldOtherContacts,
	FieldType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_menti",
	"user_mentor",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type valueobject.UserType) error {
	switch _type {
	case "menti", "mentor":
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByLogin orders the results by the login field.
func ByLogin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogin, opts...).ToFunc()
}

// ByHashedPassword orders the results by the hashed_password field.
func ByHashedPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHashedPassword, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByBio orders the results by the bio field.
func ByBio(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBio, opts...).ToFunc()
}

// ByEducationInfo orders the results by the education_info field.
func ByEducationInfo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEducationInfo, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByTelegram orders the results by the telegram field.
func ByTelegram(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTelegram, opts...).ToFunc()
}

// ByOtherContacts orders the results by the other_contacts field.
func ByOtherContacts(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOtherContacts, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByRealExperiencesCount orders the results by real_experiences count.
func ByRealExperiencesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRealExperiencesStep(), opts...)
	}
}

// ByRealExperiences orders the results by real_experiences terms.
func ByRealExperiences(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRealExperiencesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMentiField orders the results by menti field.
func ByMentiField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMentiStep(), sql.OrderByField(field, opts...))
	}
}

// ByMentorField orders the results by mentor field.
func ByMentorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMentorStep(), sql.OrderByField(field, opts...))
	}
}

// BySkillsCount orders the results by skills count.
func BySkillsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSkillsStep(), opts...)
	}
}

// BySkills orders the results by skills terms.
func BySkills(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSkillsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRealExperiencesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RealExperiencesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RealExperiencesTable, RealExperiencesColumn),
	)
}
func newMentiStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MentiInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, MentiTable, MentiColumn),
	)
}
func newMentorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MentorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, MentorTable, MentorColumn),
	)
}
func newSkillsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SkillsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SkillsTable, SkillsColumn),
	)
}
