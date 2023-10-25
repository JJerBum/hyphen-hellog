package database

import (
	"context"
	"testing"
)

func TestDatabase(t *testing.T) {
	// New().CreateAuthor(context.Background(), &ent.Author{AuthorID: 1})
	New().DeleteAuthor(context.Background(), 1)

}
