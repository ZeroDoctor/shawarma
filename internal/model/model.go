package model

import "time"

type Action string

const (
	Continue Action = "continue"
	Pause    Action = "pause"
	Stop     Action = "stop"
)

type Organization struct {
	ID         int       `db:"id"`
	Owner      string    `db:"owner"`
	Name       string    `db:"name"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`

	Environments []Environment
	Env          map[string]string
}

type Repository struct {
	ID         int       `db:"id"`
	Owner      string    `db:"owner"`
	Name       string    `db:"name"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`

	Branches     []Branch
	Environments []Environment
	Env          map[string]string
}

type Branch struct {
	ID           int       `db:"id"`
	Name         string    `db:"name"`
	LatestCommit string    `db:"latest_commit"`
	CreatedAt    time.Time `db:"created_at"`
	ModifiedAt   time.Time `db:"modified_at"`

	Commits []Commit
}

type Commit struct {
	Hash      string    `db:"commit"`
	Author    string    `db:"author"`
	CreatedAt time.Time `db:"created_at"`
}

type Runner struct {
	Type       string    `db:"type"`
	Hostname   string    `db:"hostname"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`

	RunningPipelineIDs []int
}

type Pipeline struct {
	ID         int       `db:"id"`
	Type       string    `db:"type"`
	Status     string    `db:"status"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`

	Steps  []Step
	Events []Event
	Repo   Repository
}

type Step struct {
	ID         int       `db:"id"`
	Name       string    `db:"name"`
	Image      string    `db:"image"`
	Commands   []string  `db:"commands"`
	Privileged bool      `db:"privileged"`
	Detach     bool      `db:"detach"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`

	Events []Event
}

type Environment struct {
	Key        string    `db:"key"`
	Data       string    `db:"data"`
	Protected  bool      `db:"protected"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`
}

type Event struct {
	Webhook    string    `db:"webhook"`
	Type       string    `db:"type"`
	Action     Action    `db:"action"`
	Deadline   string    `db:"deadline"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`
}
