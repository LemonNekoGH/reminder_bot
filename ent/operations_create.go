// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/lemonnekogh/reminderbot/ent/operations"
)

// OperationsCreate is the builder for creating a Operations entity.
type OperationsCreate struct {
	config
	mutation *OperationsMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (oc *OperationsCreate) SetType(i int) *OperationsCreate {
	oc.mutation.SetType(i)
	return oc
}

// SetCompleted sets the "completed" field.
func (oc *OperationsCreate) SetCompleted(b bool) *OperationsCreate {
	oc.mutation.SetCompleted(b)
	return oc
}

// SetNillableCompleted sets the "completed" field if the given value is not nil.
func (oc *OperationsCreate) SetNillableCompleted(b *bool) *OperationsCreate {
	if b != nil {
		oc.SetCompleted(*b)
	}
	return oc
}

// SetSuccess sets the "success" field.
func (oc *OperationsCreate) SetSuccess(b bool) *OperationsCreate {
	oc.mutation.SetSuccess(b)
	return oc
}

// SetNillableSuccess sets the "success" field if the given value is not nil.
func (oc *OperationsCreate) SetNillableSuccess(b *bool) *OperationsCreate {
	if b != nil {
		oc.SetSuccess(*b)
	}
	return oc
}

// SetOperator sets the "operator" field.
func (oc *OperationsCreate) SetOperator(i int64) *OperationsCreate {
	oc.mutation.SetOperator(i)
	return oc
}

// SetMessageID sets the "message_id" field.
func (oc *OperationsCreate) SetMessageID(i int) *OperationsCreate {
	oc.mutation.SetMessageID(i)
	return oc
}

// SetRemindID sets the "remind_id" field.
func (oc *OperationsCreate) SetRemindID(u uuid.UUID) *OperationsCreate {
	oc.mutation.SetRemindID(u)
	return oc
}

// SetNillableRemindID sets the "remind_id" field if the given value is not nil.
func (oc *OperationsCreate) SetNillableRemindID(u *uuid.UUID) *OperationsCreate {
	if u != nil {
		oc.SetRemindID(*u)
	}
	return oc
}

// SetID sets the "id" field.
func (oc *OperationsCreate) SetID(u uuid.UUID) *OperationsCreate {
	oc.mutation.SetID(u)
	return oc
}

// Mutation returns the OperationsMutation object of the builder.
func (oc *OperationsCreate) Mutation() *OperationsMutation {
	return oc.mutation
}

// Save creates the Operations in the database.
func (oc *OperationsCreate) Save(ctx context.Context) (*Operations, error) {
	oc.defaults()
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OperationsCreate) SaveX(ctx context.Context) *Operations {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OperationsCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OperationsCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OperationsCreate) defaults() {
	if _, ok := oc.mutation.Completed(); !ok {
		v := operations.DefaultCompleted
		oc.mutation.SetCompleted(v)
	}
	if _, ok := oc.mutation.Success(); !ok {
		v := operations.DefaultSuccess
		oc.mutation.SetSuccess(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OperationsCreate) check() error {
	if _, ok := oc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Operations.type"`)}
	}
	if _, ok := oc.mutation.Completed(); !ok {
		return &ValidationError{Name: "completed", err: errors.New(`ent: missing required field "Operations.completed"`)}
	}
	if _, ok := oc.mutation.Success(); !ok {
		return &ValidationError{Name: "success", err: errors.New(`ent: missing required field "Operations.success"`)}
	}
	if _, ok := oc.mutation.Operator(); !ok {
		return &ValidationError{Name: "operator", err: errors.New(`ent: missing required field "Operations.operator"`)}
	}
	if _, ok := oc.mutation.MessageID(); !ok {
		return &ValidationError{Name: "message_id", err: errors.New(`ent: missing required field "Operations.message_id"`)}
	}
	return nil
}

func (oc *OperationsCreate) sqlSave(ctx context.Context) (*Operations, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OperationsCreate) createSpec() (*Operations, *sqlgraph.CreateSpec) {
	var (
		_node = &Operations{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(operations.Table, sqlgraph.NewFieldSpec(operations.FieldID, field.TypeUUID))
	)
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := oc.mutation.GetType(); ok {
		_spec.SetField(operations.FieldType, field.TypeInt, value)
		_node.Type = value
	}
	if value, ok := oc.mutation.Completed(); ok {
		_spec.SetField(operations.FieldCompleted, field.TypeBool, value)
		_node.Completed = value
	}
	if value, ok := oc.mutation.Success(); ok {
		_spec.SetField(operations.FieldSuccess, field.TypeBool, value)
		_node.Success = value
	}
	if value, ok := oc.mutation.Operator(); ok {
		_spec.SetField(operations.FieldOperator, field.TypeInt64, value)
		_node.Operator = value
	}
	if value, ok := oc.mutation.MessageID(); ok {
		_spec.SetField(operations.FieldMessageID, field.TypeInt, value)
		_node.MessageID = value
	}
	if value, ok := oc.mutation.RemindID(); ok {
		_spec.SetField(operations.FieldRemindID, field.TypeUUID, value)
		_node.RemindID = value
	}
	return _node, _spec
}

// OperationsCreateBulk is the builder for creating many Operations entities in bulk.
type OperationsCreateBulk struct {
	config
	builders []*OperationsCreate
}

// Save creates the Operations entities in the database.
func (ocb *OperationsCreateBulk) Save(ctx context.Context) ([]*Operations, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Operations, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OperationsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OperationsCreateBulk) SaveX(ctx context.Context) []*Operations {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OperationsCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OperationsCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}