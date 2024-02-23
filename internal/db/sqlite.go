package db

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zerodoctor/shawarma/internal/model"
)

const SQLITE_SCHEMA_DIR = "./sql/sqlite"

type SqliteDB struct {
	db  *sqlx.DB
	ctx context.Context
}

func NewSqliteDB(ctx context.Context) (*SqliteDB, error) {
	db, err := sqlx.ConnectContext(ctx, "sqlite3", "shawarma.db")
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

		return LoadSchemaFromFile(db, SQLITE_SCHEMA_DIR+"/"+path)
	})

	return &SqliteDB{db: db, ctx: ctx}, err
}

func (s *SqliteDB) InsertPipeline(pipeline model.Pipeline) (model.Pipeline, error) {
	insert := `INSERT INTO pipelines (
		"type", "status", created_at, modified_at
	) VALUES (
		:type, :status, :created_at, :modified_at
	) RETURNING id;`

	rows, err := s.db.NamedQuery(insert, pipeline)
	if err != nil {
		return pipeline, err
	}
	defer rows.Close()

	id := -1
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return pipeline, err
		}
		log.Infof("create new pipeline for [repo=%s] with [id=%d]",
			pipeline.Repo.Name, id,
		)
	}
	pipeline.ID = id

	return pipeline, nil
}

func (s *SqliteDB) InsertStep(step model.Step) (model.Step, error) {
	var err error
	step.UUID, err = uuid.NewV7()
	if err != nil {
		return step, err
	}

	insert := `INSERT INTO steps (
		uuid, "name", "image", 
		commands, privileged, detach,
		created_at, modified_at, pipeline_id
	) VALUES (
		:uuid, :name, :image, 
		:commands, :privileged, :detach,
		:created_at, :modified_at, :pipeline_id
	) ON CONFLICT ("name", pipeline_id) DO UPDATE SET 
		uuid        = excluded.uuid, 
		"name"      = excluded.name, 
		"image"     = excluded.image, 
		commands    = excluded.commands, 
		privileged  = excluded.privileged, 
		detach      = excluded.detach,
		created_at  = excluded.created_at, 
		modified_at = excluded.modified_at, 
		pipeline_id = excluded.pipeline_id
	;`

	_, err = s.db.NamedExec(insert, step)
	return step, err
}

func (s *SqliteDB) InsertEvent(event model.Event) (model.Event, error) {
	insert := `INSERT INTO events (
		webhook, "type", "action",
		deadline, created_at, modified_at,
		pipeline_id, step_id
	) VALUES (
		:webhook, :type, :action,
		:deadline, :created_at, :modified_at,
		:pipeline_id, :step_id
	) ON CONFLICT ("type", pipeline_id, step_id) DO UPDATE SET 
		webhook     = excluded.webhook, 
		"type"      = excluded.type, 
		"action"    = excluded.action,
		deadline    = excluded.deadline, 
		created_at  = excluded.created_at, 
		modified_at = excluded.modified_at,
		pipeline_id = excluded.pipline_id, 
		step_id     = excluded.step_id
	;`

	_, err := s.db.NamedExec(insert, event)
	return event, err
}

func (s *SqliteDB) InsertEnvironment(environment model.Environment) (model.Environment, error) {
	return environment, nil
}

func (s *SqliteDB) InsertOrganization(organization model.Organization) (model.Organization, error) {
	return organization, nil
}

func (s *SqliteDB) InsertRepository(repository model.Repository) (model.Repository, error) {
	return repository, nil
}

func (s *SqliteDB) InsertBranch(branch model.Branch) (model.Branch, error) {
	return branch, nil
}

func (s *SqliteDB) InsertCommit(commit model.Commit) (model.Commit, error) {
	return commit, nil
}
