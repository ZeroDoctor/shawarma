package model

import (
	"database/sql/driver"
	"errors"
	"strings"

	"github.com/google/uuid"
)

type PipelineStatus string

const (
	CREATED PipelineStatus = "created"
)

func (ps *PipelineStatus) Scan(value interface{}) error {
	str, _ := value.(string)
	*ps = PipelineStatus(str)

	return nil
}

func (ps PipelineStatus) Value() (driver.Value, error) {
	return string(ps), nil
}

type Action string

const (
	CONTINUE Action = "continue"
	PAUSE    Action = "pause"
	STOP     Action = "stop"
)

func (a *Action) Scan(value interface{}) error {
	str, _ := value.(string)
	*a = Action(str)

	return nil
}

func (a Action) Value() (driver.Value, error) {
	return string(a), nil
}

type StatusEvent string

const (
	TIME   StatusEvent = "time"
	STATUS StatusEvent = "status"
)

func (se *StatusEvent) Scan(value interface{}) error {
	str, _ := value.(string)
	*se = StatusEvent(str)

	return nil
}

func (se StatusEvent) Value() (driver.Value, error) {
	return string(se), nil
}

type StatusEventName string

const (
	NONE    StatusEventName = "none"
	FAILURE StatusEventName = "failure"
	SUCCESS StatusEventName = "success"
)

func (sen *StatusEventName) Scan(value interface{}) error {
	str, _ := value.(string)
	*sen = StatusEventName(str)

	return nil
}

func (sen StatusEventName) Value() (driver.Value, error) {
	return string(sen), nil
}

type Commands []string

func (c *Commands) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("commands is not a string")
	}
	*c = strings.Split(str, ",")

	return nil
}

func (c Commands) Value() (driver.Value, error) {
	return strings.Join(c, ","), nil
}

type Pipeline struct {
	ID         int            `db:"id"`
	Type       string         `db:"type"`
	Status     PipelineStatus `db:"status"`
	CreatedAt  Time           `db:"created_at"`
	ModifiedAt Time           `db:"modified_at"`

	RepoID   uuid.UUID `db:"repo_id"`
	RunnerID uuid.UUID `db:"runner_id"`

	Steps  []Step
	Events []Event
}

type Step struct {
	UUID       uuid.UUID `db:"uuid"`
	Name       string    `db:"name"`
	Image      string    `db:"image"`
	Commands   Commands  `db:"commands"`
	Privileged bool      `db:"privileged"`
	Detach     bool      `db:"detach"`
	CreatedAt  Time      `db:"created_at"`
	ModifiedAt Time      `db:"modified_at"`

	PipelineID int `db:"pipeline_id"`

	Events []Event
}

type Environment struct {
	Key        string `db:"key"`
	Data       string `db:"data"`
	Protected  bool   `db:"protected"`
	CreatedAt  Time   `db:"created_at"`
	ModifiedAt Time   `db:"modified_at"`

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
	CreatedAt  Time            `db:"created_at"`
	ModifiedAt Time            `db:"modified_at"`

	PipelineID int       `db:"pipeline_id"`
	StepID     uuid.UUID `db:"step_id"`
}
