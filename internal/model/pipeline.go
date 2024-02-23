package model

import (
	"time"

	"github.com/google/uuid"
)

type Action string

const (
	Continue Action = "continue"
	Pause    Action = "pause"
	Stop     Action = "stop"
)

type StatusEvent string

const (
	Time   StatusEvent = "time"
	Status StatusEvent = "status"
)

type Pipeline struct {
	ID         int       `db:"id"`
	Type       string    `db:"type"`
	Status     string    `db:"status"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`

	RepoID   uuid.UUID `db:"repo_id"`
	RunnerID uuid.UUID `db:"runner_id"`

	Steps  []Step
	Events []Event
	Repo   Repository
}

type Step struct {
	UUID       uuid.UUID `db:"uuid"`
	Name       string    `db:"name"`
	Image      string    `db:"image"`
	Commands   []string  `db:"commands"` // TODO: create type alias []string for db
	Privileged bool      `db:"privileged"`
	Detach     bool      `db:"detach"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`

	PipelineID int `db:"pipeline_id"`

	Events []Event
}

type Environment struct {
	Key        string    `db:"key"`
	Data       string    `db:"data"`
	Protected  bool      `db:"protected"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`

	RepoID uuid.UUID `db:"repo_id"`
	OrgID  uuid.UUID `db:"org_id"`
}

type Event struct {
	Webhook    string      `db:"webhook"`
	Type       StatusEvent `db:"type"`
	Action     Action      `db:"action"`
	Deadline   string      `db:"deadline"`
	CreatedAt  time.Time   `db:"created_at"`
	ModifiedAt time.Time   `db:"modified_at"`

	PipelineID int       `db:"pipeline_id"`
	StepID     uuid.UUID `db:"step_id"`
}
