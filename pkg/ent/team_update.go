// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"subflow-core-go/pkg/ent/predicate"
	"subflow-core-go/pkg/ent/task"
	"subflow-core-go/pkg/ent/team"
	"subflow-core-go/pkg/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeamUpdate is the builder for updating Team entities.
type TeamUpdate struct {
	config
	hooks    []Hook
	mutation *TeamMutation
}

// Where appends a list predicates to the TeamUpdate builder.
func (tu *TeamUpdate) Where(ps ...predicate.Team) *TeamUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TeamUpdate) SetName(s string) *TeamUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetStatus sets the "status" field.
func (tu *TeamUpdate) SetStatus(i int) *TeamUpdate {
	tu.mutation.ResetStatus()
	tu.mutation.SetStatus(i)
	return tu
}

// AddStatus adds i to the "status" field.
func (tu *TeamUpdate) AddStatus(i int) *TeamUpdate {
	tu.mutation.AddStatus(i)
	return tu
}

// SetQqGroup sets the "qq_group" field.
func (tu *TeamUpdate) SetQqGroup(s string) *TeamUpdate {
	tu.mutation.SetQqGroup(s)
	return tu
}

// SetNillableQqGroup sets the "qq_group" field if the given value is not nil.
func (tu *TeamUpdate) SetNillableQqGroup(s *string) *TeamUpdate {
	if s != nil {
		tu.SetQqGroup(*s)
	}
	return tu
}

// ClearQqGroup clears the value of the "qq_group" field.
func (tu *TeamUpdate) ClearQqGroup() *TeamUpdate {
	tu.mutation.ClearQqGroup()
	return tu
}

