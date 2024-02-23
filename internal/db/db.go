package db

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/zerodoctor/shawarma/internal/logger"
	"github.com/zerodoctor/shawarma/internal/model"
)

var log = logger.Log

type DB interface {
	InsertPipeline(model.Pipeline) (model.Pipeline, error)
	InsertStep(model.Step) (model.Step, error)
	InsertEvent(model.Event) (model.Event, error)
	InsertEnvironment(model.Environment) (model.Environment, error)
	InsertOrganization(model.Organization) (model.Organization, error)
	InsertRepository(model.Repository) (model.Repository, error)
	InsertBranch(model.Branch) (model.Branch, error)
	InsertCommit(model.Commit) (model.Commit, error)
}

func LoadSchemaFromFile(db *sqlx.DB, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(data))
	return err
}
