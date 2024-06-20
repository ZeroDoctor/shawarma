package sqlite

import (
	"github.com/google/uuid"
	"github.com/zerodoctor/shawarma/internal/model"
)

func (s *SqliteDB) InsertPipeline(pipeline model.Pipeline) (model.Pipeline, error) {
	insert := `INSERT INTO pipelines (
		"type", "status", 
		created_at, modified_at,
		repo_id, runner_id
	) VALUES (
		:type, :status, 
		:created_at, :modified_at,
		:repo_id, :runner_id
	) RETURNING id;`

	rows, err := s.conn.NamedQuery(insert, pipeline)
	if err != nil {
		return pipeline, err
	}
	defer rows.Close()

	id := -1
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return pipeline, err
		}
	}
	pipeline.ID = id

	for i := range pipeline.Steps {
		pipeline.Steps[i].PipelineID = pipeline.ID
		pipeline.Steps[i], err = s.InsertStep(pipeline.Steps[i])
		if err != nil {
			return pipeline, err
		}
	}

	for i := range pipeline.Events {
		pipeline.Events[i].PipelineID = pipeline.ID
		pipeline.Events[i], err = s.InsertEvent(pipeline.Events[i])
		if err != nil {
			return pipeline, err
		}
	}

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

	if _, err = s.conn.NamedExec(insert, step); err != nil {
		return step, err
	}

	for i := range step.Events {
		step.Events[i].StepID = step.UUID
		step.Events[i], err = s.InsertEvent(step.Events[i])
		if err != nil {
			return step, err
		}
	}

	return step, nil
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

	_, err := s.conn.NamedExec(insert, event)
	return event, err
}

func (s *SqliteDB) InsertEnvironment(environment model.Environment) (model.Environment, error) {
	insert := `INSERT INTO environments (
		"key", "data", protected, 
		created_at, modified_at, 
		repo_id, org_id
	) VALUES (
		:key, :data, :protected, 
		:created_at, :modified_at, 
		:repo_id, :org_id
	) ON CONFLICT ("key", repo_id, org_id) DO UPDATE SET 
		"key"       = excluded.key, 
		"data"      = excluded.data, 
		protected   = excluded.protected, 
		created_at  = excluded.created_at, 
		modified_at = excluded.modified_at, 
		repo_id     = excluded.repo_id, 
		org_id      = excluded.org_id
	;`

	_, err := s.conn.NamedExec(insert, environment)
	return environment, err
}
