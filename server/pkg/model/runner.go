package model

import (
	"github.com/google/uuid"
)

type Runner struct {
	UUID       uuid.UUID `db:"uuid"`
	Type       string    `db:"type"`
	Hostname   string    `db:"hostname"`
	CreatedAt  Time      `db:"created_at"`
	ModifiedAt Time      `db:"modified_at"`

	Pipelines []Pipeline
}
