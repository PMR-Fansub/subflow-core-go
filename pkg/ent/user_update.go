// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"subflow-core-go/pkg/ent/predicate"
	"subflow-core-go/pkg/ent/role"
	"subflow-core-go/pkg/ent/taskrecord"
	"subflow-core-go/pkg/ent/team"
	"subflow-core-go/pkg/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNickname sets the "nickname" field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.mutation.SetNickname(s)
	return uu
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uu *UserUpdate) SetNillableNickname(s *string) *UserUpdate {
	if s != nil {
		uu.SetNickname(*s)
	}
	return uu
}

// ClearNickname clears the value of the "nickname" field.
func (uu *UserUpdate) ClearNickname() *UserUpdate {
	uu.mutation.ClearNickname()
	return uu
}

// SetStatus sets the "status" field.
func (uu *UserUpdate) SetStatus(i int) *UserUpdate {
	uu.mutation.ResetStatus()
	uu.mutation.SetStatus(i)
	return uu
}

// AddStatus adds i to the "status" field.
func (uu *UserUpdate) AddStatus(i int) *UserUpdate {
	uu.mutation.AddStatus(i)
	return uu
}

// SetRegisteredAt sets the "registered_at" field.
func (uu *UserUpdate) SetRegisteredAt(t time.Time) *UserUpdate {
	uu.mutation.SetRegisteredAt(t)
	return uu
}

// SetNillableRegisteredAt sets the "registered_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRegisteredAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetRegisteredAt(*t)
	}
	return uu
}

// SetRegisterIP sets the "register_ip" field.
func (uu *UserUpdate) SetRegisterIP(s string) *UserUpdate {
	uu.mutation.SetRegisterIP(s)
	return uu
}

// SetNillableRegisterIP sets the "register_ip" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRegisterIP(s *string) *UserUpdate {
	if s != nil {
		uu.SetRegisterIP(*s)
	}
	return uu
}

// ClearRegisterIP clears the value of the "register_ip" field.
func (uu *UserUpdate) ClearRegisterIP() *UserUpdate {
	uu.mutation.ClearRegisterIP()
	return uu
}

// SetLastLoggedAt sets the "last_logged_at" field.
func (uu *UserUpdate) SetLastLoggedAt(t time.Time) *UserUpdate {
	uu.mutation.SetLastLoggedAt(t)
	return uu
}

// SetNillableLastLoggedAt sets the "last_logged_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastLoggedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetLastLoggedAt(*t)
	}
	return uu
}

// SetLoginIP sets the "login_ip" field.
func (uu *UserUpdate) SetLoginIP(s string) *UserUpdate {
	uu.mutation.SetLoginIP(s)
	return uu
}

// SetNillableLoginIP sets the "login_ip" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLoginIP(s *string) *UserUpdate {
	if s != nil {
		uu.SetLoginIP(*s)
	}
	return uu
}

// ClearLoginIP clears the value of the "login_ip" field.
func (uu *UserUpdate) ClearLoginIP() *UserUpdate {
	uu.mutation.ClearLoginIP()
	return uu
}

// SetAvatar sets the "avatar" field.
func (uu *UserUpdate) SetAvatar(s string) *UserUpdate {
	uu.mutation.SetAvatar(s)
	return uu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatar(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatar(*s)
	}
	return uu
}

// ClearAvatar clears the value of the "avatar" field.
func (uu *UserUpdate) ClearAvatar() *UserUpdate {
	uu.mutation.ClearAvatar()
	return uu
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (uu *UserUpdate) AddRoleIDs(ids ...int) *UserUpdate {
	uu.mutation.AddRoleIDs(ids...)
	return uu
}

// AddRoles adds the "roles" edges to the Role entity.
func (uu *UserUpdate) AddRoles(r ...*Role) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.AddRoleIDs(ids...)
}

// AddTaskRecordIDs adds the "task_records" edge to the TaskRecord entity by IDs.
func (uu *UserUpdate) AddTaskRecordIDs(ids ...int) *UserUpdate {
	uu.mutation.AddTaskRecordIDs(ids...)
	return uu
}

// AddTaskRecords adds the "task_records" edges to the TaskRecord entity.
func (uu *UserUpdate) AddTaskRecords(t ...*TaskRecord) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddTaskRecordIDs(ids...)
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (uu *UserUpdate) AddTeamIDs(ids ...int) *UserUpdate {
	uu.mutation.AddTeamIDs(ids...)
	return uu
}

// AddTeams adds the "teams" edges to the Team entity.
func (uu *UserUpdate) AddTeams(t ...*Team) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddTeamIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (uu *UserUpdate) ClearRoles() *UserUpdate {
	uu.mutation.ClearRoles()
	return uu
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (uu *UserUpdate) RemoveRoleIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveRoleIDs(ids...)
	return uu
}

