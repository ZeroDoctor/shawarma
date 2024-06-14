package service

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/zerodoctor/shawarma/internal/logger"
)

var log *logrus.Logger = logger.Log

const (
	GITHUB_API_VERSION string = "2022-11-28"
	GITHUB_API_ACCEPT  string = "application/vnd.github+json"
)

func addGithubHeaders(req *http.Request, token string) {
	req.Header.Add("Accept", GITHUB_API_ACCEPT)
	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}
	req.Header.Add("X-GitHub-Api-Version", GITHUB_API_VERSION)
}
