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
	ID         int            `db:"id" json:"id"`
	Type       string         `db:"type" json:"type"`
	Status     PipelineStatus `db:"status" json:"status"`
	CreatedAt  time.Time      `db:"created_at" json:"created_at"`
	ModifiedAt time.Time      `db:"modified_at" json:"modified_at"`

	RepoID   uuid.UUID `db:"repo_id" json:"repo_id"`
	RunnerID uuid.UUID `db:"runner_id" json:"runner_id"`

	Steps  []Step  `json:"steps"`
	Events []Event `json:"events"`
}

type Step struct {
	UUID       UUID      `db:"uuid" json:"uuid"`
	Name       string    `db:"name" json:"name"`
	Image      string    `db:"image" json:"image"`
	Commands   Commands  `db:"commands" json:"commands"`
	Privileged bool      `db:"privileged" json:"privileged"`
	Detach     bool      `db:"detach" json:"detach"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	ModifiedAt time.Time `db:"modified_at" json:"modified_at"`

	PipelineID int `db:"pipeline_id" json:"pipeline_id"`

	Events []Event `json:"events"`
}

type Environment struct {
	Key        string    `db:"key" json:"key"`
	Data       string    `db:"data" json:"data"`
	Protected  bool      `db:"protected" json:"protected"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	ModifiedAt time.Time `db:"modified_at" json:"modified_at"`

	RepoID UUID `db:"repo_id" json:"repo_id"`
	OrgID  UUID `db:"org_id" json:"org_id"`
}

type Event struct {
	Webhook    string          `db:"webhook" json:"webhook"`
	Type       StatusEvent     `db:"type" json:"type"`
	StatusName StatusEventName `db:"status_name" json:"status_name"`
	Action     Action          `db:"action" json:"action"`
	Deadline   string          `db:"deadline" json:"deadline"`
	After      string          `db:"after" json:"after"`
	CreatedAt  time.Time       `db:"created_at" json:"created_at"`
	ModifiedAt time.Time       `db:"modified_at" json:"modified_at"`

	PipelineID int  `db:"pipeline_id" json:"pipeline_id"`
	StepID     UUID `db:"step_id" json:"step_id"`
}
