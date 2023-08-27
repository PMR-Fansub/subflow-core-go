package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// WorkflowNode holds the schema definition for the WorkflowNode entity.
type WorkflowNode struct {
	ent.Schema
}

// Fields of the WorkflowNode.
func (WorkflowNode) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("type"),
		field.String("desc").Optional(),
		field.Int("seq").Optional(),
	}
}

// Edges of the WorkflowNode.
func (WorkflowNode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("task_records", TaskRecord.Type),
		edge.From("workflow", Workflow.Type).
			Ref("workflow_nodes").
			Unique(),
	}
}
