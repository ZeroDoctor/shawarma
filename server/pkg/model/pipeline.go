package model

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"

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
	ID         int            `db:"id" json:"id,omitempty"`
	Type       string         `db:"type" json:"type,omitempty"`
	Status     PipelineStatus `db:"status" json:"status,omitempty"`
	CreatedAt  time.Time      `db:"created_at" json:"created_at,omitempty"`
	ModifiedAt time.Time      `db:"modified_at" json:"modified_at,omitempty"`

	RepoID   uuid.UUID `db:"repo_id" json:"repo_id,omitempty"`
	RunnerID uuid.UUID `db:"runner_id" json:"runner_id,omitempty"`

	Steps  []Step  `json:"steps,omitempty"`
	Events []Event `json:"events,omitempty"`
}

type Step struct {
	UUID       UUID      `db:"uuid" json:"uuid,omitempty"`
	Name       string    `db:"name" json:"name,omitempty"`
	Image      string    `db:"image" json:"image,omitempty"`
	Commands   Commands  `db:"commands" json:"commands,omitempty"`
	Privileged bool      `db:"privileged" json:"privileged,omitempty"`
	Detach     bool      `db:"detach" json:"detach,omitempty"`
	CreatedAt  time.Time `db:"created_at" json:"created_at,omitempty"`
	ModifiedAt time.Time `db:"modified_at" json:"modified_at,omitempty"`

	PipelineID int `db:"pipeline_id" json:"pipeline_id,omitempty"`

	Events []Event `json:"events,omitempty"`
}

type Environment struct {
	Key        string    `db:"key" json:"key,omitempty"`
	Data       string    `db:"data" json:"data,omitempty"`
	Protected  bool      `db:"protected" json:"protected,omitempty"`
	CreatedAt  time.Time `db:"created_at" json:"created_at,omitempty"`
	ModifiedAt time.Time `db:"modified_at" json:"modified_at,omitempty"`

	RepoID UUID `db:"repo_id" json:"repo_id,omitempty"`
	OrgID  UUID `db:"org_id" json:"org_id,omitempty"`
}

type Event struct {
	Webhook    string          `db:"webhook" json:"webhook,omitempty"`
	Type       StatusEvent     `db:"type" json:"type,omitempty"`
	StatusName StatusEventName `db:"status_name" json:"status_name,omitempty"`
	Action     Action          `db:"action" json:"action,omitempty"`
	Deadline   string          `db:"deadline" json:"deadline,omitempty"`
	After      string          `db:"after" json:"after,omitempty"`
	CreatedAt  time.Time       `db:"created_at" json:"created_at,omitempty"`
	ModifiedAt time.Time       `db:"modified_at" json:"modified_at,omitempty"`

	PipelineID int  `db:"pipeline_id" json:"pipeline_id,omitempty"`
	StepID     UUID `db:"step_id" json:"step_id,omitempty"`
}
