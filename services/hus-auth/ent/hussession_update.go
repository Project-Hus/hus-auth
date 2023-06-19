// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hus-auth/ent/connectedsession"
	"hus-auth/ent/hussession"
	"hus-auth/ent/predicate"
	"hus-auth/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// HusSessionUpdate is the builder for updating HusSession entities.
type HusSessionUpdate struct {
	config
	hooks    []Hook
	mutation *HusSessionMutation
}

// Where appends a list predicates to the HusSessionUpdate builder.
func (hsu *HusSessionUpdate) Where(ps ...predicate.HusSession) *HusSessionUpdate {
	hsu.mutation.Where(ps...)
	return hsu
}

// SetTid sets the "tid" field.
func (hsu *HusSessionUpdate) SetTid(u uuid.UUID) *HusSessionUpdate {
	hsu.mutation.SetTid(u)
	return hsu
}

// SetNillableTid sets the "tid" field if the given value is not nil.
func (hsu *HusSessionUpdate) SetNillableTid(u *uuid.UUID) *HusSessionUpdate {
	if u != nil {
		hsu.SetTid(*u)
	}
	return hsu
}

// SetIat sets the "iat" field.
func (hsu *HusSessionUpdate) SetIat(t time.Time) *HusSessionUpdate {
	hsu.mutation.SetIat(t)
	return hsu
}

// SetNillableIat sets the "iat" field if the given value is not nil.
func (hsu *HusSessionUpdate) SetNillableIat(t *time.Time) *HusSessionUpdate {
	if t != nil {
		hsu.SetIat(*t)
	}
	return hsu
}

// SetPreserved sets the "preserved" field.
func (hsu *HusSessionUpdate) SetPreserved(b bool) *HusSessionUpdate {
	hsu.mutation.SetPreserved(b)
	return hsu
}

// SetNillablePreserved sets the "preserved" field if the given value is not nil.
func (hsu *HusSessionUpdate) SetNillablePreserved(b *bool) *HusSessionUpdate {
	if b != nil {
		hsu.SetPreserved(*b)
	}
	return hsu
}

