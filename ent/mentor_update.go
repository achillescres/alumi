// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"itamconnect/ent/internal"
	"itamconnect/ent/match"
	"itamconnect/ent/menti"
	"itamconnect/ent/mentor"
	"itamconnect/ent/predicate"
	"itamconnect/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MentorUpdate is the builder for updating Mentor entities.
type MentorUpdate struct {
	config
	hooks    []Hook
	mutation *MentorMutation
}

// Where appends a list predicates to the MentorUpdate builder.
func (mu *MentorUpdate) Where(ps ...predicate.Mentor) *MentorUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetSearchFor sets the "search_for" field.
func (mu *MentorUpdate) SetSearchFor(s string) *MentorUpdate {
	mu.mutation.SetSearchFor(s)
	return mu
}

// SetNillableSearchFor sets the "search_for" field if the given value is not nil.
func (mu *MentorUpdate) SetNillableSearchFor(s *string) *MentorUpdate {
	if s != nil {
		mu.SetSearchFor(*s)
	}
	return mu
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (mu *MentorUpdate) AddUserIDs(ids ...int) *MentorUpdate {
	mu.mutation.AddUserIDs(ids...)
	return mu
}

// AddUser adds the "user" edges to the User entity.
func (mu *MentorUpdate) AddUser(u ...*User) *MentorUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mu.AddUserIDs(ids...)
}

// AddMentiIDs adds the "mentis" edge to the Menti entity by IDs.
func (mu *MentorUpdate) AddMentiIDs(ids ...int) *MentorUpdate {
	mu.mutation.AddMentiIDs(ids...)
	return mu
}

// AddMentis adds the "mentis" edges to the Menti entity.
func (mu *MentorUpdate) AddMentis(m ...*Menti) *MentorUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddMentiIDs(ids...)
}

// AddMatchIDs adds the "matches" edge to the Match entity by IDs.
func (mu *MentorUpdate) AddMatchIDs(ids ...int) *MentorUpdate {
	mu.mutation.AddMatchIDs(ids...)
	return mu
}

// AddMatches adds the "matches" edges to the Match entity.
func (mu *MentorUpdate) AddMatches(m ...*Match) *MentorUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddMatchIDs(ids...)
}

// Mutation returns the MentorMutation object of the builder.
func (mu *MentorUpdate) Mutation() *MentorMutation {
	return mu.mutation
}

// ClearUser clears all "user" edges to the User entity.
func (mu *MentorUpdate) ClearUser() *MentorUpdate {
	mu.mutation.ClearUser()
	return mu
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (mu *MentorUpdate) RemoveUserIDs(ids ...int) *MentorUpdate {
	mu.mutation.RemoveUserIDs(ids...)
	return mu
}

// RemoveUser removes "user" edges to User entities.
func (mu *MentorUpdate) RemoveUser(u ...*User) *MentorUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mu.RemoveUserIDs(ids...)
}

// ClearMentis clears all "mentis" edges to the Menti entity.
func (mu *MentorUpdate) ClearMentis() *MentorUpdate {
	mu.mutation.ClearMentis()
	return mu
}

// RemoveMentiIDs removes the "mentis" edge to Menti entities by IDs.
func (mu *MentorUpdate) RemoveMentiIDs(ids ...int) *MentorUpdate {
	mu.mutation.RemoveMentiIDs(ids...)
	return mu
}

// RemoveMentis removes "mentis" edges to Menti entities.
func (mu *MentorUpdate) RemoveMentis(m ...*Menti) *MentorUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveMentiIDs(ids...)
}

// ClearMatches clears all "matches" edges to the Match entity.
func (mu *MentorUpdate) ClearMatches() *MentorUpdate {
	mu.mutation.ClearMatches()
	return mu
}

// RemoveMatchIDs removes the "matches" edge to Match entities by IDs.
func (mu *MentorUpdate) RemoveMatchIDs(ids ...int) *MentorUpdate {
	mu.mutation.RemoveMatchIDs(ids...)
	return mu
}

