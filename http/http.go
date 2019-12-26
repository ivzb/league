package http

import (
	"errors"
	"io"
	nhttp "net/http"
)

var errNewRequestFailed = errors.New("http: new request failed")

type (
	HTTP interface {
		Do(method, url string, body io.Reader, headers map[string]string) (r *nhttp.Response, err error)
	}

	http struct {
	}
)

func New() HTTP {
	return &http{}
}

func (h *http) Do(method, url string, body io.Reader, headers map[string]string) (r *nhttp.Response, err error) {
	request, err := nhttp.NewRequest(method, url, body)

	if err != nil {
		return nil, errNewRequestFailed
	}

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	client := &nhttp.Client{}

	return client.Do(request)
}