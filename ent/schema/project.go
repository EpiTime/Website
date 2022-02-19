package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("Start").Optional(),
		field.String("end").Optional(),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("projects").
			Unique(),
	}
}
