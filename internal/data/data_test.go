package data

import (
	"testing"

	"usm/internal/data/ent/enttest"

	_ "github.com/mattn/go-sqlite3"
)

func NewTestData(t *testing.T) (*Data, func()) {
	db := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	return &Data{db}, func() {
		db.Close()
	}
}
