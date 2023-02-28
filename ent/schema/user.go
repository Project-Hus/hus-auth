package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uuid").Unique(),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
		field.String("google_sub").Unique(),

		field.String("email").Unique(),
		field.Bool("email_verified"),

		// User real info
		field.String("name"),
		field.Time("birthdate").Nillable(),
		field.String("given_name"),
		field.String("family_name"),

		// User Info in the service
		field.Text("google_profile_picture"),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("groups", Community.Type).Ref("users"),
	}
}
