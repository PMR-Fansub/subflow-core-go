// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"subflow-core-go/pkg/ent/task"
	"subflow-core-go/pkg/ent/taskrecord"
	"subflow-core-go/pkg/ent/tasktag"
	"subflow-core-go/pkg/ent/team"
	"subflow-core-go/pkg/ent/workflow"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (tc *TaskCreate) SetName(s string) *TaskCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetType sets the "type" field.
func (tc *TaskCreate) SetType(i int) *TaskCreate {
	tc.mutation.SetType(i)
	return tc
}

// SetStatus sets the "status" field.
func (tc *TaskCreate) SetStatus(i int) *TaskCreate {
	tc.mutation.SetStatus(i)
	return tc
}

// SetDesc sets the "desc" field.
func (tc *TaskCreate) SetDesc(s string) *TaskCreate {
	tc.mutation.SetDesc(s)
	return tc
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (tc *TaskCreate) SetNillableDesc(s *string) *TaskCreate {
	if s != nil {
		tc.SetDesc(*s)
	}
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TaskCreate) SetCreatedAt(t time.Time) *TaskCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCreatedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetCompletedAt sets the "completed_at" field.
func (tc *TaskCreate) SetCompletedAt(t time.Time) *TaskCreate {
	tc.mutation.SetCompletedAt(t)
	return tc
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCompletedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetCompletedAt(*t)
	}
	return tc
}

// AddTaskRecordIDs adds the "task_records" edge to the TaskRecord entity by IDs.
func (tc *TaskCreate) AddTaskRecordIDs(ids ...int) *TaskCreate {
	tc.mutation.AddTaskRecordIDs(ids...)
	return tc
}

// AddTaskRecords adds the "task_records" edges to the TaskRecord entity.
func (tc *TaskCreate) AddTaskRecords(t ...*TaskRecord) *TaskCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTaskRecordIDs(ids...)
}

// AddTaskTagIDs adds the "task_tags" edge to the TaskTag entity by IDs.
func (tc *TaskCreate) AddTaskTagIDs(ids ...int) *TaskCreate {
	tc.mutation.AddTaskTagIDs(ids...)
	return tc
}

// AddTaskTags adds the "task_tags" edges to the TaskTag entity.
func (tc *TaskCreate) AddTaskTags(t ...*TaskTag) *TaskCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTaskTagIDs(ids...)
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (tc *TaskCreate) SetWorkflowID(id int) *TaskCreate {
	tc.mutation.SetWorkflowID(id)
	return tc
}

// SetNillableWorkflowID sets the "workflow" edge to the Workflow entity by ID if the given value is not nil.
func (tc *TaskCreate) SetNillableWorkflowID(id *int) *TaskCreate {
	if id != nil {
		tc = tc.SetWorkflowID(*id)
	}
	return tc
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (tc *TaskCreate) SetWorkflow(w *Workflow) *TaskCreate {
	return tc.SetWorkflowID(w.ID)
}

// SetTeamID sets the "team" edge to the Team entity by ID.
func (tc *TaskCreate) SetTeamID(id int) *TaskCreate {
	tc.mutation.SetTeamID(id)
	return tc
}

// SetNillableTeamID sets the "team" edge to the Team entity by ID if the given value is not nil.
func (tc *TaskCreate) SetNillableTeamID(id *int) *TaskCreate {
	if id != nil {
		tc = tc.SetTeamID(*id)
	}
	return tc
}

// SetTeam sets the "team" edge to the Team entity.
func (tc *TaskCreate) SetTeam(t *Team) *TaskCreate {
	return tc.SetTeamID(t.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tc *TaskCreate) Mutation() *TaskMutation {
	return tc.mutation
}

// Save creates the Task in the database.
func (tc *TaskCreate) Save(ctx context.Context) (*Task, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TaskCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TaskCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TaskCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := task.DefaultCreatedAt
		tc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaskCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Task.name"`)}
	}
	if _, ok := tc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Task.type"`)}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Task.status"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Task.created_at"`)}
	}
	return nil
}

