// Code generated by ent, DO NOT EDIT.

package message

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the message type in the database.
	Label = "message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// EdgeFrom holds the string denoting the from edge name in mutations.
	EdgeFrom = "from"
	// EdgeMatch holds the string denoting the match edge name in mutations.
	EdgeMatch = "match"
	// Table holds the table name of the message in the database.
	Table = "messages"
	// FromTable is the table that holds the from relation/edge.
	FromTable = "messages"
	// FromInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	FromInverseTable = "users"
	// FromColumn is the table column denoting the from relation/edge.
	FromColumn = "message_from"
	// MatchTable is the table that holds the match relation/edge.
	MatchTable = "messages"
	// MatchInverseTable is the table name for the Match entity.
	// It exists in this package in order to avoid circular dependency with the "match" package.
	MatchInverseTable = "matches"
	// MatchColumn is the table column denoting the match relation/edge.
	MatchColumn = "message_match"
)

// Columns holds all SQL columns for message fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldText,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "messages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"message_from",
	"message_match",
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

// OrderOption defines the ordering options for the Message queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByText orders the results by the text field.
func ByText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldText, opts...).ToFunc()
}

// ByFromField orders the results by from field.
func ByFromField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFromStep(), sql.OrderByField(field, opts...))
	}
}

// ByMatchField orders the results by match field.
func ByMatchField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMatchStep(), sql.OrderByField(field, opts...))
	}
}
func newFromStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FromInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, FromTable, FromColumn),
	)
}
func newMatchStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MatchInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, MatchTable, MatchColumn),
	)
}
