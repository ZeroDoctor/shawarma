package service

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
	"github.com/zerodoctor/shawarma/internal/db"
	"github.com/zerodoctor/shawarma/internal/logger"
)

var log *logrus.Logger = logger.Log

const (
	GITHUB_API_VERSION string = "2022-11-28"
	GITHUB_API_ACCEPT  string = "application/vnd.github+json"
)

const (
	HTTP_GET     string = "GET"
	HTTP_POST    string = "POST"
	HTTP_PUT     string = "PUT"
	HTTP_OPTIONS string = "OPTIONS"
	HTTP_PATCH   string = "PATCH"
	HTTP_DELETE  string = "DELETE"
)

type Service struct {
	db db.DB
}

func NewService(db db.DB) *Service {
	return &Service{
		db: db,
	}
}

type Request struct {
	req *http.Request
	err error
}

func NewRequest(method, u string, body io.Reader) *Request {
	uri, err := url.Parse(u)
	if err != nil {
		return &Request{
			err: fmt.Errorf("failed to parse [url=%s] [error=%w]", u, err),
		}
	}

	req, err := http.NewRequest(method, uri.String(), body)
	if err != nil {
		return &Request{err: fmt.Errorf("failed to create request [error=%w]", err)}
	}

	return &Request{req: req}
}

func (r *Request) OptionGithubHeaders(token string) *Request {
	if r.err != nil {
		return r

	}

	r.req.Header.Add("Accept", GITHUB_API_ACCEPT)
	if token != "" {
		r.req.Header.Add("Authorization", "Bearer "+token)
	}
	r.req.Header.Add("X-GitHub-Api-Version", GITHUB_API_VERSION)

	return r
}

func (r *Request) Do() (*http.Response, error) {
	if r.err != nil {
		return nil, r.err
	}

	client := http.DefaultClient
	return client.Do(r.req)
}
