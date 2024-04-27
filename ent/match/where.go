// Code generated by ent, DO NOT EDIT.

package match

import (
	"itamconnect/ent/internal"
	"itamconnect/ent/predicate"
	"itamconnect/internal/domain/valueobject"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Match {
	return predicate.Match(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Match {
	return predicate.Match(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Match {
	return predicate.Match(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Match {
	return predicate.Match(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Match {
	return predicate.Match(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Match {
	return predicate.Match(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Match {
	return predicate.Match(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldCreatedAt, v))
}

// MentorID applies equality check predicate on the "mentor_id" field. It's identical to MentorIDEQ.
func MentorID(v int) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldMentorID, v))
}

// MentiID applies equality check predicate on the "menti_id" field. It's identical to MentiIDEQ.
func MentiID(v int) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldMentiID, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v valueobject.MatchStatus) predicate.Match {
	vc := v
	return predicate.Match(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v valueobject.MatchStatus) predicate.Match {
	vc := v
	return predicate.Match(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...valueobject.MatchStatus) predicate.Match {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Match(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...valueobject.MatchStatus) predicate.Match {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Match(sql.FieldNotIn(FieldStatus, v...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Match {
	return predicate.Match(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Match {
	return predicate.Match(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Match {
	return predicate.Match(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Match {
	return predicate.Match(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Match {
	return predicate.Match(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Match {
	return predicate.Match(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Match {
	return predicate.Match(sql.FieldLTE(FieldCreatedAt, v))
}

// MentorIDEQ applies the EQ predicate on the "mentor_id" field.
func MentorIDEQ(v int) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldMentorID, v))
}

// MentorIDNEQ applies the NEQ predicate on the "mentor_id" field.
func MentorIDNEQ(v int) predicate.Match {
	return predicate.Match(sql.FieldNEQ(FieldMentorID, v))
}

// MentorIDIn applies the In predicate on the "mentor_id" field.
func MentorIDIn(vs ...int) predicate.Match {
	return predicate.Match(sql.FieldIn(FieldMentorID, vs...))
}

// MentorIDNotIn applies the NotIn predicate on the "mentor_id" field.
func MentorIDNotIn(vs ...int) predicate.Match {
	return predicate.Match(sql.FieldNotIn(FieldMentorID, vs...))
}

// MentiIDEQ applies the EQ predicate on the "menti_id" field.
func MentiIDEQ(v int) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldMentiID, v))
}

// MentiIDNEQ applies the NEQ predicate on the "menti_id" field.
func MentiIDNEQ(v int) predicate.Match {
	return predicate.Match(sql.FieldNEQ(FieldMentiID, v))
}

// MentiIDIn applies the In predicate on the "menti_id" field.
func MentiIDIn(vs ...int) predicate.Match {
	return predicate.Match(sql.FieldIn(FieldMentiID, vs...))
}

// MentiIDNotIn applies the NotIn predicate on the "menti_id" field.
func MentiIDNotIn(vs ...int) predicate.Match {
	return predicate.Match(sql.FieldNotIn(FieldMentiID, vs...))
}

// HasMentor applies the HasEdge predicate on the "mentor" edge.
func HasMentor() predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MentorTable, MentorColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Mentor
		step.Edge.Schema = schemaConfig.Match
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMentorWith applies the HasEdge predicate on the "mentor" edge with a given conditions (other predicates).
func HasMentorWith(preds ...predicate.Mentor) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := newMentorStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Mentor
		step.Edge.Schema = schemaConfig.Match
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMenti applies the HasEdge predicate on the "menti" edge.
func HasMenti() predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MentiTable, MentiColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Menti
		step.Edge.Schema = schemaConfig.Match
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMentiWith applies the HasEdge predicate on the "menti" edge with a given conditions (other predicates).
func HasMentiWith(preds ...predicate.Menti) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := newMentiStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Menti
		step.Edge.Schema = schemaConfig.Match
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMessages applies the HasEdge predicate on the "messages" edge.
func HasMessages() predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, MessagesTable, MessagesColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Message
		step.Edge.Schema = schemaConfig.Message
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMessagesWith applies the HasEdge predicate on the "messages" edge with a given conditions (other predicates).
func HasMessagesWith(preds ...predicate.Message) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := newMessagesStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Message
		step.Edge.Schema = schemaConfig.Message
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Match) predicate.Match {
	return predicate.Match(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Match) predicate.Match {
	return predicate.Match(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Match) predicate.Match {
	return predicate.Match(sql.NotPredicates(p))
}