// RemoveRoles removes "roles" edges to Role entities.
func (uu *UserUpdate) RemoveRoles(r ...*Role) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.RemoveRoleIDs(ids...)
}

// ClearTaskRecords clears all "task_records" edges to the TaskRecord entity.
func (uu *UserUpdate) ClearTaskRecords() *UserUpdate {
	uu.mutation.ClearTaskRecords()
	return uu
}

// RemoveTaskRecordIDs removes the "task_records" edge to TaskRecord entities by IDs.
func (uu *UserUpdate) RemoveTaskRecordIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveTaskRecordIDs(ids...)
	return uu
}

// RemoveTaskRecords removes "task_records" edges to TaskRecord entities.
func (uu *UserUpdate) RemoveTaskRecords(t ...*TaskRecord) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveTaskRecordIDs(ids...)
}

// ClearTeams clears all "teams" edges to the Team entity.
func (uu *UserUpdate) ClearTeams() *UserUpdate {
	uu.mutation.ClearTeams()
	return uu
}

// RemoveTeamIDs removes the "teams" edge to Team entities by IDs.
func (uu *UserUpdate) RemoveTeamIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveTeamIDs(ids...)
	return uu
}

// RemoveTeams removes "teams" edges to Team entities.
func (uu *UserUpdate) RemoveTeams(t ...*Team) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveTeamIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if uu.mutation.NicknameCleared() {
		_spec.ClearField(user.FieldNickname, field.TypeString)
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedStatus(); ok {
		_spec.AddField(user.FieldStatus, field.TypeInt, value)
	}
	if value, ok := uu.mutation.RegisteredAt(); ok {
		_spec.SetField(user.FieldRegisteredAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.RegisterIP(); ok {
		_spec.SetField(user.FieldRegisterIP, field.TypeString, value)
	}
	if uu.mutation.RegisterIPCleared() {
		_spec.ClearField(user.FieldRegisterIP, field.TypeString)
	}
	if value, ok := uu.mutation.LastLoggedAt(); ok {
		_spec.SetField(user.FieldLastLoggedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.LoginIP(); ok {
		_spec.SetField(user.FieldLoginIP, field.TypeString, value)
	}
	if uu.mutation.LoginIPCleared() {
		_spec.ClearField(user.FieldLoginIP, field.TypeString)
	}
	if value, ok := uu.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uu.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if uu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedRolesIDs(); len(nodes) > 0 && !uu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.TaskRecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TaskRecordsTable,
			Columns: []string{user.TaskRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedTaskRecordsIDs(); len(nodes) > 0 && !uu.mutation.TaskRecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TaskRecordsTable,
			Columns: []string{user.TaskRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.TaskRecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TaskRecordsTable,
			Columns: []string{user.TaskRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedTeamsIDs(); len(nodes) > 0 && !uu.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNickname sets the "nickname" field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.mutation.SetNickname(s)
	return uuo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNickname(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNickname(*s)
	}
	return uuo
}

// ClearNickname clears the value of the "nickname" field.
func (uuo *UserUpdateOne) ClearNickname() *UserUpdateOne {
	uuo.mutation.ClearNickname()
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UserUpdateOne) SetStatus(i int) *UserUpdateOne {
	uuo.mutation.ResetStatus()
	uuo.mutation.SetStatus(i)
	return uuo
}

// AddStatus adds i to the "status" field.
func (uuo *UserUpdateOne) AddStatus(i int) *UserUpdateOne {
	uuo.mutation.AddStatus(i)
	return uuo
}

// SetRegisteredAt sets the "registered_at" field.
func (uuo *UserUpdateOne) SetRegisteredAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetRegisteredAt(t)
	return uuo
}

// SetNillableRegisteredAt sets the "registered_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRegisteredAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetRegisteredAt(*t)
	}
	return uuo
}

// SetRegisterIP sets the "register_ip" field.
func (uuo *UserUpdateOne) SetRegisterIP(s string) *UserUpdateOne {
	uuo.mutation.SetRegisterIP(s)
	return uuo
}

// SetNillableRegisterIP sets the "register_ip" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRegisterIP(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetRegisterIP(*s)
	}
	return uuo
}

// ClearRegisterIP clears the value of the "register_ip" field.
func (uuo *UserUpdateOne) ClearRegisterIP() *UserUpdateOne {
	uuo.mutation.ClearRegisterIP()
	return uuo
}

// SetLastLoggedAt sets the "last_logged_at" field.
func (uuo *UserUpdateOne) SetLastLoggedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetLastLoggedAt(t)
	return uuo
}

// SetNillableLastLoggedAt sets the "last_logged_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastLoggedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetLastLoggedAt(*t)
	}
	return uuo
}

// SetLoginIP sets the "login_ip" field.
func (uuo *UserUpdateOne) SetLoginIP(s string) *UserUpdateOne {
	uuo.mutation.SetLoginIP(s)
	return uuo
}

