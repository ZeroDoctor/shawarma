package httputils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	GET     string = "GET"
	POST    string = "POST"
	PUT     string = "PUT"
	OPTIONS string = "OPTIONS"
	PATCH   string = "PATCH"
	DELETE  string = "DELETE"
)

var (
	ErrNextPageFuncUndefined error = errors.New("next page function undefined")
)

type RequestLimiter struct {
	totalRequest    int
	nextRequestTime func(*http.Response) time.Duration
	prevResponse    *http.Response
}

func NewRequestLimiter(next func(*http.Response) time.Duration) *RequestLimiter {
	rl := &RequestLimiter{
		totalRequest:    0,
		nextRequestTime: next,
	}

	if next == nil {
		rl.nextRequestTime = func(*http.Response) time.Duration {
			rl.totalRequest++
			return 0 * time.Second
		}
	}

	return rl
}

func (rl *RequestLimiter) NextRequestTime(resp *http.Response) time.Duration {
	if resp == nil {
		return 0
	}

	if rl.nextRequestTime == nil {
		rl.nextRequestTime = func(*http.Response) time.Duration {
			rl.totalRequest++
			return 0 * time.Second
		}
	}

	return rl.nextRequestTime(resp)
}

func (rl *RequestLimiter) NewRequest(method, u string, body io.Reader) *Request {
	uri, err := url.Parse(u)
	if err != nil {
		return &Request{
			Err: fmt.Errorf("failed to parse [url=%s] [error=%w]", u, err),
		}
	}

	return &Request{
		Uri:    uri,
		Method: method,
		Body:   body,

		Limiter: rl,
	}
}

type Request struct {
	Uri     *url.URL
	Method  string
	Body    io.Reader
	Headers map[string]string

	NextPage func(args ...interface{}) (string, bool, error)
	Err      error

	Limiter *RequestLimiter
}

func (r *Request) Do() (*http.Response, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	req, err := http.NewRequest(r.Method, r.Uri.String(), r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request [url=%s] [error=%s]",
			r.Uri.String(), err.Error(),
		)
	}
	addHeaders(r.Headers, req)

	time.Sleep(r.Limiter.NextRequestTime(r.Limiter.prevResponse))
	resp, err := http.DefaultClient.Do(req)
	r.Limiter.prevResponse = resp
	return resp, err
}

func (r *Request) DoAll() ([]*http.Response, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	if r.NextPage == nil {
		return nil, ErrNextPageFuncUndefined
	}

	var resps []*http.Response
	client := http.DefaultClient

	hasNext := true
	for hasNext {
		req, err := http.NewRequest(r.Method, r.Uri.String(), r.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to create new request [url=%s] [error=%s]",
				r.Uri.String(), err.Error(),
			)
		}
		addHeaders(r.Headers, req)

		fmt.Printf("[url=%s] [method=%s] [headers=%+v]\n", r.Uri.String(), r.Method, r.Headers)

		time.Sleep(r.Limiter.NextRequestTime(r.Limiter.prevResponse))
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		resps = append(resps, resp)
		r.Limiter.prevResponse = resp

		var uri string
		uri, hasNext, err = r.NextPage(resp)
		if err != nil {
			return resps, err
		}

		r.Uri, err = url.Parse(uri)
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
