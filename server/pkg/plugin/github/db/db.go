package db

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/zerodoctor/shawarma/pkg/plugin/github/model"
)

type DB struct {
	dbType string
	conn   *sqlx.DB
}

func NewDB(dbType string, conn *sqlx.DB) *DB {
	return &DB{
		dbType: dbType,
		conn:   conn,
	}
}

func (s *DB) SaveGithubAuthUser(user model.GithubUser) (model.GithubUser, error) {
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

func (s *DB) SaveGithubUserOrg(userID int, org model.GithubOrg) (model.GithubOrg, error) {
	var err error
	org, err = s.SaveGithubOrgs(org)
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

func (s *DB) SaveGithubOrgs(org model.GithubOrg) (model.GithubOrg, error) {
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
	) ON CONFLICT (id) DO UPDATE SET
		"url"              = excluded.url,
		repos_url          = excluded.repos_url,
		hooks_url          = excluded.hooks_url,
		issues_url         = excluded.issues_url,
		members_url        = excluded.members_url,
		public_members_url = excluded.public_members_url,
		avatar_url         = excluded.avatar_url,
		"description"      = excluded.description,
		"name"             = excluded.name,
		company            = excluded.company,
		created_at         = excluded.created_at,
		updated_at         = excluded.updated_at,
		archived_at        = excluded.archived_at,
		"type"             = excluded.type;`

	_, err = s.conn.NamedExec(insert, org)
	return org, err
}

func (s *DB) SaveGithubOwner(owner model.GithubOwner) (model.GithubOwner, error) {
	var err error

	insert := `INSERT INTO github_owners (
		id, avatar_url, gravatar_id, "url",
		organizations_url, repos_url, "type"
	) VALUES (
		:id, :avatar_url, :gravatar_id, :url,
		:organizations_url, :repos_url, :type
	) ON CONFLICT (id) DO UPDATE SET 
		avatar_url        = excluded.avatar_url, 
		gravatar_id       = excluded.gravatar_id,
		"url"             = excluded.url,
		organizations_url = excluded.organizations_url,
		repos_url         = excluded.repos_url,
		"type"            = excluded.type;`

	_, err = s.conn.NamedExec(insert, owner)
	return owner, err
}

func (s *DB) SaveGithubRepo(repos []model.GithubRepo) ([]model.GithubRepo, error) {
	var err error

	for i := range repos {
		owner, err := s.SaveGithubOwner(repos[i].Owner)
		if err != nil {
			return repos, err
		}
		repos[i].OwnerID = owner.ID
	}

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

	_, err = s.conn.NamedExec(insert, repos)
	return repos, err
}

func (s *DB) SaveGithubBranch(branch []model.GithubBranch) ([]model.GithubBranch, error) {

	return branch, nil
}

func (s *DB) SaveGithubCommit(commits []model.GithubCommit) ([]model.GithubCommit, error) {
	insert := `INSERT INTO github_commits (
		sha, "message", "url"
	) VALUES (
		:sha, :message, :url
	) ON CONFLICT (sha) DO NOTHING;`

	_, err := s.conn.NamedExec(insert, commits)
	if err != nil {
		return commits, err
	}

	for i := range commits {
		parentCommits, err := s.SaveGithubCommit(commits[i].Parents)
		if err != nil {
			return commits, err
		}

		for j := range parentCommits {
			err = s.SaveGithubCommitParents(parentCommits[j].SHA, commits[i].SHA)
			if err != nil {
				return commits, err
			}
		}
	}

	return commits, err
}

func (s *DB) SaveGithubCommitParents(parentSHA string, childSHA string) error {
	insert := `INSERT INTO github_commit_parents (
		parent_sha, child_sha
	) VALUES (
		$1, $2
	) ON CONFLICT (parent_sha, child_sha) DO NOTHING;`

	_, err := s.conn.Exec(insert, parentSHA, childSHA)
	return err
}
