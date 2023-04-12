// Code generated by ent, DO NOT EDIT.

package ent

import (
	"LinkTer-api/internel/ent/predicate"
	"LinkTer-api/internel/ent/session"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks    []Hook
	mutation *SessionMutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUserID sets the "user_id" field.
func (su *SessionUpdate) SetUserID(u uuid.UUID) *SessionUpdate {
	su.mutation.SetUserID(u)
	return su
}

// SetToken sets the "token" field.
func (su *SessionUpdate) SetToken(s string) *SessionUpdate {
	su.mutation.SetToken(s)
	return su
}

// SetIP sets the "ip" field.
func (su *SessionUpdate) SetIP(s string) *SessionUpdate {
	su.mutation.SetIP(s)
	return su
}

// SetStartedTime sets the "started_time" field.
func (su *SessionUpdate) SetStartedTime(t time.Time) *SessionUpdate {
	su.mutation.SetStartedTime(t)
	return su
}

// SetNillableStartedTime sets the "started_time" field if the given value is not nil.
func (su *SessionUpdate) SetNillableStartedTime(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetStartedTime(*t)
	}
	return su
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, SessionMutation](ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SessionUpdate) check() error {
	if v, ok := su.mutation.Token(); ok {
		if err := session.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`ent: validator failed for field "Session.token": %w`, err)}
		}
	}
	if v, ok := su.mutation.IP(); ok {
		if err := session.IPValidator(v); err != nil {
			return &ValidationError{Name: "ip", err: fmt.Errorf(`ent: validator failed for field "Session.ip": %w`, err)}
		}
	}
	return nil
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UserID(); ok {
		_spec.SetField(session.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := su.mutation.Token(); ok {
		_spec.SetField(session.FieldToken, field.TypeString, value)
	}
	if value, ok := su.mutation.IP(); ok {
		_spec.SetField(session.FieldIP, field.TypeString, value)
	}
	if value, ok := su.mutation.StartedTime(); ok {
		_spec.SetField(session.FieldStartedTime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SessionMutation
}

// SetUserID sets the "user_id" field.
func (suo *SessionUpdateOne) SetUserID(u uuid.UUID) *SessionUpdateOne {
	suo.mutation.SetUserID(u)
	return suo
}

// SetToken sets the "token" field.
func (suo *SessionUpdateOne) SetToken(s string) *SessionUpdateOne {
	suo.mutation.SetToken(s)
	return suo
}

// SetIP sets the "ip" field.
func (suo *SessionUpdateOne) SetIP(s string) *SessionUpdateOne {
	suo.mutation.SetIP(s)
	return suo
}

// SetStartedTime sets the "started_time" field.
func (suo *SessionUpdateOne) SetStartedTime(t time.Time) *SessionUpdateOne {
	suo.mutation.SetStartedTime(t)
	return suo
}

// SetNillableStartedTime sets the "started_time" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableStartedTime(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetStartedTime(*t)
	}
	return suo
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (suo *SessionUpdateOne) Where(ps ...predicate.Session) *SessionUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SessionUpdateOne) Select(field string, fields ...string) *SessionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Session entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {
	return withHooks[*Session, SessionMutation](ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SessionUpdateOne) check() error {
	if v, ok := suo.mutation.Token(); ok {
		if err := session.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`ent: validator failed for field "Session.token": %w`, err)}
		}
	}
	if v, ok := suo.mutation.IP(); ok {
		if err := session.IPValidator(v); err != nil {
			return &ValidationError{Name: "ip", err: fmt.Errorf(`ent: validator failed for field "Session.ip": %w`, err)}
		}
	}
	return nil
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (_node *Session, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Session.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for _, f := range fields {
			if !session.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != session.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UserID(); ok {
		_spec.SetField(session.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := suo.mutation.Token(); ok {
		_spec.SetField(session.FieldToken, field.TypeString, value)
	}
	if value, ok := suo.mutation.IP(); ok {
		_spec.SetField(session.FieldIP, field.TypeString, value)
	}
	if value, ok := suo.mutation.StartedTime(); ok {
		_spec.SetField(session.FieldStartedTime, field.TypeTime, value)
	}
	_node = &Session{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