// SetUID sets the "uid" field.
func (hsu *HusSessionUpdate) SetUID(u uint64) *HusSessionUpdate {
	hsu.mutation.SetUID(u)
	return hsu
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (hsu *HusSessionUpdate) SetNillableUID(u *uint64) *HusSessionUpdate {
	if u != nil {
		hsu.SetUID(*u)
	}
	return hsu
}

// ClearUID clears the value of the "uid" field.
func (hsu *HusSessionUpdate) ClearUID() *HusSessionUpdate {
	hsu.mutation.ClearUID()
	return hsu
}

// SetSignedAt sets the "signed_at" field.
func (hsu *HusSessionUpdate) SetSignedAt(t time.Time) *HusSessionUpdate {
	hsu.mutation.SetSignedAt(t)
	return hsu
}

// SetNillableSignedAt sets the "signed_at" field if the given value is not nil.
func (hsu *HusSessionUpdate) SetNillableSignedAt(t *time.Time) *HusSessionUpdate {
	if t != nil {
		hsu.SetSignedAt(*t)
	}
	return hsu
}

// ClearSignedAt clears the value of the "signed_at" field.
func (hsu *HusSessionUpdate) ClearSignedAt() *HusSessionUpdate {
	hsu.mutation.ClearSignedAt()
	return hsu
}

// SetUpdatedAt sets the "updated_at" field.
func (hsu *HusSessionUpdate) SetUpdatedAt(t time.Time) *HusSessionUpdate {
	hsu.mutation.SetUpdatedAt(t)
	return hsu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (hsu *HusSessionUpdate) SetUserID(id uint64) *HusSessionUpdate {
	hsu.mutation.SetUserID(id)
	return hsu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (hsu *HusSessionUpdate) SetNillableUserID(id *uint64) *HusSessionUpdate {
	if id != nil {
		hsu = hsu.SetUserID(*id)
	}
	return hsu
}

// SetUser sets the "user" edge to the User entity.
func (hsu *HusSessionUpdate) SetUser(u *User) *HusSessionUpdate {
	return hsu.SetUserID(u.ID)
}

// AddConnectedSessionIDs adds the "connected_session" edge to the ConnectedSession entity by IDs.
func (hsu *HusSessionUpdate) AddConnectedSessionIDs(ids ...int) *HusSessionUpdate {
	hsu.mutation.AddConnectedSessionIDs(ids...)
	return hsu
}

// AddConnectedSession adds the "connected_session" edges to the ConnectedSession entity.
func (hsu *HusSessionUpdate) AddConnectedSession(c ...*ConnectedSession) *HusSessionUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return hsu.AddConnectedSessionIDs(ids...)
}

// Mutation returns the HusSessionMutation object of the builder.
func (hsu *HusSessionUpdate) Mutation() *HusSessionMutation {
	return hsu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (hsu *HusSessionUpdate) ClearUser() *HusSessionUpdate {
	hsu.mutation.ClearUser()
	return hsu
}

// ClearConnectedSession clears all "connected_session" edges to the ConnectedSession entity.
func (hsu *HusSessionUpdate) ClearConnectedSession() *HusSessionUpdate {
	hsu.mutation.ClearConnectedSession()
	return hsu
}

// RemoveConnectedSessionIDs removes the "connected_session" edge to ConnectedSession entities by IDs.
func (hsu *HusSessionUpdate) RemoveConnectedSessionIDs(ids ...int) *HusSessionUpdate {
	hsu.mutation.RemoveConnectedSessionIDs(ids...)
	return hsu
}

// RemoveConnectedSession removes "connected_session" edges to ConnectedSession entities.
func (hsu *HusSessionUpdate) RemoveConnectedSession(c ...*ConnectedSession) *HusSessionUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return hsu.RemoveConnectedSessionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hsu *HusSessionUpdate) Save(ctx context.Context) (int, error) {
	hsu.defaults()
	return withHooks[int, HusSessionMutation](ctx, hsu.sqlSave, hsu.mutation, hsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hsu *HusSessionUpdate) SaveX(ctx context.Context) int {
	affected, err := hsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hsu *HusSessionUpdate) Exec(ctx context.Context) error {
	_, err := hsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hsu *HusSessionUpdate) ExecX(ctx context.Context) {
	if err := hsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hsu *HusSessionUpdate) defaults() {
	if _, ok := hsu.mutation.UpdatedAt(); !ok {
		v := hussession.UpdateDefaultUpdatedAt()
		hsu.mutation.SetUpdatedAt(v)
	}
}

func (hsu *HusSessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(hussession.Table, hussession.Columns, sqlgraph.NewFieldSpec(hussession.FieldID, field.TypeUUID))
	if ps := hsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hsu.mutation.Tid(); ok {
		_spec.SetField(hussession.FieldTid, field.TypeUUID, value)
	}
	if value, ok := hsu.mutation.Iat(); ok {
		_spec.SetField(hussession.FieldIat, field.TypeTime, value)
	}
	if value, ok := hsu.mutation.Preserved(); ok {
		_spec.SetField(hussession.FieldPreserved, field.TypeBool, value)
	}
	if value, ok := hsu.mutation.SignedAt(); ok {
		_spec.SetField(hussession.FieldSignedAt, field.TypeTime, value)
	}
	if hsu.mutation.SignedAtCleared() {
		_spec.ClearField(hussession.FieldSignedAt, field.TypeTime)
	}
	if value, ok := hsu.mutation.UpdatedAt(); ok {
		_spec.SetField(hussession.FieldUpdatedAt, field.TypeTime, value)
	}
	if hsu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hussession.UserTable,
			Columns: []string{hussession.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hsu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hussession.UserTable,
			Columns: []string{hussession.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hsu.mutation.ConnectedSessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hussession.ConnectedSessionTable,
			Columns: []string{hussession.ConnectedSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: connectedsession.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hsu.mutation.RemovedConnectedSessionIDs(); len(nodes) > 0 && !hsu.mutation.ConnectedSessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hussession.ConnectedSessionTable,
			Columns: []string{hussession.ConnectedSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: connectedsession.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hsu.mutation.ConnectedSessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hussession.ConnectedSessionTable,
			Columns: []string{hussession.ConnectedSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: connectedsession.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hussession.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hsu.mutation.done = true
	return n, nil
}

// HusSessionUpdateOne is the builder for updating a single HusSession entity.
type HusSessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HusSessionMutation
}

// SetTid sets the "tid" field.
func (hsuo *HusSessionUpdateOne) SetTid(u uuid.UUID) *HusSessionUpdateOne {
	hsuo.mutation.SetTid(u)
	return hsuo
}

// SetNillableTid sets the "tid" field if the given value is not nil.
func (hsuo *HusSessionUpdateOne) SetNillableTid(u *uuid.UUID) *HusSessionUpdateOne {
	if u != nil {
		hsuo.SetTid(*u)
	}
	return hsuo
}

// SetIat sets the "iat" field.
func (hsuo *HusSessionUpdateOne) SetIat(t time.Time) *HusSessionUpdateOne {
	hsuo.mutation.SetIat(t)
	return hsuo
}

// SetNillableIat sets the "iat" field if the given value is not nil.
func (hsuo *HusSessionUpdateOne) SetNillableIat(t *time.Time) *HusSessionUpdateOne {
	if t != nil {
		hsuo.SetIat(*t)
	}
	return hsuo
}

// SetPreserved sets the "preserved" field.
func (hsuo *HusSessionUpdateOne) SetPreserved(b bool) *HusSessionUpdateOne {
	hsuo.mutation.SetPreserved(b)
	return hsuo
}

// SetNillablePreserved sets the "preserved" field if the given value is not nil.
func (hsuo *HusSessionUpdateOne) SetNillablePreserved(b *bool) *HusSessionUpdateOne {
	if b != nil {
		hsuo.SetPreserved(*b)
	}
	return hsuo
}

// SetUID sets the "uid" field.
func (hsuo *HusSessionUpdateOne) SetUID(u uint64) *HusSessionUpdateOne {
	hsuo.mutation.SetUID(u)
	return hsuo
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (hsuo *HusSessionUpdateOne) SetNillableUID(u *uint64) *HusSessionUpdateOne {
	if u != nil {
		hsuo.SetUID(*u)
	}
	return hsuo
}

// ClearUID clears the value of the "uid" field.
func (hsuo *HusSessionUpdateOne) ClearUID() *HusSessionUpdateOne {
	hsuo.mutation.ClearUID()
	return hsuo
}

// SetSignedAt sets the "signed_at" field.
func (hsuo *HusSessionUpdateOne) SetSignedAt(t time.Time) *HusSessionUpdateOne {
	hsuo.mutation.SetSignedAt(t)
	return hsuo
}

// SetNillableSignedAt sets the "signed_at" field if the given value is not nil.
func (hsuo *HusSessionUpdateOne) SetNillableSignedAt(t *time.Time) *HusSessionUpdateOne {
	if t != nil {
		hsuo.SetSignedAt(*t)
	}
	return hsuo
}

// ClearSignedAt clears the value of the "signed_at" field.
func (hsuo *HusSessionUpdateOne) ClearSignedAt() *HusSessionUpdateOne {
	hsuo.mutation.ClearSignedAt()
	return hsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (hsuo *HusSessionUpdateOne) SetUpdatedAt(t time.Time) *HusSessionUpdateOne {
	hsuo.mutation.SetUpdatedAt(t)
	return hsuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (hsuo *HusSessionUpdateOne) SetUserID(id uint64) *HusSessionUpdateOne {
	hsuo.mutation.SetUserID(id)
	return hsuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (hsuo *HusSessionUpdateOne) SetNillableUserID(id *uint64) *HusSessionUpdateOne {
	if id != nil {
		hsuo = hsuo.SetUserID(*id)
	}
	return hsuo
}

// SetUser sets the "user" edge to the User entity.
func (hsuo *HusSessionUpdateOne) SetUser(u *User) *HusSessionUpdateOne {
	return hsuo.SetUserID(u.ID)
}

// AddConnectedSessionIDs adds the "connected_session" edge to the ConnectedSession entity by IDs.
func (hsuo *HusSessionUpdateOne) AddConnectedSessionIDs(ids ...int) *HusSessionUpdateOne {
	hsuo.mutation.AddConnectedSessionIDs(ids...)
	return hsuo
}

// AddConnectedSession adds the "connected_session" edges to the ConnectedSession entity.
func (hsuo *HusSessionUpdateOne) AddConnectedSession(c ...*ConnectedSession) *HusSessionUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return hsuo.AddConnectedSessionIDs(ids...)
}

// Mutation returns the HusSessionMutation object of the builder.
func (hsuo *HusSessionUpdateOne) Mutation() *HusSessionMutation {
	return hsuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (hsuo *HusSessionUpdateOne) ClearUser() *HusSessionUpdateOne {
	hsuo.mutation.ClearUser()
	return hsuo
}

// ClearConnectedSession clears all "connected_session" edges to the ConnectedSession entity.
func (hsuo *HusSessionUpdateOne) ClearConnectedSession() *HusSessionUpdateOne {
	hsuo.mutation.ClearConnectedSession()
	return hsuo
}

// RemoveConnectedSessionIDs removes the "connected_session" edge to ConnectedSession entities by IDs.
func (hsuo *HusSessionUpdateOne) RemoveConnectedSessionIDs(ids ...int) *HusSessionUpdateOne {
	hsuo.mutation.RemoveConnectedSessionIDs(ids...)
	return hsuo
}

// RemoveConnectedSession removes "connected_session" edges to ConnectedSession entities.
func (hsuo *HusSessionUpdateOne) RemoveConnectedSession(c ...*ConnectedSession) *HusSessionUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return hsuo.RemoveConnectedSessionIDs(ids...)
}

// Where appends a list predicates to the HusSessionUpdate builder.
func (hsuo *HusSessionUpdateOne) Where(ps ...predicate.HusSession) *HusSessionUpdateOne {
	hsuo.mutation.Where(ps...)
	return hsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (hsuo *HusSessionUpdateOne) Select(field string, fields ...string) *HusSessionUpdateOne {
	hsuo.fields = append([]string{field}, fields...)
	return hsuo
}

// Save executes the query and returns the updated HusSession entity.
func (hsuo *HusSessionUpdateOne) Save(ctx context.Context) (*HusSession, error) {
	hsuo.defaults()
	return withHooks[*HusSession, HusSessionMutation](ctx, hsuo.sqlSave, hsuo.mutation, hsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hsuo *HusSessionUpdateOne) SaveX(ctx context.Context) *HusSession {
	node, err := hsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (hsuo *HusSessionUpdateOne) Exec(ctx context.Context) error {
	_, err := hsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hsuo *HusSessionUpdateOne) ExecX(ctx context.Context) {
	if err := hsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hsuo *HusSessionUpdateOne) defaults() {
	if _, ok := hsuo.mutation.UpdatedAt(); !ok {
		v := hussession.UpdateDefaultUpdatedAt()
		hsuo.mutation.SetUpdatedAt(v)
	}
}

func (hsuo *HusSessionUpdateOne) sqlSave(ctx context.Context) (_node *HusSession, err error) {
	_spec := sqlgraph.NewUpdateSpec(hussession.Table, hussession.Columns, sqlgraph.NewFieldSpec(hussession.FieldID, field.TypeUUID))
	id, ok := hsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HusSession.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := hsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hussession.FieldID)
		for _, f := range fields {
			if !hussession.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != hussession.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := hsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hsuo.mutation.Tid(); ok {
		_spec.SetField(hussession.FieldTid, field.TypeUUID, value)
	}
	if value, ok := hsuo.mutation.Iat(); ok {
		_spec.SetField(hussession.FieldIat, field.TypeTime, value)
	}
	if value, ok := hsuo.mutation.Preserved(); ok {
		_spec.SetField(hussession.FieldPreserved, field.TypeBool, value)
	}
	if value, ok := hsuo.mutation.SignedAt(); ok {
		_spec.SetField(hussession.FieldSignedAt, field.TypeTime, value)
	}
	if hsuo.mutation.SignedAtCleared() {
		_spec.ClearField(hussession.FieldSignedAt, field.TypeTime)
	}
	if value, ok := hsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(hussession.FieldUpdatedAt, field.TypeTime, value)
	}
	if hsuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hussession.UserTable,
			Columns: []string{hussession.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hsuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hussession.UserTable,
			Columns: []string{hussession.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hsuo.mutation.ConnectedSessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hussession.ConnectedSessionTable,
			Columns: []string{hussession.ConnectedSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: connectedsession.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hsuo.mutation.RemovedConnectedSessionIDs(); len(nodes) > 0 && !hsuo.mutation.ConnectedSessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hussession.ConnectedSessionTable,
			Columns: []string{hussession.ConnectedSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: connectedsession.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hsuo.mutation.ConnectedSessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hussession.ConnectedSessionTable,
			Columns: []string{hussession.ConnectedSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: connectedsession.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &HusSession{config: hsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, hsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hussession.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	hsuo.mutation.done = true
	return _node, nil
}
