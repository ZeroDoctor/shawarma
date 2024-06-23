package service

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
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
	uri     *url.URL
	method  string
	body    io.Reader
	headers map[string]string

	next func(args ...interface{}) (string, bool, error)
	err  error
}

func NewRequest(method, u string, body io.Reader) *Request {
	uri, err := url.Parse(u)
	if err != nil {
		return &Request{
			err: fmt.Errorf("failed to parse [url=%s] [error=%w]", u, err),
		}
	}

	return &Request{
		uri:    uri,
		method: method,
		body:   body,
	}
}

func (r *Request) OptionGithubHeaders(token string) *Request {
	if r.err != nil {
		return r
	}

	if r.headers == nil {
		r.headers = make(map[string]string, 3)
	}

	r.headers["Accept"] = GITHUB_API_ACCEPT
	if token != "" {
		r.headers["Authorization"] = "Bearer " + token
	}
	r.headers["X-GitHub-Api-Version"] = GITHUB_API_VERSION

	return r
}

func (r *Request) OptionGithubPages(total int) *Request {
	if r.err != nil {
		return r
	}

	u := r.uri.String() + "?per_page=" + strconv.Itoa(total)
	if strings.Contains(r.uri.String(), "?") {
		u = r.uri.String() + "&per_page=" + strconv.Itoa(total)
	}

	uri, err := url.Parse(u)
	if err != nil {
		r.err = fmt.Errorf("failed to parse [url=%s] [error=%w]", u, err)
		return r
	}

	r.uri = uri
	return r
}

func (r *Request) OptionGithubPagination() *Request {
	if r.err != nil {
		return r
	}

	if r.next != nil {
		r.err = ErrPaginationTypeExists
		return r
	}

	r.next = func(args ...interface{}) (string, bool, error) {
		if len(args) <= 0 {
			log.Warn("arguments not set for github pagination")
		}

		resp, ok := args[0].(*http.Response)
		if !ok {
			return "", false, fmt.Errorf(
				"pagination failed to type cast args[0] to *http.Response when actual [type=%T]", args[0],
			)
		}

		link := resp.Header.Get("link")
		if link == "" {
			return "", false, nil
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
			return "", false, nil
		}

		split = strings.Split(next, ";")
		if len(split) <= 1 {
			return "", false, fmt.Errorf("pagination failed unexpected [result=%s]", split)
		}

		if len(split[0]) <= 2 {
			return "", false, fmt.Errorf("unexpected [url=%s] format", split[0])
		}

		return split[0][1 : len(split)-1], true, nil
	}

	return r
}

func (r *Request) Do() (*http.Response, error) {
	if r.err != nil {
		return nil, r.err
	}

	req, err := http.NewRequest(r.method, r.uri.String(), r.body)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request [url=%s] [error=%s]",
			r.uri.String(), err.Error(),
		)
	}
	addHeaders(r.headers, req)

	client := http.DefaultClient
	return client.Do(req)
}

func (r *Request) DoAll() ([]*http.Response, error) {
	if r.err != nil {
		return nil, r.err
	}

	var resps []*http.Response
	client := http.DefaultClient

	hasNext := true
	for hasNext {
		req, err := http.NewRequest(r.method, r.uri.String(), r.body)
		if err != nil {
			return nil, fmt.Errorf("failed to create new request [url=%s] [error=%s]",
				r.uri.String(), err.Error(),
			)
		}
		addHeaders(r.headers, req)

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		resps = append(resps, resp)

		var uri string
		uri, hasNext, err = r.next(resp)
		if err != nil {
			return resps, err
		}

		r.uri, err = url.Parse(uri)
		if err != nil {
			return nil, err
		}
	}

	return resps, nil
}

func addHeaders(headers map[string]string, req *http.Request) {
	for k, v := range headers {
		req.Header.Add(k, v)
	}
}
