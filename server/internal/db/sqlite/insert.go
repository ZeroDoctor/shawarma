package sqlite

import (
	"github.com/google/uuid"
	"github.com/zerodoctor/shawarma/internal/model"
)

// TODO: use transaction when possible

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

func (s *SqliteDB) InsertOrganization(organization model.Organization) (model.Organization, error) {
	var err error
	organization.UUID, err = uuid.NewV7()
	if err != nil {
		return organization, err
	}

	insert := `INSERT INTO organizations (
		uuid, "owner", "name", created_at, modified_at
	) VALUES (
		:uuid, :owner, :name, :created_at, :modified_at
	) ON CONFLICT ("owner", "name") DO UPDATE SET 
		uuid        = excluded.uuid, 
		"owner"     = excluded.owner, 
		"name"      = excluded.name, 
		created_at  = excluded.created_at, 
		modified_at = excluded.modified_at
	;`

	if _, err = s.conn.NamedExec(insert, organization); err != nil {
		return organization, err
	}

	for i := range organization.Repositories {
		organization.Repositories[i].OrgID = organization.UUID
		organization.Repositories[i], err = s.InsertRepository(organization.Repositories[i])
		if err != nil {
			return organization, err
		}
	}

	for i := range organization.Environments {
		organization.Environments[i].OrgID = organization.UUID
		organization.Environments[i], err = s.InsertEnvironment(organization.Environments[i])
		if err != nil {
			return organization, err
		}
	}

	return organization, nil
}

func (s *SqliteDB) InsertRepository(repository model.Repository) (model.Repository, error) {
	var err error
	repository.UUID, err = uuid.NewV7()
	if err != nil {
		return repository, err
	}

	insert := `INSERT INTO repositories (
		uuid, "owner", "name", 
		created_at, modified_at,
		org_id
	) VALUES (
		:uuid, :owner, :name, 
		:created_at, :modified_at,
		:org_id
	) ON CONFLICT ("owner", "name") DO UPDATE SET 
		uuid        = excluded.uuid, 
		"owner"     = excluded.owner, 
		"name"      = excluded.name, 
		created_at  = excluded.created_at, 
		modified_at = excluded.modified_at,
		org_id      = excluded.org_id
	;`

	if _, err = s.conn.NamedExec(insert, repository); err != nil {
		return repository, err
	}

	for i := range repository.Branches {
		repository.Branches[i].RepoID = repository.UUID
		repository.Branches[i], err = s.InsertBranch(repository.Branches[i])
		if err != nil {
			return repository, err
		}
	}

	for i := range repository.Environments {
		repository.Environments[i].RepoID = repository.UUID
		repository.Environments[i], err = s.InsertEnvironment(repository.Environments[i])
		if err != nil {
			return repository, err
		}
	}

	return repository, nil
}

func (s *SqliteDB) InsertBranch(branch model.Branch) (model.Branch, error) {
	insert := `INSERT INTO branches (
		id, "name", created_at,
		modified_at, latest_commit,
		repo_id
	) VALUES (
		:id, :name, :created_at,
		:modified_at, :latest_commit,
		:repo_id
	) RETURNING id;`

	rows, err := s.conn.NamedQuery(insert, branch)
	if err != nil {
		return branch, err
	}
	defer rows.Close()

	id := -1
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return branch, err
		}
	}
	branch.ID = id

	for i := range branch.Commits {
		branch.Commits[i].BranchID = branch.ID
		branch.Commits[i], err = s.InsertCommit(branch.Commits[i])
		if err != nil {
			return branch, err
		}
	}

	return branch, nil
}

func (s *SqliteDB) InsertCommit(commit model.Commit) (model.Commit, error) {
	insert := `INSERT INTO commits (
		"hash", author, created_at, branch_id
	) VALUES (
		:hash, :author, :created_at, :branch_id
	);`

	_, err := s.conn.NamedExec(insert, commit)
	return commit, err
}

func (s *SqliteDB) InsertRunner(runner model.Runner) (model.Runner, error) {
	var err error
	runner.UUID, err = uuid.NewV7()
	if err != nil {
		return runner, err
	}

	insert := `INSERT INTO runners (
		uuid, "type", hostname, 
		created_at, modified_at
	) VALUES (
		:uuid, :type, :hostname, 
		:created_at, :modified_at
	) ON CONFLICT (hostname) DO UPDATE SET 
		uuid        = excluded.uuid, 
		"type"      = excluded.type, 
		hostname    = excluded.hostname, 
		created_at  = excluded.created_at, 
		modified_at = excluded.modified_at
	;`

	_, err = s.conn.NamedExec(insert, runner)
	return runner, err
}

func (s *SqliteDB) InsertGithubUser(user model.User) (model.User, error) {
	var err error

	insert := `INSERT INTO users (
		uuid, "name", github_token,
		"session", created_at, modified_at
	) VALUES (
		:uuid, :name, :github_token,
		:session, :created_at, :modified_at
	) ON CONFLICT("name", "session") DO UPDATE SET
		uuid         = excluded.uuid, 
		github_token = excluded.github_token,
		session      = excluded.session,
		created_at   = excluded.created_at, 
		modified_at  = excluded.modified_at
	;`

	_, err = s.conn.NamedExec(insert, user)
	return user, err
}
