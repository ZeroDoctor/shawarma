package model

import (
	"time"
)

type Runner struct {
	UUID       UUID      `db:"uuid" json:"uuid,omitempty"`
	Type       string    `db:"type" json:"type,omitempty"`
	Hostname   string    `db:"hostname" json:"hostname,omitempty"`
	CreatedAt  time.Time `db:"created_at" json:"created_at,omitempty"`
	ModifiedAt time.Time `db:"modified_at" json:"moified_at,omitempty"`

	Pipelines []Pipeline `json:"pipelines,omitempty"`
}
