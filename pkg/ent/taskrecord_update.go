// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"subflow-core-go/pkg/ent/predicate"
	"subflow-core-go/pkg/ent/task"
	"subflow-core-go/pkg/ent/taskrecord"
	"subflow-core-go/pkg/ent/user"
	"subflow-core-go/pkg/ent/workflownode"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskRecordUpdate is the builder for updating TaskRecord entities.
type TaskRecordUpdate struct {
	config
	hooks    []Hook
	mutation *TaskRecordMutation
}

// Where appends a list predicates to the TaskRecordUpdate builder.
func (tru *TaskRecordUpdate) Where(ps ...predicate.TaskRecord) *TaskRecordUpdate {
	tru.mutation.Where(ps...)
	return tru
}

// SetStatus sets the "status" field.
func (tru *TaskRecordUpdate) SetStatus(i int) *TaskRecordUpdate {
	tru.mutation.ResetStatus()
	tru.mutation.SetStatus(i)
	return tru
}

// AddStatus adds i to the "status" field.
func (tru *TaskRecordUpdate) AddStatus(i int) *TaskRecordUpdate {
	tru.mutation.AddStatus(i)
	return tru
}

// SetCreatedAt sets the "created_at" field.
func (tru *TaskRecordUpdate) SetCreatedAt(t time.Time) *TaskRecordUpdate {
	tru.mutation.SetCreatedAt(t)
	return tru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tru *TaskRecordUpdate) SetNillableCreatedAt(t *time.Time) *TaskRecordUpdate {
	if t != nil {
		tru.SetCreatedAt(*t)
	}
	return tru
}

// SetAssignedAt sets the "assigned_at" field.
func (tru *TaskRecordUpdate) SetAssignedAt(t time.Time) *TaskRecordUpdate {
	tru.mutation.SetAssignedAt(t)
	return tru
}

// SetNillableAssignedAt sets the "assigned_at" field if the given value is not nil.
func (tru *TaskRecordUpdate) SetNillableAssignedAt(t *time.Time) *TaskRecordUpdate {
	if t != nil {
		tru.SetAssignedAt(*t)
	}
	return tru
}

// ClearAssignedAt clears the value of the "assigned_at" field.
func (tru *TaskRecordUpdate) ClearAssignedAt() *TaskRecordUpdate {
	tru.mutation.ClearAssignedAt()
	return tru
}

