// Code generated from Pkl module `zerodoctor.shawarma.pkg.config`. DO NOT EDIT.
package config

import (
	"github.com/apple/pkl-go/pkl"
	"github.com/zerodoctor/shawarma/pkg/config/action"
)

type TimeEvent struct {
	Deadline *pkl.Duration `pkl:"deadline"`

	After string `pkl:"after"`

	Action action.Action `pkl:"action"`

	Webhook string `pkl:"webhook"`
}
