package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TaskRecord holds the schema definition for the TaskRecord entity.
type TaskRecord struct {
	ent.Schema
}

// Fields of the TaskRecord.
func (TaskRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int("status"),
		field.Time("created_at").Default(time.Now()),
		field.Time("assigned_at").Optional(),
		field.Time("completed_at").Optional(),
		field.String("remark").Optional(),
	}
}

// Edges of the TaskRecord.
func (TaskRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("task_records").
			Unique(),
		edge.From("task", Task.Type).
			Ref("task_records").
			Unique(),
		edge.From("workflow_node", WorkflowNode.Type).
			Ref("task_records").
			Unique(),
	}
}
