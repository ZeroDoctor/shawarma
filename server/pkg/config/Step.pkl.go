// Code generated from Pkl module `zerodoctor.shawarma.pkg.config`. DO NOT EDIT.
package config

type Step struct {
	Name string `pkl:"name"`

	Image string `pkl:"image"`

	Commands []string `pkl:"commands"`

	Environment map[string]string `pkl:"environment"`

	Privileged bool `pkl:"privileged"`

	Detach bool `pkl:"detach"`

	OnBranch string `pkl:"on_branch"`

	OnFailure []*StatusEvent `pkl:"on_failure"`

	OnSuccess []*StatusEvent `pkl:"on_success"`

	OnTime []*TimeEvent `pkl:"on_time"`
}
