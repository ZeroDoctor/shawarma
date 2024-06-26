package model

import (
	"time"

	"github.com/google/uuid"
)

type Runner struct {
	UUID       uuid.UUID `db:"uuid"`
	Type       string    `db:"type"`
	Hostname   string    `db:"hostname"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`

	Pipelines []Pipeline
}
