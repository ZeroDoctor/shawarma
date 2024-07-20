package model

import (
	"time"
)

type Repository struct {
	UUID          UUID      `db:"uuid" json:"uuid,omitempty"`
	Owner         string    `db:"owner" json:"owner,omitempty"`
	Name          string    `db:"name" json:"name,omitempty"`
	DefaultBranch string    `db:"default_branch" json:"default_branch,omitempty"`
	CreatedAt     time.Time `db:"created_at" json:"created_at,omitempty"`
	ModifiedAt    time.Time `db:"modified_at" json:"modified_at,omitempty"`
	OwnerType     string    `db:"owner_type" json:"owner_type,omitempty"`
	OwnerID       UUID      `db:"owner_id" json:"owner_id,omitempty"`

	Branches     []Branch      `json:"branches,omitempty"`
	Environments []Environment `json:"environments,omitempty"`
	Env          map[string]string
}

type Branch struct {
	ID         int       `db:"id" json:"id,omitempty"`
	Name       string    `db:"name" json:"name,omitempty"`
	Hash       string    `json:"hash" db:"hash,omitempty"`
	CreatedAt  time.Time `db:"created_at" json:"created_at,omitempty"`
	ModifiedAt time.Time `db:"modified_at" json:"moified_at,omitempty"`

	RepoID UUID `db:"repo_id" json:"repo_id,omitempty"`

	Commits []Commit `json:"commits,omitempty"`
}

type Commit struct {
	Hash      string    `db:"commit" json:"hash,omitempty"`
	Author    string    `db:"author" json:"author,omitempty"`
	Message   string    `db:"message" json:"message,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`

	Parents []Commit `json:"commit,omitempty"`
}
