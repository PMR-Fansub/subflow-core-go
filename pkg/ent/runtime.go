// Code generated by ent, DO NOT EDIT.

package ent

import (
	"subflow-core-go/pkg/ent/role"
	"subflow-core-go/pkg/ent/schema"
	"subflow-core-go/pkg/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescName is the schema descriptor for name field.
	roleDescName := roleFields[1].Descriptor()
	// role.NameValidator is a validator for the "name" field. It is called by the builders before save.
	role.NameValidator = roleDescName.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescRegisteredAt is the schema descriptor for registered_at field.
	userDescRegisteredAt := userFields[5].Descriptor()
	// user.DefaultRegisteredAt holds the default value on creation for the registered_at field.
	user.DefaultRegisteredAt = userDescRegisteredAt.Default.(time.Time)
	// userDescLastLoggedAt is the schema descriptor for last_logged_at field.
	userDescLastLoggedAt := userFields[7].Descriptor()
	// user.DefaultLastLoggedAt holds the default value on creation for the last_logged_at field.
	user.DefaultLastLoggedAt = userDescLastLoggedAt.Default.(time.Time)
}