// RemoveMatches removes "matches" edges to Match entities.
func (mu *MentorUpdate) RemoveMatches(m ...*Match) *MentorUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveMatchIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MentorUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MentorUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MentorUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MentorUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MentorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(mentor.Table, mentor.Columns, sqlgraph.NewFieldSpec(mentor.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.SearchFor(); ok {
		_spec.SetField(mentor.FieldSearchFor, field.TypeString, value)
	}
	if mu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   mentor.UserTable,
			Columns: []string{mentor.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = mu.schemaConfig.User
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedUserIDs(); len(nodes) > 0 && !mu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   mentor.UserTable,
			Columns: []string{mentor.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = mu.schemaConfig.User
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   mentor.UserTable,
			Columns: []string{mentor.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = mu.schemaConfig.User
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.MentisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   mentor.MentisTable,
			Columns: mentor.MentisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menti.FieldID, field.TypeInt),
			},
		}
		edge.Schema = mu.schemaConfig.Match
		createE := &MatchCreate{config: mu.config, mutation: newMatchMutation(mu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedMentisIDs(); len(nodes) > 0 && !mu.mutation.MentisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   mentor.MentisTable,
			Columns: mentor.MentisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menti.FieldID, field.TypeInt),
			},
		}
		edge.Schema = mu.schemaConfig.Match
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &MatchCreate{config: mu.config, mutation: newMatchMutation(mu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MentisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   mentor.MentisTable,
			Columns: mentor.MentisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menti.FieldID, field.TypeInt),
			},
		}
		edge.Schema = mu.schemaConfig.Match
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &MatchCreate{config: mu.config, mutation: newMatchMutation(mu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.MatchesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   mentor.MatchesTable,
			Columns: []string{mentor.MatchesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(match.FieldID, field.TypeInt),
			},
		}
		edge.Schema = mu.schemaConfig.Match
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedMatchesIDs(); len(nodes) > 0 && !mu.mutation.MatchesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   mentor.MatchesTable,
			Columns: []string{mentor.MatchesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(match.FieldID, field.TypeInt),
			},
		}
		edge.Schema = mu.schemaConfig.Match
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MatchesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   mentor.MatchesTable,
			Columns: []string{mentor.MatchesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(match.FieldID, field.TypeInt),
			},
		}
		edge.Schema = mu.schemaConfig.Match
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = mu.schemaConfig.Mentor
	ctx = internal.NewSchemaConfigContext(ctx, mu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mentor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MentorUpdateOne is the builder for updating a single Mentor entity.
type MentorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MentorMutation
}

// SetSearchFor sets the "search_for" field.
func (muo *MentorUpdateOne) SetSearchFor(s string) *MentorUpdateOne {
	muo.mutation.SetSearchFor(s)
	return muo
}

// SetNillableSearchFor sets the "search_for" field if the given value is not nil.
func (muo *MentorUpdateOne) SetNillableSearchFor(s *string) *MentorUpdateOne {
	if s != nil {
		muo.SetSearchFor(*s)
	}
	return muo
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (muo *MentorUpdateOne) AddUserIDs(ids ...int) *MentorUpdateOne {
	muo.mutation.AddUserIDs(ids...)
	return muo
}

// AddUser adds the "user" edges to the User entity.
func (muo *MentorUpdateOne) AddUser(u ...*User) *MentorUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return muo.AddUserIDs(ids...)
}

// AddMentiIDs adds the "mentis" edge to the Menti entity by IDs.
func (muo *MentorUpdateOne) AddMentiIDs(ids ...int) *MentorUpdateOne {
	muo.mutation.AddMentiIDs(ids...)
	return muo
}

// AddMentis adds the "mentis" edges to the Menti entity.
func (muo *MentorUpdateOne) AddMentis(m ...*Menti) *MentorUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddMentiIDs(ids...)
}

// AddMatchIDs adds the "matches" edge to the Match entity by IDs.
func (muo *MentorUpdateOne) AddMatchIDs(ids ...int) *MentorUpdateOne {
	muo.mutation.AddMatchIDs(ids...)
	return muo
}

// AddMatches adds the "matches" edges to the Match entity.
func (muo *MentorUpdateOne) AddMatches(m ...*Match) *MentorUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddMatchIDs(ids...)
}

// Mutation returns the MentorMutation object of the builder.
func (muo *MentorUpdateOne) Mutation() *MentorMutation {
	return muo.mutation
}

// ClearUser clears all "user" edges to the User entity.
func (muo *MentorUpdateOne) ClearUser() *MentorUpdateOne {
	muo.mutation.ClearUser()
	return muo
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (muo *MentorUpdateOne) RemoveUserIDs(ids ...int) *MentorUpdateOne {
	muo.mutation.RemoveUserIDs(ids...)
	return muo
}

// RemoveUser removes "user" edges to User entities.
func (muo *MentorUpdateOne) RemoveUser(u ...*User) *MentorUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return muo.RemoveUserIDs(ids...)
}

