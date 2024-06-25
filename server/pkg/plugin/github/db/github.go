package db

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/zerodoctor/shawarma/pkg/plugin/github/model"
)

type DB struct {
	conn *sqlx.DB
}

func (s *DB) InsertGithubAuthUser(user model.GithubUser) (model.GithubUser, error) {
	var err error

	insert := `INSERT INTO github_users (
		id, avatar_url, gravatar_id, "url",
		organizations_url, repos_url, "type",
		"name", created_at, updated_at
	) VALUES (
		:id, :avatar_url, :gravatar_id, :url,
		:organizations_url, :repos_url, :type,
		:name, :created_at, :updated_at
	);` // TODO: change to upsert

	_, err = s.conn.NamedExec(insert, user)
	return user, err
}

func (s *DB) InsertGithubUserOrg(userID int, org model.GithubOrg) (model.GithubOrg, error) {
	var err error
	org, err = s.InsertGithubOrgs(org)
	if err != nil {
		return org, err
	}

	insert := `INSERT INTO github_users_orgs (
		user_id, org_id, created_at
	) VALUES (
		$1, $2, $3
	) ON CONFLICT (user_id, org_id) DO NOTHING;`

	_, err = s.conn.Exec(
		insert, userID, org.ID, model.Time(time.Now()),
	)
	return org, err
}

func (s *DB) InsertGithubOrgs(org model.GithubOrg) (model.GithubOrg, error) {
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
	) ON CONFLICT (id) DO NOTHING;` // TODO: change to upsert

	_, err = s.conn.NamedExec(insert, org)
	return org, err
}

func (s *DB) InsertGithubOwner(owner model.GithubOwner) (model.GithubOwner, error) {
	var err error

	insert := `INSERT INTO github_owners (
		id, avatar_url, gravatar_id, "url",
		organizations_url, repos_url, "type"
	) VALUES (
		:id, :avatar_url, :gravatar_id, :url,
		:organizations_url, :repos_url, :type
	) ON CONFLICT (id) DO NOTHING;` // TODO: change to upsert for avatar

	_, err = s.conn.NamedExec(insert, owner)
	return owner, err
}

func (s *DB) InsertGithubRepo(repo model.GithubRepo) (model.GithubRepo, error) {
	var err error

	owner, err := s.InsertGithubOwner(repo.Owner)
	if err != nil {
		return repo, err
	}
	repo.OwnerID = owner.ID

	insert := `INSERT INTO github_repos (
		id, owner_id, "name", full_name, "description",
		"url", collaborators_url, hooks_url, issue_events_url,
		branches_url, tags_url, statuses_url, commits_url,
		merges_url, issues_url, pulls_url, created_at, updated_at,
		pushed_at, has_issues, archived, open_issues_count, visibility
	) VALUES (
		:id, :owner_id, :name, :full_name, :description,
		:url, :collaborators_url, :hooks_url, :issue_events_url,
		:branches_url, :tags_url, :statuses_url, :commits_url,
		:merges_url, :issues_url, :pulls_url, :created_at, :updated_at,
		:pushed_at, :has_issues, :archived, :open_issues_count, :visibility
	) ON CONFLICT (id) DO UPDATE SET
		owner_id          = excluded.owner_id,
		"name"            = excluded.name,
		full_name         = excluded.full_name,
		"description"     = excluded.description,
		"url"             = excluded.url,
		collaborators_url = excluded.collaborators_url,
		hooks_url         = excluded.hooks_url,
		issue_events_url  = excluded.issue_events_url,
		branches_url      = excluded.branches_url,
		tags_url          = excluded.tags_url,
		statuses_url      = excluded.statuses_url,
		commits_url       = excluded.commits_url,
		merges_url        = excluded.merges_url,
		issues_url        = excluded.issues_url,
		pulls_url         = excluded.pulls_url,
		created_at        = excluded.created_at,
		updated_at        = excluded.updated_at,
		pushed_at         = excluded.pushed_at,
		has_issues        = excluded.has_issues,
		archived          = excluded.archived,
		open_issues_count = excluded.open_issues_count,
		visibility        = excluded.visibility;`

	_, err = s.conn.NamedExec(insert, repo)
	return repo, err
}
