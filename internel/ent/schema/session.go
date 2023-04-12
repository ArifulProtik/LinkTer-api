package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.Text("token").NotEmpty(),
		field.String("ip").NotEmpty(),
		field.Time("started_time").Default(time.Now),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return nil
}
