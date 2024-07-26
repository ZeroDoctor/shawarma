package db

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/zerodoctor/shawarma/internal/logger"
	"github.com/zerodoctor/shawarma/pkg/model"
)

var log = logger.Log

type DB interface {
	GetConnection() *sqlx.DB
	GetType() string

	SavePipeline(model.Pipeline) (model.Pipeline, error)
	SaveStep(model.Step) (model.Step, error)
	SaveEvent(model.Event) (model.Event, error)
	SaveEnvironment(model.Environment) (model.Environment, error)
	SaveOrganization(model.Organization) (model.Organization, error)
	SaveRepository(model.Repository) (model.Repository, error)
	SaveBranch(model.Branch) (model.Branch, error)
	SaveCommit(model.Commit) (model.Commit, error)
	SaveRunner(model.Runner) (model.Runner, error)
	SaveUser(model.User) (model.User, error)

	QueryUserByName(string) (model.User, error)
	QueryUserCount() (int, error)
}

func LoadSchemaFromFile(db *sqlx.DB, path string) error {
	log.Debugf("loading schema [file=%s]...", path)

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(data))
	return err
}
