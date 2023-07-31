package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.String("email").NotEmpty(),
		field.String("nickname").Optional(),
		field.Int("status"),
		field.Time("registered_at").Default(time.Now()),
		field.String("register_ip").Optional(),
		field.Time("last_logged_at").Default(time.Now()),
		field.String("login_ip").Optional(),
		field.String("avatar").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
