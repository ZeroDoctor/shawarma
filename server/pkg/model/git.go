package model

import (
	"github.com/google/uuid"
)

type Organization struct {
	UUID       uuid.UUID `db:"uuid"`
	Owner      string    `db:"owner" json:"owner"`
	Name       string    `db:"name" json:"name"`
	CreatedAt  Time      `db:"created_at"`
	ModifiedAt Time      `db:"modified_at"`

	Repositories []Repository  `json:"repository"`
	Environments []Environment `json:"environments"`
	Env          map[string]string
}

type Repository struct {
	UUID       uuid.UUID `db:"uuid"`
	Owner      string    `db:"owner" json:"owner"`
	Name       string    `db:"name" json:"name"`
	CreatedAt  Time      `db:"created_at"`
	ModifiedAt Time      `db:"modified_at"`

	OrgID uuid.UUID `db:"org_id"`

	Branches     []Branch      `json:"branches"`
	Environments []Environment `json:"environments"`
	Env          map[string]string
}

type Branch struct {
	ID         int    `db:"id"`
	Name       string `db:"name" json:"name"`
	CreatedAt  Time   `db:"created_at"`
	ModifiedAt Time   `db:"modified_at"`

	LatestCommit string    `db:"latest_commit" json:"latest_commit"`
	RepoID       uuid.UUID `db:"repo_id"`

	Commits []Commit `json:"commits"`
}

type Commit struct {
	Hash      string `db:"commit" json:"hash"`
	Author    string `db:"author" json:"author"`
	CreatedAt Time   `db:"created_at"`

	BranchID int `db:"branch_id"`
}
