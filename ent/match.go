// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"itamconnect/ent/match"
	"itamconnect/ent/menti"
	"itamconnect/ent/mentor"
	"itamconnect/internal/domain/valueobject"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Match is the model entity for the Match schema.
type Match struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Status holds the value of the "status" field.
	Status valueobject.MatchStatus `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// MentorID holds the value of the "mentor_id" field.
	MentorID int `json:"mentor_id,omitempty"`
	// MentiID holds the value of the "menti_id" field.
	MentiID int `json:"menti_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MatchQuery when eager-loading is set.
	Edges        MatchEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MatchEdges holds the relations/edges for other nodes in the graph.
type MatchEdges struct {
	// Mentor holds the value of the mentor edge.
	Mentor *Mentor `json:"mentor,omitempty"`
	// Menti holds the value of the menti edge.
	Menti *Menti `json:"menti,omitempty"`
	// Messages holds the value of the messages edge.
	Messages []*Message `json:"messages,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// MentorOrErr returns the Mentor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MatchEdges) MentorOrErr() (*Mentor, error) {
	if e.Mentor != nil {
		return e.Mentor, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: mentor.Label}
	}
	return nil, &NotLoadedError{edge: "mentor"}
}

// MentiOrErr returns the Menti value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MatchEdges) MentiOrErr() (*Menti, error) {
	if e.Menti != nil {
		return e.Menti, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: menti.Label}
	}
	return nil, &NotLoadedError{edge: "menti"}
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading.
func (e MatchEdges) MessagesOrErr() ([]*Message, error) {
	if e.loadedTypes[2] {
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Match) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case match.FieldID, match.FieldMentorID, match.FieldMentiID:
			values[i] = new(sql.NullInt64)
		case match.FieldStatus:
			values[i] = new(sql.NullString)
		case match.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Match fields.
func (m *Match) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case match.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case match.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				m.Status = valueobject.MatchStatus(value.String)
			}
		case match.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case match.FieldMentorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mentor_id", values[i])
			} else if value.Valid {
				m.MentorID = int(value.Int64)
			}
		case match.FieldMentiID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field menti_id", values[i])
			} else if value.Valid {
				m.MentiID = int(value.Int64)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Match.
// This includes values selected through modifiers, order, etc.
func (m *Match) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryMentor queries the "mentor" edge of the Match entity.
func (m *Match) QueryMentor() *MentorQuery {
	return NewMatchClient(m.config).QueryMentor(m)
}

// QueryMenti queries the "menti" edge of the Match entity.
func (m *Match) QueryMenti() *MentiQuery {
	return NewMatchClient(m.config).QueryMenti(m)
}

// QueryMessages queries the "messages" edge of the Match entity.
func (m *Match) QueryMessages() *MessageQuery {
	return NewMatchClient(m.config).QueryMessages(m)
}

// Update returns a builder for updating this Match.
// Note that you need to call Match.Unwrap() before calling this method if this Match
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Match) Update() *MatchUpdateOne {
	return NewMatchClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Match entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Match) Unwrap() *Match {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Match is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Match) String() string {
	var builder strings.Builder
	builder.WriteString("Match(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", m.Status))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("mentor_id=")
	builder.WriteString(fmt.Sprintf("%v", m.MentorID))
	builder.WriteString(", ")
	builder.WriteString("menti_id=")
	builder.WriteString(fmt.Sprintf("%v", m.MentiID))
	builder.WriteByte(')')
	return builder.String()
}

// Matches is a parsable slice of Match.
type Matches []*Match
