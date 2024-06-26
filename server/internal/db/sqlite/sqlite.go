package sqlite

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"github.com/zerodoctor/shawarma/internal/db"
	"github.com/zerodoctor/shawarma/internal/logger"
)

var log *logrus.Logger = logger.Log

var ErrModelConvert error = errors.New("failed to convert result to model")

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

func (s *SqliteDB) GetConnection() *sqlx.DB { return s.conn }

func (s *SqliteDB) GetType() string { return "sqlite3" }

type Time time.Time

func (t *Time) Scan(src interface{}) error {
	if unix, ok := src.(int64); ok {
		*t = Time(time.Unix(unix, 0))
		return nil
	}

	return fmt.Errorf("unexcepted format [src=%T] wanted type int64", src)
}

func (t Time) Value() (driver.Value, error) {
	ti := time.Time(t)
	return ti.Unix(), nil
}

func convertNamedSqlite(object interface{}) map[string]interface{} {
	if reflect.TypeOf(object).Kind() != reflect.Struct {
		return map[string]interface{}{}
	}

	named := make(map[string]interface{})

	vObject := reflect.ValueOf(object)
	tObject := vObject.Type()

	for i := 0; i < tObject.NumField(); i++ {
		value := vObject.Field(i).Interface()

		tag := tObject.Field(i).Tag.Get("db")
		switch tag {
		case "created_at", "modified_at":
			value = Time(value.(time.Time))
		}
		named[tObject.Field(i).Tag.Get("db")] = value
	}

	return named
}

func convertModel(objectMap map[string]interface{}, object interface{}) interface{} {
	if reflect.TypeOf(object).Kind() != reflect.Struct {
		return object
	}

	vObject := reflect.ValueOf(object)
	tObject := vObject.Type()

	for i := 0; i < vObject.NumField(); i++ {
		tag := tObject.Field(i).Tag.Get("db")
		value := objectMap[tag]

		switch tag {
		case "created_at", "modified_at":
			value = time.Time(value.(Time))
		}
		vObject.Field(i).Set(reflect.ValueOf(value))
	}

	return object
}
