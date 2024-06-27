package github

import (
	"github.com/zerodoctor/shawarma/pkg/plugin/github/service"
	"github.com/zerodoctor/shawarma/pkg/remote"
)

func init() {
	remote.Register("github", service.NewGithubService())
}
