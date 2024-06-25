package sqlite

import (
	"github.com/google/uuid"
	"github.com/zerodoctor/shawarma/pkg/model"
)

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
