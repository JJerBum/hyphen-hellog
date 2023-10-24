package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Like holds the schema definition for the Like entity.
type Like struct {
	ent.Schema
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return nil
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", Author.Type).
			Ref("likes").
			Unique(),

		edge.From("post", Post.Type).
			Ref("likes").
			Unique(),
	}
}
