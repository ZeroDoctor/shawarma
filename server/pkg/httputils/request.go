package httputils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	GET     string = "GET"
	POST    string = "POST"
	PUT     string = "PUT"
	OPTIONS string = "OPTIONS"
	PATCH   string = "PATCH"
	DELETE  string = "DELETE"
)

type Request struct {
	Uri     *url.URL
	Method  string
	Body    io.Reader
	Headers map[string]string

	Next func(args ...interface{}) (string, bool, error)
	Err  error
}

func NewRequest(method, u string, body io.Reader) *Request {
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
	}
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

	client := http.DefaultClient
	return client.Do(req)
}

func (r *Request) DoAll() ([]*http.Response, error) {
	if r.Err != nil {
		return nil, r.Err
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

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		resps = append(resps, resp)

		var uri string
		uri, hasNext, err = r.Next(resp)
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
