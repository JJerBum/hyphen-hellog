package schema

import "entgo.io/ent"

// Vote holds the schema definition for the Vote entity.
type Vote struct {
	ent.Schema
}

// Fields of the Vote.
func (Vote) Fields() []ent.Field {
	return nil
}

// Edges of the Vote.
func (Vote) Edges() []ent.Edge {
	return nil
}
