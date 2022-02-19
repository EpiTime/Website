package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	uuid "github.com/satori/go.uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			Unique(),
		field.String("password"),
		field.Int("year").
			Default(2026),
		field.String("hideModules").
			Optional(),
		field.String("othersModules").
			Optional(),
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.NewV4).
			Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