// SetLogo sets the "logo" field.
func (tu *TeamUpdate) SetLogo(s string) *TeamUpdate {
	tu.mutation.SetLogo(s)
	return tu
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (tu *TeamUpdate) SetNillableLogo(s *string) *TeamUpdate {
	if s != nil {
		tu.SetLogo(*s)
	}
	return tu
}

// ClearLogo clears the value of the "logo" field.
func (tu *TeamUpdate) ClearLogo() *TeamUpdate {
	tu.mutation.ClearLogo()
	return tu
}

// SetDesc sets the "desc" field.
func (tu *TeamUpdate) SetDesc(s string) *TeamUpdate {
	tu.mutation.SetDesc(s)
	return tu
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (tu *TeamUpdate) SetNillableDesc(s *string) *TeamUpdate {
	if s != nil {
		tu.SetDesc(*s)
	}
	return tu
}

// ClearDesc clears the value of the "desc" field.
func (tu *TeamUpdate) ClearDesc() *TeamUpdate {
	tu.mutation.ClearDesc()
	return tu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (tu *TeamUpdate) AddUserIDs(ids ...int) *TeamUpdate {
	tu.mutation.AddUserIDs(ids...)
	return tu
}

// AddUsers adds the "users" edges to the User entity.
func (tu *TeamUpdate) AddUsers(u ...*User) *TeamUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.AddUserIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (tu *TeamUpdate) AddTaskIDs(ids ...int) *TeamUpdate {
	tu.mutation.AddTaskIDs(ids...)
	return tu
}

// AddTasks adds the "tasks" edges to the Task entity.
func (tu *TeamUpdate) AddTasks(t ...*Task) *TeamUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTaskIDs(ids...)
}

// Mutation returns the TeamMutation object of the builder.
func (tu *TeamUpdate) Mutation() *TeamMutation {
	return tu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (tu *TeamUpdate) ClearUsers() *TeamUpdate {
	tu.mutation.ClearUsers()
	return tu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (tu *TeamUpdate) RemoveUserIDs(ids ...int) *TeamUpdate {
	tu.mutation.RemoveUserIDs(ids...)
	return tu
}

// RemoveUsers removes "users" edges to User entities.
func (tu *TeamUpdate) RemoveUsers(u ...*User) *TeamUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.RemoveUserIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (tu *TeamUpdate) ClearTasks() *TeamUpdate {
	tu.mutation.ClearTasks()
	return tu
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (tu *TeamUpdate) RemoveTaskIDs(ids ...int) *TeamUpdate {
	tu.mutation.RemoveTaskIDs(ids...)
	return tu
}

// RemoveTasks removes "tasks" edges to Task entities.
func (tu *TeamUpdate) RemoveTasks(t ...*Task) *TeamUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TeamUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TeamUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TeamUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TeamUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TeamUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := team.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Team.name": %w`, err)}
		}
	}
	return nil
}

func (tu *TeamUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(team.Table, team.Columns, sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(team.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.SetField(team.FieldStatus, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedStatus(); ok {
		_spec.AddField(team.FieldStatus, field.TypeInt, value)
	}
	if value, ok := tu.mutation.QqGroup(); ok {
		_spec.SetField(team.FieldQqGroup, field.TypeString, value)
	}
	if tu.mutation.QqGroupCleared() {
		_spec.ClearField(team.FieldQqGroup, field.TypeString)
	}
	if value, ok := tu.mutation.Logo(); ok {
		_spec.SetField(team.FieldLogo, field.TypeString, value)
	}
	if tu.mutation.LogoCleared() {
		_spec.ClearField(team.FieldLogo, field.TypeString)
	}
	if value, ok := tu.mutation.Desc(); ok {
		_spec.SetField(team.FieldDesc, field.TypeString, value)
	}
	if tu.mutation.DescCleared() {
		_spec.ClearField(team.FieldDesc, field.TypeString)
	}
	if tu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !tu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
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
	if tu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.TasksTable,
			Columns: []string{team.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTasksIDs(); len(nodes) > 0 && !tu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.TasksTable,
			Columns: []string{team.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.TasksTable,
			Columns: []string{team.TasksColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{team.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TeamUpdateOne is the builder for updating a single Team entity.
type TeamUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeamMutation
}

// SetName sets the "name" field.
func (tuo *TeamUpdateOne) SetName(s string) *TeamUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TeamUpdateOne) SetStatus(i int) *TeamUpdateOne {
	tuo.mutation.ResetStatus()
	tuo.mutation.SetStatus(i)
	return tuo
}

// AddStatus adds i to the "status" field.
func (tuo *TeamUpdateOne) AddStatus(i int) *TeamUpdateOne {
	tuo.mutation.AddStatus(i)
	return tuo
}

// SetQqGroup sets the "qq_group" field.
func (tuo *TeamUpdateOne) SetQqGroup(s string) *TeamUpdateOne {
	tuo.mutation.SetQqGroup(s)
	return tuo
}

// SetNillableQqGroup sets the "qq_group" field if the given value is not nil.
func (tuo *TeamUpdateOne) SetNillableQqGroup(s *string) *TeamUpdateOne {
	if s != nil {
		tuo.SetQqGroup(*s)
	}
	return tuo
}

// ClearQqGroup clears the value of the "qq_group" field.
func (tuo *TeamUpdateOne) ClearQqGroup() *TeamUpdateOne {
	tuo.mutation.ClearQqGroup()
	return tuo
}

// SetLogo sets the "logo" field.
func (tuo *TeamUpdateOne) SetLogo(s string) *TeamUpdateOne {
	tuo.mutation.SetLogo(s)
	return tuo
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (tuo *TeamUpdateOne) SetNillableLogo(s *string) *TeamUpdateOne {
	if s != nil {
		tuo.SetLogo(*s)
	}
	return tuo
}

// ClearLogo clears the value of the "logo" field.
func (tuo *TeamUpdateOne) ClearLogo() *TeamUpdateOne {
	tuo.mutation.ClearLogo()
	return tuo
}

// SetDesc sets the "desc" field.
func (tuo *TeamUpdateOne) SetDesc(s string) *TeamUpdateOne {
	tuo.mutation.SetDesc(s)
	return tuo
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (tuo *TeamUpdateOne) SetNillableDesc(s *string) *TeamUpdateOne {
	if s != nil {
		tuo.SetDesc(*s)
	}
	return tuo
}

// ClearDesc clears the value of the "desc" field.
func (tuo *TeamUpdateOne) ClearDesc() *TeamUpdateOne {
	tuo.mutation.ClearDesc()
	return tuo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (tuo *TeamUpdateOne) AddUserIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.AddUserIDs(ids...)
	return tuo
}

// AddUsers adds the "users" edges to the User entity.
func (tuo *TeamUpdateOne) AddUsers(u ...*User) *TeamUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.AddUserIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (tuo *TeamUpdateOne) AddTaskIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.AddTaskIDs(ids...)
	return tuo
}

// AddTasks adds the "tasks" edges to the Task entity.
func (tuo *TeamUpdateOne) AddTasks(t ...*Task) *TeamUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTaskIDs(ids...)
}

// Mutation returns the TeamMutation object of the builder.
func (tuo *TeamUpdateOne) Mutation() *TeamMutation {
	return tuo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (tuo *TeamUpdateOne) ClearUsers() *TeamUpdateOne {
	tuo.mutation.ClearUsers()
	return tuo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (tuo *TeamUpdateOne) RemoveUserIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.RemoveUserIDs(ids...)
	return tuo
}

// RemoveUsers removes "users" edges to User entities.
func (tuo *TeamUpdateOne) RemoveUsers(u ...*User) *TeamUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.RemoveUserIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (tuo *TeamUpdateOne) ClearTasks() *TeamUpdateOne {
	tuo.mutation.ClearTasks()
	return tuo
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (tuo *TeamUpdateOne) RemoveTaskIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.RemoveTaskIDs(ids...)
	return tuo
}

// RemoveTasks removes "tasks" edges to Task entities.
func (tuo *TeamUpdateOne) RemoveTasks(t ...*Task) *TeamUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTaskIDs(ids...)
}

// Where appends a list predicates to the TeamUpdate builder.
func (tuo *TeamUpdateOne) Where(ps ...predicate.Team) *TeamUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TeamUpdateOne) Select(field string, fields ...string) *TeamUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Team entity.
func (tuo *TeamUpdateOne) Save(ctx context.Context) (*Team, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TeamUpdateOne) SaveX(ctx context.Context) *Team {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TeamUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TeamUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TeamUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := team.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Team.name": %w`, err)}
		}
	}
	return nil
}

func (tuo *TeamUpdateOne) sqlSave(ctx context.Context) (_node *Team, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(team.Table, team.Columns, sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Team.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, team.FieldID)
		for _, f := range fields {
			if !team.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != team.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(team.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.SetField(team.FieldStatus, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedStatus(); ok {
		_spec.AddField(team.FieldStatus, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.QqGroup(); ok {
		_spec.SetField(team.FieldQqGroup, field.TypeString, value)
	}
	if tuo.mutation.QqGroupCleared() {
		_spec.ClearField(team.FieldQqGroup, field.TypeString)
	}
	if value, ok := tuo.mutation.Logo(); ok {
		_spec.SetField(team.FieldLogo, field.TypeString, value)
	}
	if tuo.mutation.LogoCleared() {
		_spec.ClearField(team.FieldLogo, field.TypeString)
	}
	if value, ok := tuo.mutation.Desc(); ok {
		_spec.SetField(team.FieldDesc, field.TypeString, value)
	}
	if tuo.mutation.DescCleared() {
		_spec.ClearField(team.FieldDesc, field.TypeString)
	}
	if tuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !tuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
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
	if tuo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.TasksTable,
			Columns: []string{team.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !tuo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.TasksTable,
			Columns: []string{team.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.TasksTable,
			Columns: []string{team.TasksColumn},
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
	_node = &Team{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{team.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
