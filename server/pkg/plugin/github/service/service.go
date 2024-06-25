package service

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/zerodoctor/shawarma/pkg/service"
)

const (
	GITHUB_API_VERSION string = "2022-11-28"
	GITHUB_API_ACCEPT  string = "application/vnd.github+json"
)

var (
	ErrPaginationTypeExists error = errors.New("pagination option already set")
	ErrPaginationArgsNotSet error = errors.New("arguments not set for github pagination")
)

type GithubRequest struct {
	*service.Request
}

func NewRequest(method, u string, body io.Reader) *GithubRequest {
	uri, err := url.Parse(u)
	if err != nil {
		return &GithubRequest{
			Request: &service.Request{

				Err: fmt.Errorf("failed to parse [url=%s] [error=%w]", u, err),
			},
		}
	}

	return &GithubRequest{
		Request: &service.Request{
			Uri:    uri,
			Method: method,
			Body:   body,
		},
	}
}

func (r *GithubRequest) OptionGithubHeaders(token string) *GithubRequest {
	if r.Err != nil {
		return r
	}

	if r.Headers == nil {
		r.Headers = make(map[string]string, 3)
	}

	r.Headers["Accept"] = GITHUB_API_ACCEPT
	if token != "" {
		r.Headers["Authorization"] = "Bearer " + token
	}
	r.Headers["X-GitHub-Api-Version"] = GITHUB_API_VERSION

	return r
}

func (r *GithubRequest) OptionGithubPages(total int) *GithubRequest {
	if r.Err != nil {
		return r
	}

	u := r.Uri.String() + "?per_page=" + strconv.Itoa(total)
	if strings.Contains(r.Uri.String(), "?") {
		u = r.Uri.String() + "&per_page=" + strconv.Itoa(total)
	}

	uri, err := url.Parse(u)
	if err != nil {
		r.Err = fmt.Errorf("failed to parse [url=%s] [error=%w]", u, err)
		return r
	}

	r.Uri = uri
	return r
}

func (r *GithubRequest) OptionGithubPagination() *GithubRequest {
	if r.Err != nil {
		return r
	}

	if r.Next != nil {
		r.Err = ErrPaginationTypeExists
		return r
	}

	r.Next = func(args ...interface{}) (string, bool, error) {
		if len(args) <= 0 {
			return "", false, ErrPaginationArgsNotSet
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