func (tc *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TaskCreate) createSpec() (*Task, *sqlgraph.CreateSpec) {
	var (
		_node = &Task{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(task.Table, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	)
	_spec.OnConflict = tc.conflict
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(task.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.GetType(); ok {
		_spec.SetField(task.FieldType, field.TypeInt, value)
		_node.Type = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	if value, ok := tc.mutation.Desc(); ok {
		_spec.SetField(task.FieldDesc, field.TypeString, value)
		_node.Desc = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(task.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.CompletedAt(); ok {
		_spec.SetField(task.FieldCompletedAt, field.TypeTime, value)
		_node.CompletedAt = value
	}
	if nodes := tc.mutation.TaskRecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.TaskRecordsTable,
			Columns: []string{task.TaskRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TaskTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TaskTagsTable,
			Columns: task.TaskTagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tasktag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.WorkflowTable,
			Columns: []string{task.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.workflow_tasks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TeamTable,
			Columns: []string{task.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.team_tasks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Task.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (tc *TaskCreate) OnConflict(opts ...sql.ConflictOption) *TaskUpsertOne {
	tc.conflict = opts
	return &TaskUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TaskCreate) OnConflictColumns(columns ...string) *TaskUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TaskUpsertOne{
		create: tc,
	}
}

type (
	// TaskUpsertOne is the builder for "upsert"-ing
	//  one Task node.
	TaskUpsertOne struct {
		create *TaskCreate
	}

	// TaskUpsert is the "OnConflict" setter.
	TaskUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *TaskUpsert) SetName(v string) *TaskUpsert {
	u.Set(task.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TaskUpsert) UpdateName() *TaskUpsert {
	u.SetExcluded(task.FieldName)
	return u
}

// SetType sets the "type" field.
func (u *TaskUpsert) SetType(v int) *TaskUpsert {
	u.Set(task.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *TaskUpsert) UpdateType() *TaskUpsert {
	u.SetExcluded(task.FieldType)
	return u
}

// AddType adds v to the "type" field.
func (u *TaskUpsert) AddType(v int) *TaskUpsert {
	u.Add(task.FieldType, v)
	return u
}

// SetStatus sets the "status" field.
func (u *TaskUpsert) SetStatus(v int) *TaskUpsert {
	u.Set(task.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *TaskUpsert) UpdateStatus() *TaskUpsert {
	u.SetExcluded(task.FieldStatus)
	return u
}

// AddStatus adds v to the "status" field.
func (u *TaskUpsert) AddStatus(v int) *TaskUpsert {
	u.Add(task.FieldStatus, v)
	return u
}

// SetDesc sets the "desc" field.
func (u *TaskUpsert) SetDesc(v string) *TaskUpsert {
	u.Set(task.FieldDesc, v)
	return u
}

// UpdateDesc sets the "desc" field to the value that was provided on create.
func (u *TaskUpsert) UpdateDesc() *TaskUpsert {
	u.SetExcluded(task.FieldDesc)
	return u
}

// ClearDesc clears the value of the "desc" field.
func (u *TaskUpsert) ClearDesc() *TaskUpsert {
	u.SetNull(task.FieldDesc)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskUpsert) SetCreatedAt(v time.Time) *TaskUpsert {
	u.Set(task.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskUpsert) UpdateCreatedAt() *TaskUpsert {
	u.SetExcluded(task.FieldCreatedAt)
	return u
}

// SetCompletedAt sets the "completed_at" field.
func (u *TaskUpsert) SetCompletedAt(v time.Time) *TaskUpsert {
	u.Set(task.FieldCompletedAt, v)
	return u
}

// UpdateCompletedAt sets the "completed_at" field to the value that was provided on create.
func (u *TaskUpsert) UpdateCompletedAt() *TaskUpsert {
	u.SetExcluded(task.FieldCompletedAt)
	return u
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (u *TaskUpsert) ClearCompletedAt() *TaskUpsert {
	u.SetNull(task.FieldCompletedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TaskUpsertOne) UpdateNewValues() *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Task.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TaskUpsertOne) Ignore() *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskUpsertOne) DoNothing() *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskCreate.OnConflict
// documentation for more info.
func (u *TaskUpsertOne) Update(set func(*TaskUpsert)) *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *TaskUpsertOne) SetName(v string) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateName() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateName()
	})
}

// SetType sets the "type" field.
func (u *TaskUpsertOne) SetType(v int) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetType(v)
	})
}

// AddType adds v to the "type" field.
func (u *TaskUpsertOne) AddType(v int) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.AddType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateType() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateType()
	})
}

// SetStatus sets the "status" field.
func (u *TaskUpsertOne) SetStatus(v int) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *TaskUpsertOne) AddStatus(v int) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateStatus() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateStatus()
	})
}

