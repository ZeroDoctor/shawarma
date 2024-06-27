package model

import (
	"time"
)

type User struct {
	UUID       UUID       `json:"uuid" db:"uuid"`
	Name       string     `json:"name" db:"name"`
	Session    UUID       `json:"session" db:"session"`
	AvatarURL  string     `json:"avatar_url" db:"avatar_url"`
	GitRemote  StringList `json:"git_remote" db:"git_remote"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	ModifiedAt time.Time  `json:"modified_at" db:"modified_at"`

	Organizations []Organization
	Repositories  []Repository
}
