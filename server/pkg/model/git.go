package model

import (
	"time"
)

type Organization struct {
	UUID       UUID      `db:"uuid"`
	Owner      string    `db:"owner" json:"owner"`
	Name       string    `db:"name" json:"name"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`

	Repositories []Repository  `json:"repository"`
	Environments []Environment `json:"environments"`
	Env          map[string]string
}

type Repository struct {
	UUID       UUID      `db:"uuid"`
	Owner      string    `db:"owner" json:"owner"`
	Name       string    `db:"name" json:"name"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`

	OwnerType string `json:"owner_type"`
	OwnerID   UUID   `db:"owner_id"`

	Branches     []Branch      `json:"branches"`
	Environments []Environment `json:"environments"`
	Env          map[string]string
}

type Branch struct {
	ID         int       `db:"id"`
	Name       string    `db:"name" json:"name"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`

	LatestCommit string `db:"latest_commit" json:"latest_commit"`
	RepoID       UUID   `db:"repo_id"`

	Commits []Commit `json:"commits"`
}

type Commit struct {
	Hash      string    `db:"commit" json:"hash"`
	Author    string    `db:"author" json:"author"`
	CreatedAt time.Time `db:"created_at"`

	BranchID int `db:"branch_id"`
}
