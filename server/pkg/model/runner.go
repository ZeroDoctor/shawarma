package model

import (
	"time"
)

type Runner struct {
	UUID       UUID      `db:"uuid" json:"uuid"`
	Type       string    `db:"type" json:"type"`
	Hostname   string    `db:"hostname" json:"hostname"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	ModifiedAt time.Time `db:"modified_at" json:"moified_at"`

	Pipelines []Pipeline `json:"pipelines"`
}
