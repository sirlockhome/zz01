package datastore

import (
	"foxomni/pkg/database"
)

type Datastore struct {
	sql *database.SQL
}

func NewDatastore(sql *database.SQL) *Datastore {
	return &Datastore{
		sql: sql,
	}
}

const (
	ErrProductIDNotFound = "product with id `%d` not found"
)
