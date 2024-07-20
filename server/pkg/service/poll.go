package service

import (
	"context"

	"github.com/google/uuid"
)

type PollStatus int

const (
	INIT PollStatus = iota
	PENDING
	DONE
)

type Poll struct {
	ID  uuid.UUID
	URL string

	failedChan chan error
	status     PollStatus
	ctx        context.Context
}

func NewPoll(ctx context.Context, url string, fn func(context.Context, *Poll) error) *Poll {
	poll := &Poll{
		ID:     uuid.New(),
		URL:    url,
		status: INIT,
		ctx:    ctx,
	}

	go func(ctx context.Context) {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()

		poll.failedChan <- fn(ctx, poll)
		close(poll.failedChan)
		poll.Update(DONE)
	}(ctx)

	return poll
}

func (p *Poll) Update(status PollStatus) {
	p.status = status
}

func (p Poll) Status() PollStatus {
	return p.status
}

func (p *Poll) Failed() error {
	return <-p.failedChan
}
