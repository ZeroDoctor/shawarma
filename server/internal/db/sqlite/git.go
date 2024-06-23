package sqlite

import (
	"time"

	"github.com/google/uuid"
	"github.com/zerodoctor/shawarma/internal/model"
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

func (s *SqliteDB) InsertGithubAuthUser(user model.GithubUser) (model.GithubUser, error) {
	var err error

	insert := `INSERT INTO github_users (
		id, avatar_url, gravatar_id, "url",
		organizations_url, repos_url, "type",
		"name", created_at, updated_at
	) VALUES (
		:id, :avatar_url, :gravatar_id, :url,
		:organizations_url, :repos_url, :type,
		:name, :created_at, :updated_at
	);`

	_, err = s.conn.NamedExec(insert, user)
	return user, err
}

func (s *SqliteDB) InsertGithubUserOrgs(userID int, org model.GithubOrg) (model.GithubOrg, error) {
	var err error
	org, err = s.InsertGithubOrgs(org)
	if err != nil {
		return org, err
	}

	insert := `INSERT INTO github_users_orgs (
		github_user_id, github_org_id, created_at
	) VALUES (
		$1, $2, $3
	) ON CONFLICT (github_user_id, github_org_id) DO NOTHING;`

	_, err = s.conn.Exec(
		insert, userID, org.ID, model.Time(time.Now()),
	)
	return org, err
}

func (s *SqliteDB) InsertGithubOrgs(org model.GithubOrg) (model.GithubOrg, error) {
	var err error

	insert := `INSERT INTO github_orgs (
		id, "url", repos_url, hooks_url,
		issues_url, members_url, public_members_url,
		avatar_url, "description", "name", company,
		created_at, updated_at, archived_at, "type"
	) VALUES (
		:id, :url, :repos_url, :hooks_url,
		:issues_url, :members_url, :public_members_url,
		:avatar_url, :description, :name, :company,
		:created_at, :updated_at, :archived_at, :type
	) ON CONFLICT (id) DO NOTHING;`

	_, err = s.conn.NamedExec(insert, org)
	return org, err
}
