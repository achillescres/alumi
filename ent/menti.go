// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"itamconnect/ent/menti"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Menti is the model entity for the Menti schema.
type Menti struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MentiQuery when eager-loading is set.
	Edges        MentiEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MentiEdges holds the relations/edges for other nodes in the graph.
type MentiEdges struct {
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// Mentors holds the value of the mentors edge.
	Mentors []*Mentor `json:"mentors,omitempty"`
	// Matches holds the value of the matches edge.
	Matches []*Match `json:"matches,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e MentiEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// MentorsOrErr returns the Mentors value or an error if the edge
// was not loaded in eager-loading.
func (e MentiEdges) MentorsOrErr() ([]*Mentor, error) {
	if e.loadedTypes[1] {
		return e.Mentors, nil
	}
	return nil, &NotLoadedError{edge: "mentors"}
}

// MatchesOrErr returns the Matches value or an error if the edge
// was not loaded in eager-loading.
func (e MentiEdges) MatchesOrErr() ([]*Match, error) {
	if e.loadedTypes[2] {
		return e.Matches, nil
	}
	return nil, &NotLoadedError{edge: "matches"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Menti) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case menti.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Menti fields.
func (m *Menti) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case menti.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Menti.
// This includes values selected through modifiers, order, etc.
func (m *Menti) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Menti entity.
func (m *Menti) QueryUser() *UserQuery {
	return NewMentiClient(m.config).QueryUser(m)
}

// QueryMentors queries the "mentors" edge of the Menti entity.
func (m *Menti) QueryMentors() *MentorQuery {
	return NewMentiClient(m.config).QueryMentors(m)
}

// QueryMatches queries the "matches" edge of the Menti entity.
func (m *Menti) QueryMatches() *MatchQuery {
	return NewMentiClient(m.config).QueryMatches(m)
}

// Update returns a builder for updating this Menti.
// Note that you need to call Menti.Unwrap() before calling this method if this Menti
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Menti) Update() *MentiUpdateOne {
	return NewMentiClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Menti entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Menti) Unwrap() *Menti {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Menti is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Menti) String() string {
	var builder strings.Builder
	builder.WriteString("Menti(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Mentis is a parsable slice of Menti.
type Mentis []*Menti