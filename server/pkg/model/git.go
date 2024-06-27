package model

import (
	"time"
)

type Repository struct {
	UUID       UUID      `db:"uuid" json:"uuid"`
	Owner      string    `db:"owner" json:"owner"`
	Name       string    `db:"name" json:"name"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	ModifiedAt time.Time `db:"modified_at" json:"modified_at"`

	OwnerType string `db:"owner_type" json:"owner_type"`
	OwnerID   UUID   `db:"owner_id" json:"owner_id"`

	Branches     []Branch      `json:"branches"`
	Environments []Environment `json:"environments"`
	Env          map[string]string
}

type Branch struct {
	ID         int       `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	ModifiedAt time.Time `db:"modified_at" json:"moified_at"`

	RepoID UUID `db:"repo_id" json:"repo_id"`

	Commits []Commit `json:"commits"`
}

type Commit struct {
	Hash      string    `db:"commit" json:"hash"`
	Author    string    `db:"author" json:"author"`
	Message   string    `db:"message" json:"message"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`

	BranchID int `db:"branch_id" json:"branch_id"`
}
