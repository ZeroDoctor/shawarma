package model

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	UUID       uuid.UUID `db:"uuid"`
	Owner      string    `db:"owner"`
	Name       string    `db:"name"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`

	Environments []Environment
	Env          map[string]string
}

type Repository struct {
	UUID       uuid.UUID `db:"uuid"`
	Owner      string    `db:"owner"`
	Name       string    `db:"name"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`

	Branches     []Branch
	Environments []Environment
	Env          map[string]string
}

type Branch struct {
	ID         int       `db:"id"`
	Name       string    `db:"name"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`

	LatestCommit string `db:"latest_commit"`
	RepoID       string `db:"repo_id"`

	Commits []Commit
}

type Commit struct {
	Hash      string    `db:"commit"`
	Author    string    `db:"author"`
	CreatedAt time.Time `db:"created_at"`

	BranchID int `db:"branch_id"`
}
