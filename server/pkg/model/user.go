package model

import "github.com/google/uuid"

type User struct {
	UUID       uuid.UUID `json:"uuid" db:"uuid"`
	Name       string    `json:"name" db:"name"`
	Session    string    `json:"session" db:"session"`
	CreatedAt  Time      `json:"created_at" db:"created_at"`
	ModifiedAt Time      `json:"modified_at" db:"modified_at"`
}