// SetNillableLoginIP sets the "login_ip" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLoginIP(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLoginIP(*s)
	}
	return uuo
}

// ClearLoginIP clears the value of the "login_ip" field.
func (uuo *UserUpdateOne) ClearLoginIP() *UserUpdateOne {
	uuo.mutation.ClearLoginIP()
	return uuo
}

// SetAvatar sets the "avatar" field.
func (uuo *UserUpdateOne) SetAvatar(s string) *UserUpdateOne {
	uuo.mutation.SetAvatar(s)
	return uuo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatar(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatar(*s)
	}
	return uuo
}

// ClearAvatar clears the value of the "avatar" field.
func (uuo *UserUpdateOne) ClearAvatar() *UserUpdateOne {
	uuo.mutation.ClearAvatar()
	return uuo
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (uuo *UserUpdateOne) AddRoleIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddRoleIDs(ids...)
	return uuo
}

// AddRoles adds the "roles" edges to the Role entity.
func (uuo *UserUpdateOne) AddRoles(r ...*Role) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.AddRoleIDs(ids...)
}

// AddTaskRecordIDs adds the "task_records" edge to the TaskRecord entity by IDs.
func (uuo *UserUpdateOne) AddTaskRecordIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddTaskRecordIDs(ids...)
	return uuo
}

// AddTaskRecords adds the "task_records" edges to the TaskRecord entity.
func (uuo *UserUpdateOne) AddTaskRecords(t ...*TaskRecord) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddTaskRecordIDs(ids...)
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (uuo *UserUpdateOne) AddTeamIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddTeamIDs(ids...)
	return uuo
}

// AddTeams adds the "teams" edges to the Team entity.
func (uuo *UserUpdateOne) AddTeams(t ...*Team) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddTeamIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (uuo *UserUpdateOne) ClearRoles() *UserUpdateOne {
	uuo.mutation.ClearRoles()
	return uuo
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (uuo *UserUpdateOne) RemoveRoleIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveRoleIDs(ids...)
	return uuo
}

// RemoveRoles removes "roles" edges to Role entities.
func (uuo *UserUpdateOne) RemoveRoles(r ...*Role) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.RemoveRoleIDs(ids...)
}

// ClearTaskRecords clears all "task_records" edges to the TaskRecord entity.
func (uuo *UserUpdateOne) ClearTaskRecords() *UserUpdateOne {
	uuo.mutation.ClearTaskRecords()
	return uuo
}

// RemoveTaskRecordIDs removes the "task_records" edge to TaskRecord entities by IDs.
func (uuo *UserUpdateOne) RemoveTaskRecordIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveTaskRecordIDs(ids...)
	return uuo
}

// RemoveTaskRecords removes "task_records" edges to TaskRecord entities.
func (uuo *UserUpdateOne) RemoveTaskRecords(t ...*TaskRecord) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveTaskRecordIDs(ids...)
}

// ClearTeams clears all "teams" edges to the Team entity.
func (uuo *UserUpdateOne) ClearTeams() *UserUpdateOne {
	uuo.mutation.ClearTeams()
	return uuo
}

// RemoveTeamIDs removes the "teams" edge to Team entities by IDs.
func (uuo *UserUpdateOne) RemoveTeamIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveTeamIDs(ids...)
	return uuo
}

// RemoveTeams removes "teams" edges to Team entities.
func (uuo *UserUpdateOne) RemoveTeams(t ...*Team) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveTeamIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if uuo.mutation.NicknameCleared() {
		_spec.ClearField(user.FieldNickname, field.TypeString)
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedStatus(); ok {
		_spec.AddField(user.FieldStatus, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.RegisteredAt(); ok {
		_spec.SetField(user.FieldRegisteredAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.RegisterIP(); ok {
		_spec.SetField(user.FieldRegisterIP, field.TypeString, value)
	}
	if uuo.mutation.RegisterIPCleared() {
		_spec.ClearField(user.FieldRegisterIP, field.TypeString)
	}
	if value, ok := uuo.mutation.LastLoggedAt(); ok {
		_spec.SetField(user.FieldLastLoggedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.LoginIP(); ok {
		_spec.SetField(user.FieldLoginIP, field.TypeString, value)
	}
	if uuo.mutation.LoginIPCleared() {
		_spec.ClearField(user.FieldLoginIP, field.TypeString)
	}
	if value, ok := uuo.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uuo.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if uuo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !uuo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.TaskRecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TaskRecordsTable,
			Columns: []string{user.TaskRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedTaskRecordsIDs(); len(nodes) > 0 && !uuo.mutation.TaskRecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TaskRecordsTable,
			Columns: []string{user.TaskRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.TaskRecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TaskRecordsTable,
			Columns: []string{user.TaskRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedTeamsIDs(); len(nodes) > 0 && !uuo.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