// ClearMentis clears all "mentis" edges to the Menti entity.
func (muo *MentorUpdateOne) ClearMentis() *MentorUpdateOne {
	muo.mutation.ClearMentis()
	return muo
}

// RemoveMentiIDs removes the "mentis" edge to Menti entities by IDs.
func (muo *MentorUpdateOne) RemoveMentiIDs(ids ...int) *MentorUpdateOne {
	muo.mutation.RemoveMentiIDs(ids...)
	return muo
}

// RemoveMentis removes "mentis" edges to Menti entities.
func (muo *MentorUpdateOne) RemoveMentis(m ...*Menti) *MentorUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveMentiIDs(ids...)
}

// ClearMatches clears all "matches" edges to the Match entity.
func (muo *MentorUpdateOne) ClearMatches() *MentorUpdateOne {
	muo.mutation.ClearMatches()
	return muo
}

// RemoveMatchIDs removes the "matches" edge to Match entities by IDs.
func (muo *MentorUpdateOne) RemoveMatchIDs(ids ...int) *MentorUpdateOne {
	muo.mutation.RemoveMatchIDs(ids...)
	return muo
}

// RemoveMatches removes "matches" edges to Match entities.
func (muo *MentorUpdateOne) RemoveMatches(m ...*Match) *MentorUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveMatchIDs(ids...)
}

// Where appends a list predicates to the MentorUpdate builder.
func (muo *MentorUpdateOne) Where(ps ...predicate.Mentor) *MentorUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MentorUpdateOne) Select(field string, fields ...string) *MentorUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Mentor entity.
func (muo *MentorUpdateOne) Save(ctx context.Context) (*Mentor, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MentorUpdateOne) SaveX(ctx context.Context) *Mentor {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MentorUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MentorUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MentorUpdateOne) sqlSave(ctx context.Context) (_node *Mentor, err error) {
	_spec := sqlgraph.NewUpdateSpec(mentor.Table, mentor.Columns, sqlgraph.NewFieldSpec(mentor.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Mentor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mentor.FieldID)
		for _, f := range fields {
			if !mentor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mentor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.SearchFor(); ok {
		_spec.SetField(mentor.FieldSearchFor, field.TypeString, value)
	}
	if muo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   mentor.UserTable,
			Columns: []string{mentor.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = muo.schemaConfig.User
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedUserIDs(); len(nodes) > 0 && !muo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   mentor.UserTable,
			Columns: []string{mentor.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = muo.schemaConfig.User
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   mentor.UserTable,
			Columns: []string{mentor.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = muo.schemaConfig.User
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.MentisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   mentor.MentisTable,
			Columns: mentor.MentisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menti.FieldID, field.TypeInt),
			},
		}
		edge.Schema = muo.schemaConfig.Match
		createE := &MatchCreate{config: muo.config, mutation: newMatchMutation(muo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedMentisIDs(); len(nodes) > 0 && !muo.mutation.MentisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   mentor.MentisTable,
			Columns: mentor.MentisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menti.FieldID, field.TypeInt),
			},
		}
		edge.Schema = muo.schemaConfig.Match
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &MatchCreate{config: muo.config, mutation: newMatchMutation(muo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MentisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   mentor.MentisTable,
			Columns: mentor.MentisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menti.FieldID, field.TypeInt),
			},
		}
		edge.Schema = muo.schemaConfig.Match
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &MatchCreate{config: muo.config, mutation: newMatchMutation(muo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.MatchesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   mentor.MatchesTable,
			Columns: []string{mentor.MatchesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(match.FieldID, field.TypeInt),
			},
		}
		edge.Schema = muo.schemaConfig.Match
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedMatchesIDs(); len(nodes) > 0 && !muo.mutation.MatchesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   mentor.MatchesTable,
			Columns: []string{mentor.MatchesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(match.FieldID, field.TypeInt),
			},
		}
		edge.Schema = muo.schemaConfig.Match
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MatchesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   mentor.MatchesTable,
			Columns: []string{mentor.MatchesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(match.FieldID, field.TypeInt),
			},
		}
		edge.Schema = muo.schemaConfig.Match
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = muo.schemaConfig.Mentor
	ctx = internal.NewSchemaConfigContext(ctx, muo.schemaConfig)
	_node = &Mentor{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mentor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}