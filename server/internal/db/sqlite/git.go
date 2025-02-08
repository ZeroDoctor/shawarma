package sqlite

import (
	"github.com/google/uuid"
	"github.com/zerodoctor/shawarma/pkg/model"
)

func (s *SqliteDB) SaveRepository(repository model.Repository) (model.Repository, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return repository, err
	}
	repository.UUID = model.UUID(id)

	insert := `INSERT INTO repositories (
		uuid, "owner", "name", 
		default_branch, owner_type,
		owner_id, created_at, modified_at
	) VALUES (
		:uuid, :owner, :name, 
		:default_branch, :owner_type,
		:owner_id, :created_at, :modified_at
	) ON CONFLICT ("owner", "name") DO UPDATE SET 
		uuid           = excluded.uuid, 
		"owner"        = excluded.owner, 
		"name"         = excluded.name, 
		default_branch = excluded.default_branch,
 		owner_type     = excluded.owner_type,
		created_at     = excluded.created_at, 
		modified_at    = excluded.modified_at
	;`

	if _, err = s.conn.NamedExec(insert, convertNamedSqlite(repository)); err != nil {
		return repository, err
	}

	for i := range repository.Branches {
		repository.Branches[i].RepoID = repository.UUID
		repository.Branches[i], err = s.SaveBranch(repository.Branches[i])
		if err != nil {
			return repository, err
		}
	}

	for i := range repository.Environments {
		repository.Environments[i].RepoID = repository.UUID
		repository.Environments[i], err = s.SaveEnvironment(repository.Environments[i])
		if err != nil {
			return repository, err
		}
	}

	return repository, nil
}

func (s *SqliteDB) GetAllUserRepos(user model.User) ([]model.Repository, error) {
	var repos []model.Repository
	query := `SELECT * FROM repositories where "owner" = ?;`

	rows, err := s.conn.Queryx(query, user.Name)
	if err != nil {
		return repos, err
	}
	defer rows.Close()

	repoMap := make(map[string]interface{})
	for rows.Next() {
		if err := rows.MapScan(repoMap); err != nil {
			return repos, err
		}

		repo := model.Repository{}
		convertModel(repoMap, &repo)
		repos = append(repos, repo)
	}

	return repos, err
}

func (s *SqliteDB) SaveBranch(branch model.Branch) (model.Branch, error) {
	insert := `INSERT INTO branches (
		id, "name", "hash", created_at,
		modified_at, repo_id
	) VALUES (
		:id, :name, :hash, :created_at,
		:modified_at, :repo_id
	) RETURNING id;`

	rows, err := s.conn.NamedQuery(insert, convertNamedSqlite(branch))
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
		branch.Commits[i], err = s.SaveCommit(branch.Commits[i])
		if err != nil {
			return branch, err
		}
	}

	return branch, nil
}

func (s *SqliteDB) GetBranches(repoID model.UUID) ([]model.Branch, error) {
	var branches []model.Branch
	query := `SELECT * FROM branches WHERE repo_id = ?;`

	err := s.conn.Select(&branches, query, repoID)
	return branches, err
}

func (s *SqliteDB) SaveCommit(commit model.Commit) (model.Commit, error) {
	insert := `INSERT INTO commits (
		"hash", author, created_at, branch_id
	) VALUES (
		:hash, :author, :created_at, :branch_id
	);`

	_, err := s.conn.NamedExec(insert, convertNamedSqlite(commit))
	return commit, err
}
