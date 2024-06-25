package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/zerodoctor/shawarma/pkg/config"
	"github.com/zerodoctor/shawarma/pkg/model"
)

func NewPipeline(repoID string, runnerID string, pipe config.Pipeline) model.Pipeline {
	now := time.Now()

	steps := make([]model.Step, len(pipe.Steps))
	for i := range pipe.Steps {
		steps[i] = NewStep(0, *pipe.Steps[i])
	}

	eventsLen := len(pipe.OnFailure) + len(pipe.OnSuccess) + len(pipe.OnTime)
	events := make([]model.Event, eventsLen)
	for i := range pipe.OnFailure {
		events[i] = NewStatusEvent(0, "", *pipe.OnFailure[i], model.FAILURE)
	}

	for i := range pipe.OnSuccess {
		events[i+len(pipe.OnFailure)] = NewStatusEvent(0, "", *pipe.OnSuccess[i], model.SUCCESS)
	}

	for i := range pipe.OnTime {
		events[i+len(pipe.OnFailure)+len(pipe.OnSuccess)] = NewTimeEvent(0, "", *pipe.OnTime[i])
	}

	return model.Pipeline{
		Type:       pipe.Type,
		Status:     model.CREATED,
		CreatedAt:  model.Time(now),
		ModifiedAt: model.Time(now),
		RepoID:     uuid.MustParse(repoID),
		RunnerID:   uuid.MustParse(runnerID),
		Steps:      steps,
		Events:     events,
	}
}

func NewStep(pipelineID int, step config.Step) model.Step {
	now := time.Now()

	eventsLen := len(step.OnFailure) + len(step.OnSuccess) + len(step.OnTime)
	events := make([]model.Event, eventsLen)
	for i := range step.OnFailure {
		events[i] = NewStatusEvent(0, "", *step.OnFailure[i], model.FAILURE)
	}

	for i := range step.OnSuccess {
		events[i+len(step.OnFailure)] = NewStatusEvent(0, "", *step.OnSuccess[i], model.SUCCESS)
	}

	for i := range step.OnTime {
		events[i+len(step.OnFailure)+len(step.OnSuccess)] = NewTimeEvent(0, "", *step.OnTime[i])
	}

	return model.Step{
		UUID:       uuid.New(),
		Name:       step.Name,
		Image:      step.Image,
		Commands:   step.Commands,
		Privileged: step.Privileged,
		Detach:     step.Detach,
		CreatedAt:  model.Time(now),
		ModifiedAt: model.Time(now),

		PipelineID: pipelineID,
	}
}

func NewStatusEvent(pipelineID int, stepID string, event config.StatusEvent, status model.StatusEventName) model.Event {
	return newEvent(model.STATUS, event, status)
}

func NewTimeEvent(pipelineID int, stepID string, event config.TimeEvent) model.Event {
	return newEvent(model.TIME, event, model.NONE)
}

func newEvent[T config.StatusEvent | config.TimeEvent](eType model.StatusEvent, event T, status model.StatusEventName) model.Event {
	now := time.Now()
	e := model.Event{
		Type:       eType,
		CreatedAt:  model.Time(now),
		ModifiedAt: model.Time(now),
		StatusName: status,
	}

	switch ev := any(event).(type) {
	case config.StatusEvent:
		e.Webhook = ev.Webhook
		e.Action = model.Action(ev.Action)
	case config.TimeEvent:
		e.Webhook = ev.Webhook
		e.Action = model.Action(ev.Action)
		e.Deadline = ev.Deadline.Unit.String()
		e.After = ev.After
	}

	return e
}
