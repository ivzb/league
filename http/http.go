package http

import (
	"fmt"
	nhttp "net/http"

	"league/config"
)

var (
	fmtBaseURL = "%s/%s"
	xRiotToken = "X-Riot-Token"
)

type (
	HTTP interface {
		Get(url string) (r *nhttp.Response, err error)
	}

	http struct {
		config *config.Config
	}
)

func New(config *config.Config) HTTP {
	return &http{
		config: config,
	}
}

func (h *http) Get(url string) (r *nhttp.Response, err error) {
	url = fmt.Sprintf(fmtBaseURL, h.config.BaseURL, url)

	request, err := nhttp.NewRequest(nhttp.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Add(xRiotToken, h.config.ApiKey)

	client := &nhttp.Client{}

	return client.Do(request)
}
