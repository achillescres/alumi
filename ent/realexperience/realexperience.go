// Code generated by ent, DO NOT EDIT.

package realexperience

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the realexperience type in the database.
	Label = "real_experience"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldWhere holds the string denoting the where field in the database.
	FieldWhere = "where"
	// FieldAs holds the string denoting the as field in the database.
	FieldAs = "as"
	// FieldHow holds the string denoting the how field in the database.
	FieldHow = "how"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldWhenStarted holds the string denoting the when_started field in the database.
	FieldWhenStarted = "when_started"
	// FieldWhenEnded holds the string denoting the when_ended field in the database.
	FieldWhenEnded = "when_ended"
	// Table holds the table name of the realexperience in the database.
	Table = "real_experiences"
)

// Columns holds all SQL columns for realexperience fields.
var Columns = []string{
	FieldID,
	FieldWhere,
	FieldAs,
	FieldHow,
	FieldDescription,
	FieldWhenStarted,
	FieldWhenEnded,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "real_experiences"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_real_experiences",
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

// OrderOption defines the ordering options for the RealExperience queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByWhere orders the results by the where field.
func ByWhere(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWhere, opts...).ToFunc()
}

// ByAs orders the results by the as field.
func ByAs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAs, opts...).ToFunc()
}

// ByHow orders the results by the how field.
func ByHow(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHow, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByWhenStarted orders the results by the when_started field.
func ByWhenStarted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWhenStarted, opts...).ToFunc()
}

// ByWhenEnded orders the results by the when_ended field.
func ByWhenEnded(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWhenEnded, opts...).ToFunc()
}