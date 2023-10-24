package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Author holds the schema definition for the Author entity.
type Author struct {
	ent.Schema
}

// Fields of the Author.
func (Author) Fields() []ent.Field {
	return []ent.Field{
		field.Int("author_id").
			Unique(),
		field.Time("joined_at").
			Default(time.Now),
	}
}

// Edges of the Author.
func (Author) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("comments", Comment.Type),

		edge.To("likes", Like.Type),
	}
}
