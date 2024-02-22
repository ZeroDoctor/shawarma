package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const SQLITE_SCHEMA = "./sql/sqlite/schema.sql"

type SqliteDB struct {
	db *sqlx.DB
}

func NewSqliteDB() (*SqliteDB, error) {
	db, err := sqlx.Connect("sqlite3", "shawarma.db")

	if err := LoadSchemaFromFile(db, SQLITE_SCHEMA); err != nil {
		return nil, err
	}

	return &SqliteDB{db: db}, err
}
