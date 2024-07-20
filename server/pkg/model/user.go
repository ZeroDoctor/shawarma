package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID       UUID       `json:"uuid,omitempty" db:"uuid"`
	Name       string     `json:"name,omitempty" db:"name"`
	Session    UUID       `json:"session,omitempty" db:"session"`
	AvatarURL  string     `json:"avatar_url,omitempty" db:"avatar_url"`
	GitRemote  StringList `json:"git_remote,omitempty" db:"git_remote"`
	CreatedAt  time.Time  `json:"created_at,omitempty" db:"created_at"`
	ModifiedAt time.Time  `json:"modified_at,omitempty" db:"modified_at"`

	Organizations []Organization `json:"organizations,omitempty"`
	Repositories  []Repository   `json:"repositories,omitempty"`
	Polls         []uuid.UUID
}

type Organization struct {
	UUID       UUID      `db:"uuid"`
	Owner      string    `db:"owner" json:"owner,omitempty"`
	Name       string    `db:"name" json:"name,omitempty"`
	AvatarURL  string    `db:"avatar_url" json:"avatar_url,omitempty"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`

	Repositories []Repository  `json:"repository,omitempty"`
	Environments []Environment `json:"environments,omitempty"`
	Env          map[string]string
}

type UserPoll struct {
	ID   UUID   `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
