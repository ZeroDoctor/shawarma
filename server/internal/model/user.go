package model

import "github.com/google/uuid"

type User struct {
	UUID         uuid.UUID `json:"uuid" db:"uuid"`
	Name         string    `json:"name" db:"name"`
	GithubToken  string    `json:"github_token" db:"github_token"`
	GithubState  string    `json:"github_state" db:"github_state"`
	GithubUserID int       `json:"github_user_id" db:"github_user_id"`
	Session      string    `json:"session" db:"session"`
	CreatedAt    Time      `json:"created_at" db:"created_at"`
	ModifiedAt   Time      `json:"modified_at" db:"modified_at"`

	GithubUser GithubUser
}
