package model

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
)

type PipelineStatus string

func (ps *PipelineStatus) Scan(value interface{}) error {
	str, _ := value.(string)
	*ps = PipelineStatus(str)

	return nil
}

func (ps PipelineStatus) Value() (driver.Value, error) {
	return string(ps), nil
}

const (
	CREATED PipelineStatus = "created"
)

type Action string

const (
	CONTINUE Action = "continue"
	PAUSE    Action = "pause"
	STOP     Action = "stop"
)

type StatusEvent string

const (
	TIME   StatusEvent = "time"
	STATUS StatusEvent = "status"
)

type StatusEventName string

const (
	NONE    StatusEventName = "none"
	FAILURE StatusEventName = "failure"
	SUCCESS StatusEventName = "success"
)

type Pipeline struct {
	ID         int            `db:"id"`
	Type       string         `db:"type"`
	Status     PipelineStatus `db:"status"`
	CreatedAt  time.Time      `db:"created_at"`
	ModifiedAt time.Time      `db:"modified_at"`

	RepoID   uuid.UUID `db:"repo_id"`
	RunnerID uuid.UUID `db:"runner_id"`

	Steps  []Step
	Events []Event
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
	Webhook    string          `db:"webhook"`
	Type       StatusEvent     `db:"type"`
	StatusName StatusEventName `db:"status_name"`
	Action     Action          `db:"action"`
	Deadline   string          `db:"deadline"`
	After      string          `db:"after"`
	CreatedAt  time.Time       `db:"created_at"`
	ModifiedAt time.Time       `db:"modified_at"`

	PipelineID int       `db:"pipeline_id"`
	StepID     uuid.UUID `db:"step_id"`
}