// SetDesc sets the "desc" field.
func (u *TaskUpsertOne) SetDesc(v string) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetDesc(v)
	})
}

// UpdateDesc sets the "desc" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateDesc() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateDesc()
	})
}

// ClearDesc clears the value of the "desc" field.
func (u *TaskUpsertOne) ClearDesc() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.ClearDesc()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskUpsertOne) SetCreatedAt(v time.Time) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateCreatedAt() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetCompletedAt sets the "completed_at" field.
func (u *TaskUpsertOne) SetCompletedAt(v time.Time) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetCompletedAt(v)
	})
}

// UpdateCompletedAt sets the "completed_at" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateCompletedAt() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateCompletedAt()
	})
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (u *TaskUpsertOne) ClearCompletedAt() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.ClearCompletedAt()
	})
}

// Exec executes the query.
func (u *TaskUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TaskUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TaskUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	builders []*TaskCreate
	conflict []sql.ConflictOption
}

// Save creates the Task entities in the database.
func (tcb *TaskCreateBulk) Save(ctx context.Context) ([]*Task, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Task, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TaskCreateBulk) SaveX(ctx context.Context) []*Task {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TaskCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TaskCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Task.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (tcb *TaskCreateBulk) OnConflict(opts ...sql.ConflictOption) *TaskUpsertBulk {
	tcb.conflict = opts
	return &TaskUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TaskCreateBulk) OnConflictColumns(columns ...string) *TaskUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TaskUpsertBulk{
		create: tcb,
	}
}

// TaskUpsertBulk is the builder for "upsert"-ing
// a bulk of Task nodes.
type TaskUpsertBulk struct {
	create *TaskCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TaskUpsertBulk) UpdateNewValues() *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TaskUpsertBulk) Ignore() *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskUpsertBulk) DoNothing() *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskCreateBulk.OnConflict
// documentation for more info.
func (u *TaskUpsertBulk) Update(set func(*TaskUpsert)) *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *TaskUpsertBulk) SetName(v string) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateName() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateName()
	})
}

// SetType sets the "type" field.
func (u *TaskUpsertBulk) SetType(v int) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetType(v)
	})
}

// AddType adds v to the "type" field.
func (u *TaskUpsertBulk) AddType(v int) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.AddType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateType() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateType()
	})
}

// SetStatus sets the "status" field.
func (u *TaskUpsertBulk) SetStatus(v int) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *TaskUpsertBulk) AddStatus(v int) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateStatus() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateStatus()
	})
}

// SetDesc sets the "desc" field.
func (u *TaskUpsertBulk) SetDesc(v string) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetDesc(v)
	})
}

// UpdateDesc sets the "desc" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateDesc() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateDesc()
	})
}

// ClearDesc clears the value of the "desc" field.
func (u *TaskUpsertBulk) ClearDesc() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.ClearDesc()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskUpsertBulk) SetCreatedAt(v time.Time) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateCreatedAt() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetCompletedAt sets the "completed_at" field.
func (u *TaskUpsertBulk) SetCompletedAt(v time.Time) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetCompletedAt(v)
	})
}

// UpdateCompletedAt sets the "completed_at" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateCompletedAt() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateCompletedAt()
	})
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (u *TaskUpsertBulk) ClearCompletedAt() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.ClearCompletedAt()
	})
}

// Exec executes the query.
func (u *TaskUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TaskCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
