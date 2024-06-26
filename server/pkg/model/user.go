package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID       uuid.UUID `json:"uuid" db:"uuid"`
	Name       string    `json:"name" db:"name"`
	GitRemote  string    `json:"git_remote" db:"git_remote"`
	Session    string    `json:"session" db:"session"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	ModifiedAt time.Time `json:"modified_at" db:"modified_at"`
}
