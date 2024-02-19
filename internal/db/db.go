package db

import (
	"os"

	"github.com/jmoiron/sqlx"
)

type DB interface {
}

func LoadSchemaFromFile(db *sqlx.DB, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(data))
	return err
}
