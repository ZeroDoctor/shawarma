package sqlite

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zerodoctor/shawarma/internal/db"
)

const SQLITE_SCHEMA_DIR = "./sql/sqlite"

type SqliteDB struct {
	conn *sqlx.DB
	ctx  context.Context
}

func NewDB(ctx context.Context) (*SqliteDB, error) {
	conn, err := sqlx.ConnectContext(ctx, "sqlite3", "shawarma.db")
	if err != nil {
		return nil, err
	}

	err = filepath.Walk(SQLITE_SCHEMA_DIR, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && !strings.Contains(path, ".schema.") {
			return nil
		}

		return db.LoadSchemaFromFile(conn, SQLITE_SCHEMA_DIR+"/"+path)
	})

	return &SqliteDB{conn: conn, ctx: ctx}, err
}
