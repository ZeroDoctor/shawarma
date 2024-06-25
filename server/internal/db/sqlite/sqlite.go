package sqlite

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"github.com/zerodoctor/shawarma/internal/db"
	"github.com/zerodoctor/shawarma/internal/logger"
)

var log *logrus.Logger = logger.Log

const SQLITE_SCHEMA_DIR = "./server/sql/sqlite"

type SqliteDB struct {
	conn *sqlx.DB
	ctx  context.Context
}

func NewConnection(ctx context.Context, dbName string) (*SqliteDB, error) {
	log.Info("connecting to sqlite db...")
	conn, err := sqlx.ConnectContext(ctx, "sqlite3", dbName)
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

		return db.LoadSchemaFromFile(conn, path)
	})

	return &SqliteDB{conn: conn, ctx: ctx}, err
}
