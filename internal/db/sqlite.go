package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteDB struct {
	db *sqlx.DB
}

func NewSqliteDB() (*SqliteDB, error) {
	db, err := sqlx.Connect("sqlite3", "shawarma.db")

	if err := LoadSchemaFromFile(db, "./sql/sqlite/tables.sql"); err != nil {
		return nil, err
	}

	return &SqliteDB{db: db}, err
}
