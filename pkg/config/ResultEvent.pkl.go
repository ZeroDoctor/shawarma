// Code generated from Pkl module `zerodoctor.shawarma.pkg.config`. DO NOT EDIT.
package config

import "github.com/zerodoctor/shawarma/pkg/config/action"

type ResultEvent struct {
	Action action.Action `pkl:"action"`

	Webhook string `pkl:"webhook"`
}