// SetCompletedAt sets the "completed_at" field.
func (tru *TaskRecordUpdate) SetCompletedAt(t time.Time) *TaskRecordUpdate {
	tru.mutation.SetCompletedAt(t)
	return tru
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (tru *TaskRecordUpdate) SetNillableCompletedAt(t *time.Time) *TaskRecordUpdate {
	if t != nil {
		tru.SetCompletedAt(*t)
	}
	return tru
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (tru *TaskRecordUpdate) ClearCompletedAt() *TaskRecordUpdate {
	tru.mutation.ClearCompletedAt()
	return tru
}

// SetRemark sets the "remark" field.
func (tru *TaskRecordUpdate) SetRemark(s string) *TaskRecordUpdate {
	tru.mutation.SetRemark(s)
	return tru
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (tru *TaskRecordUpdate) SetNillableRemark(s *string) *TaskRecordUpdate {
	if s != nil {
		tru.SetRemark(*s)
	}
	return tru
}

// ClearRemark clears the value of the "remark" field.
func (tru *TaskRecordUpdate) ClearRemark() *TaskRecordUpdate {
	tru.mutation.ClearRemark()
	return tru
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tru *TaskRecordUpdate) SetUserID(id int) *TaskRecordUpdate {
	tru.mutation.SetUserID(id)
	return tru
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (tru *TaskRecordUpdate) SetNillableUserID(id *int) *TaskRecordUpdate {
	if id != nil {
		tru = tru.SetUserID(*id)
	}
	return tru
}

// SetUser sets the "user" edge to the User entity.
func (tru *TaskRecordUpdate) SetUser(u *User) *TaskRecordUpdate {
	return tru.SetUserID(u.ID)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (tru *TaskRecordUpdate) SetTaskID(id int) *TaskRecordUpdate {
	tru.mutation.SetTaskID(id)
	return tru
}

// SetNillableTaskID sets the "task" edge to the Task entity by ID if the given value is not nil.
func (tru *TaskRecordUpdate) SetNillableTaskID(id *int) *TaskRecordUpdate {
	if id != nil {
		tru = tru.SetTaskID(*id)
	}
	return tru
}

// SetTask sets the "task" edge to the Task entity.
func (tru *TaskRecordUpdate) SetTask(t *Task) *TaskRecordUpdate {
	return tru.SetTaskID(t.ID)
}

// SetWorkflowNodeID sets the "workflow_node" edge to the WorkflowNode entity by ID.
func (tru *TaskRecordUpdate) SetWorkflowNodeID(id int) *TaskRecordUpdate {
	tru.mutation.SetWorkflowNodeID(id)
	return tru
}

// SetNillableWorkflowNodeID sets the "workflow_node" edge to the WorkflowNode entity by ID if the given value is not nil.
func (tru *TaskRecordUpdate) SetNillableWorkflowNodeID(id *int) *TaskRecordUpdate {
	if id != nil {
		tru = tru.SetWorkflowNodeID(*id)
	}
	return tru
}

// SetWorkflowNode sets the "workflow_node" edge to the WorkflowNode entity.
func (tru *TaskRecordUpdate) SetWorkflowNode(w *WorkflowNode) *TaskRecordUpdate {
	return tru.SetWorkflowNodeID(w.ID)
}

// Mutation returns the TaskRecordMutation object of the builder.
func (tru *TaskRecordUpdate) Mutation() *TaskRecordMutation {
	return tru.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (tru *TaskRecordUpdate) ClearUser() *TaskRecordUpdate {
	tru.mutation.ClearUser()
	return tru
}

// ClearTask clears the "task" edge to the Task entity.
func (tru *TaskRecordUpdate) ClearTask() *TaskRecordUpdate {
	tru.mutation.ClearTask()
	return tru
}

// ClearWorkflowNode clears the "workflow_node" edge to the WorkflowNode entity.
func (tru *TaskRecordUpdate) ClearWorkflowNode() *TaskRecordUpdate {
	tru.mutation.ClearWorkflowNode()
	return tru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tru *TaskRecordUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tru.sqlSave, tru.mutation, tru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tru *TaskRecordUpdate) SaveX(ctx context.Context) int {
	affected, err := tru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tru *TaskRecordUpdate) Exec(ctx context.Context) error {
	_, err := tru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tru *TaskRecordUpdate) ExecX(ctx context.Context) {
	if err := tru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tru *TaskRecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(taskrecord.Table, taskrecord.Columns, sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt))
	if ps := tru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tru.mutation.Status(); ok {
		_spec.SetField(taskrecord.FieldStatus, field.TypeInt, value)
	}
	if value, ok := tru.mutation.AddedStatus(); ok {
		_spec.AddField(taskrecord.FieldStatus, field.TypeInt, value)
	}
	if value, ok := tru.mutation.CreatedAt(); ok {
		_spec.SetField(taskrecord.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tru.mutation.AssignedAt(); ok {
		_spec.SetField(taskrecord.FieldAssignedAt, field.TypeTime, value)
	}
	if tru.mutation.AssignedAtCleared() {
		_spec.ClearField(taskrecord.FieldAssignedAt, field.TypeTime)
	}
	if value, ok := tru.mutation.CompletedAt(); ok {
		_spec.SetField(taskrecord.FieldCompletedAt, field.TypeTime, value)
	}
	if tru.mutation.CompletedAtCleared() {
		_spec.ClearField(taskrecord.FieldCompletedAt, field.TypeTime)
	}
	if value, ok := tru.mutation.Remark(); ok {
		_spec.SetField(taskrecord.FieldRemark, field.TypeString, value)
	}
	if tru.mutation.RemarkCleared() {
		_spec.ClearField(taskrecord.FieldRemark, field.TypeString)
	}
	if tru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskrecord.UserTable,
			Columns: []string{taskrecord.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskrecord.UserTable,
			Columns: []string{taskrecord.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tru.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskrecord.TaskTable,
			Columns: []string{taskrecord.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskrecord.TaskTable,
			Columns: []string{taskrecord.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tru.mutation.WorkflowNodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskrecord.WorkflowNodeTable,
			Columns: []string{taskrecord.WorkflowNodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflownode.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.WorkflowNodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskrecord.WorkflowNodeTable,
			Columns: []string{taskrecord.WorkflowNodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflownode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taskrecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tru.mutation.done = true
	return n, nil
}

// TaskRecordUpdateOne is the builder for updating a single TaskRecord entity.
type TaskRecordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskRecordMutation
}

// SetStatus sets the "status" field.
func (truo *TaskRecordUpdateOne) SetStatus(i int) *TaskRecordUpdateOne {
	truo.mutation.ResetStatus()
	truo.mutation.SetStatus(i)
	return truo
}

// AddStatus adds i to the "status" field.
func (truo *TaskRecordUpdateOne) AddStatus(i int) *TaskRecordUpdateOne {
	truo.mutation.AddStatus(i)
	return truo
}

// SetCreatedAt sets the "created_at" field.
func (truo *TaskRecordUpdateOne) SetCreatedAt(t time.Time) *TaskRecordUpdateOne {
	truo.mutation.SetCreatedAt(t)
	return truo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (truo *TaskRecordUpdateOne) SetNillableCreatedAt(t *time.Time) *TaskRecordUpdateOne {
	if t != nil {
		truo.SetCreatedAt(*t)
	}
	return truo
}

// SetAssignedAt sets the "assigned_at" field.
func (truo *TaskRecordUpdateOne) SetAssignedAt(t time.Time) *TaskRecordUpdateOne {
	truo.mutation.SetAssignedAt(t)
	return truo
}

// SetNillableAssignedAt sets the "assigned_at" field if the given value is not nil.
func (truo *TaskRecordUpdateOne) SetNillableAssignedAt(t *time.Time) *TaskRecordUpdateOne {
	if t != nil {
		truo.SetAssignedAt(*t)
	}
	return truo
}

// ClearAssignedAt clears the value of the "assigned_at" field.
func (truo *TaskRecordUpdateOne) ClearAssignedAt() *TaskRecordUpdateOne {
	truo.mutation.ClearAssignedAt()
	return truo
}

// SetCompletedAt sets the "completed_at" field.
func (truo *TaskRecordUpdateOne) SetCompletedAt(t time.Time) *TaskRecordUpdateOne {
	truo.mutation.SetCompletedAt(t)
	return truo
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (truo *TaskRecordUpdateOne) SetNillableCompletedAt(t *time.Time) *TaskRecordUpdateOne {
	if t != nil {
		truo.SetCompletedAt(*t)
	}
	return truo
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (truo *TaskRecordUpdateOne) ClearCompletedAt() *TaskRecordUpdateOne {
	truo.mutation.ClearCompletedAt()
	return truo
}

// SetRemark sets the "remark" field.
func (truo *TaskRecordUpdateOne) SetRemark(s string) *TaskRecordUpdateOne {
	truo.mutation.SetRemark(s)
	return truo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (truo *TaskRecordUpdateOne) SetNillableRemark(s *string) *TaskRecordUpdateOne {
	if s != nil {
		truo.SetRemark(*s)
	}
	return truo
}

// ClearRemark clears the value of the "remark" field.
func (truo *TaskRecordUpdateOne) ClearRemark() *TaskRecordUpdateOne {
	truo.mutation.ClearRemark()
	return truo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (truo *TaskRecordUpdateOne) SetUserID(id int) *TaskRecordUpdateOne {
	truo.mutation.SetUserID(id)
	return truo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (truo *TaskRecordUpdateOne) SetNillableUserID(id *int) *TaskRecordUpdateOne {
	if id != nil {
		truo = truo.SetUserID(*id)
	}
	return truo
}

// SetUser sets the "user" edge to the User entity.
func (truo *TaskRecordUpdateOne) SetUser(u *User) *TaskRecordUpdateOne {
	return truo.SetUserID(u.ID)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (truo *TaskRecordUpdateOne) SetTaskID(id int) *TaskRecordUpdateOne {
	truo.mutation.SetTaskID(id)
	return truo
}

// SetNillableTaskID sets the "task" edge to the Task entity by ID if the given value is not nil.
func (truo *TaskRecordUpdateOne) SetNillableTaskID(id *int) *TaskRecordUpdateOne {
	if id != nil {
		truo = truo.SetTaskID(*id)
	}
	return truo
}

// SetTask sets the "task" edge to the Task entity.
func (truo *TaskRecordUpdateOne) SetTask(t *Task) *TaskRecordUpdateOne {
	return truo.SetTaskID(t.ID)
}

// SetWorkflowNodeID sets the "workflow_node" edge to the WorkflowNode entity by ID.
func (truo *TaskRecordUpdateOne) SetWorkflowNodeID(id int) *TaskRecordUpdateOne {
	truo.mutation.SetWorkflowNodeID(id)
	return truo
}

// SetNillableWorkflowNodeID sets the "workflow_node" edge to the WorkflowNode entity by ID if the given value is not nil.
func (truo *TaskRecordUpdateOne) SetNillableWorkflowNodeID(id *int) *TaskRecordUpdateOne {
	if id != nil {
		truo = truo.SetWorkflowNodeID(*id)
	}
	return truo
}

// SetWorkflowNode sets the "workflow_node" edge to the WorkflowNode entity.
func (truo *TaskRecordUpdateOne) SetWorkflowNode(w *WorkflowNode) *TaskRecordUpdateOne {
	return truo.SetWorkflowNodeID(w.ID)
}

// Mutation returns the TaskRecordMutation object of the builder.
func (truo *TaskRecordUpdateOne) Mutation() *TaskRecordMutation {
	return truo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (truo *TaskRecordUpdateOne) ClearUser() *TaskRecordUpdateOne {
	truo.mutation.ClearUser()
	return truo
}

// ClearTask clears the "task" edge to the Task entity.
func (truo *TaskRecordUpdateOne) ClearTask() *TaskRecordUpdateOne {
	truo.mutation.ClearTask()
	return truo
}

// ClearWorkflowNode clears the "workflow_node" edge to the WorkflowNode entity.
func (truo *TaskRecordUpdateOne) ClearWorkflowNode() *TaskRecordUpdateOne {
	truo.mutation.ClearWorkflowNode()
	return truo
}

// Where appends a list predicates to the TaskRecordUpdate builder.
func (truo *TaskRecordUpdateOne) Where(ps ...predicate.TaskRecord) *TaskRecordUpdateOne {
	truo.mutation.Where(ps...)
	return truo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (truo *TaskRecordUpdateOne) Select(field string, fields ...string) *TaskRecordUpdateOne {
	truo.fields = append([]string{field}, fields...)
	return truo
}

// Save executes the query and returns the updated TaskRecord entity.
func (truo *TaskRecordUpdateOne) Save(ctx context.Context) (*TaskRecord, error) {
	return withHooks(ctx, truo.sqlSave, truo.mutation, truo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (truo *TaskRecordUpdateOne) SaveX(ctx context.Context) *TaskRecord {
	node, err := truo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (truo *TaskRecordUpdateOne) Exec(ctx context.Context) error {
	_, err := truo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (truo *TaskRecordUpdateOne) ExecX(ctx context.Context) {
	if err := truo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (truo *TaskRecordUpdateOne) sqlSave(ctx context.Context) (_node *TaskRecord, err error) {
	_spec := sqlgraph.NewUpdateSpec(taskrecord.Table, taskrecord.Columns, sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt))
	id, ok := truo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TaskRecord.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := truo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, taskrecord.FieldID)
		for _, f := range fields {
			if !taskrecord.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != taskrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := truo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := truo.mutation.Status(); ok {
		_spec.SetField(taskrecord.FieldStatus, field.TypeInt, value)
	}
	if value, ok := truo.mutation.AddedStatus(); ok {
		_spec.AddField(taskrecord.FieldStatus, field.TypeInt, value)
	}
	if value, ok := truo.mutation.CreatedAt(); ok {
		_spec.SetField(taskrecord.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := truo.mutation.AssignedAt(); ok {
		_spec.SetField(taskrecord.FieldAssignedAt, field.TypeTime, value)
	}
	if truo.mutation.AssignedAtCleared() {
		_spec.ClearField(taskrecord.FieldAssignedAt, field.TypeTime)
	}
	if value, ok := truo.mutation.CompletedAt(); ok {
		_spec.SetField(taskrecord.FieldCompletedAt, field.TypeTime, value)
	}
	if truo.mutation.CompletedAtCleared() {
		_spec.ClearField(taskrecord.FieldCompletedAt, field.TypeTime)
	}
	if value, ok := truo.mutation.Remark(); ok {
		_spec.SetField(taskrecord.FieldRemark, field.TypeString, value)
	}
	if truo.mutation.RemarkCleared() {
		_spec.ClearField(taskrecord.FieldRemark, field.TypeString)
	}
	if truo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskrecord.UserTable,
			Columns: []string{taskrecord.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskrecord.UserTable,
			Columns: []string{taskrecord.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if truo.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskrecord.TaskTable,
			Columns: []string{taskrecord.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskrecord.TaskTable,
			Columns: []string{taskrecord.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if truo.mutation.WorkflowNodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskrecord.WorkflowNodeTable,
			Columns: []string{taskrecord.WorkflowNodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflownode.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.WorkflowNodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskrecord.WorkflowNodeTable,
			Columns: []string{taskrecord.WorkflowNodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflownode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TaskRecord{config: truo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, truo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taskrecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	truo.mutation.done = true
	return _node, nil
}
