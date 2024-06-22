package service

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

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

var (
	ErrPaginationTypeExists error = errors.New("pagination option already set")
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
	req    *http.Request
	method string
	body   io.Reader
	next   func(args ...interface{}) (string, error)
	err    error
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

	return &Request{
		req:    req,
		method: method,
		body:   body,
	}
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

func (r *Request) OptionGithubPagination() *Request {
	if r.next != nil {
		r.err = ErrPaginationTypeExists
		return r
	}

	r.next = func(args ...interface{}) (string, error) {
		if len(args) <= 0 {
			log.Warn("arguments not set for github pagination")
		}

		resp, ok := args[0].(*http.Response)
		if !ok {
			return "", fmt.Errorf(
				"pagination failed to type cast args[0] to *http.Response when actual [type=%T]", args[0],
			)
		}

		link := resp.Header.Get("link")
		if link == "" {
			return "", nil
		}

		split := strings.Split(link, ",")

		var next string
		for i := range split {
			if strings.Contains(split[i], "rel=\"next\"") {
				next = split[i]
				break
			}
		}

		if next == "" {
			return "", nil
		}

		split = strings.Split(next, ";")
		if len(split) <= 1 {
			return "", fmt.Errorf("pagination failed unexpected [result=%s]", split)
		}

		if len(split[0]) <= 2 {
			return "", fmt.Errorf("unexpected [url=%s] format", split[0])
		}

		return split[0][1 : len(split)-1], nil
	}

	return r
}

func (r *Request) Do() (*http.Response, error) {
	if r.err != nil {
		return nil, r.err
	}

	client := http.DefaultClient
	return client.Do(r.req)
}

func (r *Request) DoAll() ([]*http.Response, error) {
	if r.err != nil {
		return nil, r.err
	}

	var resps []*http.Response
	client := http.DefaultClient

	hasNext := true
	for hasNext {
		resp, err := client.Do(r.req)
		if err != nil {
			return nil, err
		}
		resps = append(resps, resp)

		url, err := r.next(resp)
		if err != nil {
			return resps, err
		} else if url == "" {
			break
		}

		r.req, err = http.NewRequest(r.method, url, r.body)
		if err != nil {
			return resps, fmt.Errorf("failed to create request [error=%w]", err)
		}
	}

	return resps, nil
}
