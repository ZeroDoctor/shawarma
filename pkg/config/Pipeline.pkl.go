// Code generated from Pkl module `zerodoctor.shawarma.pkg.config`. DO NOT EDIT.
package config

type Pipeline struct {
	Type string `pkl:"type"`

	Environment map[string]string `pkl:"environment"`

	Steps []*Step `pkl:"steps"`

	OnBranch string `pkl:"on_branch"`

	OnFailure []*ResultEvent `pkl:"on_failure"`

	OnSuccess []*ResultEvent `pkl:"on_success"`

	OnTime []*TimeEvent `pkl:"on_time"`
}
