// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"itamconnect/ent/mentor"
	"itamconnect/ent/roadmap"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// RoadMap is the model entity for the RoadMap schema.
type RoadMap struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Tags holds the value of the "tags" field.
	Tags []string `json:"tags,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoadMapQuery when eager-loading is set.
	Edges           RoadMapEdges `json:"edges"`
	road_map_author *int
	selectValues    sql.SelectValues
}

// RoadMapEdges holds the relations/edges for other nodes in the graph.
type RoadMapEdges struct {
	// Author holds the value of the author edge.
	Author *Mentor `json:"author,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoadMapEdges) AuthorOrErr() (*Mentor, error) {
	if e.Author != nil {
		return e.Author, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: mentor.Label}
	}
	return nil, &NotLoadedError{edge: "author"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoadMap) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case roadmap.FieldTags:
			values[i] = new([]byte)
		case roadmap.FieldID:
			values[i] = new(sql.NullInt64)
		case roadmap.FieldName, roadmap.FieldText:
			values[i] = new(sql.NullString)
		case roadmap.ForeignKeys[0]: // road_map_author
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoadMap fields.
func (rm *RoadMap) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case roadmap.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rm.ID = int(value.Int64)
		case roadmap.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				rm.Name = value.String
			}
		case roadmap.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &rm.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case roadmap.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				rm.Text = value.String
			}
		case roadmap.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field road_map_author", value)
			} else if value.Valid {
				rm.road_map_author = new(int)
				*rm.road_map_author = int(value.Int64)
			}
		default:
			rm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RoadMap.
// This includes values selected through modifiers, order, etc.
func (rm *RoadMap) Value(name string) (ent.Value, error) {
	return rm.selectValues.Get(name)
}

// QueryAuthor queries the "author" edge of the RoadMap entity.
func (rm *RoadMap) QueryAuthor() *MentorQuery {
	return NewRoadMapClient(rm.config).QueryAuthor(rm)
}

// Update returns a builder for updating this RoadMap.
// Note that you need to call RoadMap.Unwrap() before calling this method if this RoadMap
// was returned from a transaction, and the transaction was committed or rolled back.
func (rm *RoadMap) Update() *RoadMapUpdateOne {
	return NewRoadMapClient(rm.config).UpdateOne(rm)
}

// Unwrap unwraps the RoadMap entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rm *RoadMap) Unwrap() *RoadMap {
	_tx, ok := rm.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoadMap is not a transactional entity")
	}
	rm.config.driver = _tx.drv
	return rm
}

// String implements the fmt.Stringer.
func (rm *RoadMap) String() string {
	var builder strings.Builder
	builder.WriteString("RoadMap(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rm.ID))
	builder.WriteString("name=")
	builder.WriteString(rm.Name)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", rm.Tags))
	builder.WriteString(", ")
	builder.WriteString("text=")
	builder.WriteString(rm.Text)
	builder.WriteByte(')')
	return builder.String()
}

// RoadMaps is a parsable slice of RoadMap.
type RoadMaps []*RoadMap