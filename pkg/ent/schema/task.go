package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("type"),
		field.Int("status"),
		field.String("desc").Optional(),
		field.Time("created_at").Default(time.Now()),
		field.Time("completed_at").Optional(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("task_records", TaskRecord.Type),
		edge.From("workflow", Workflow.Type).
			Ref("tasks").
			Unique(),
		edge.From("team", Team.Type).
			Ref("tasks").
			Unique(),
	}
}
