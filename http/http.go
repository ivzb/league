package http

import (
	"encoding/json"
	"fmt"
	nhttp "net/http"
	"strings"

	"league/config"
)

const (
	fmtBaseURL = "%s/%s"
	xRiotToken = "X-Riot-Token"
)

type (
	HTTP interface {
		Get(url string, response interface{}) error
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

func (h *http) Get(url string, dto interface{}) error {
	if !strings.Contains(url, "http") {
		url = fmt.Sprintf(fmtBaseURL, h.config.BaseURL, url)
	}

	request, err := nhttp.NewRequest(nhttp.MethodGet, url, nil)

	if err != nil {
		return err
	}

	request.Header.Add(xRiotToken, h.config.ApiKey)

	client := &nhttp.Client{}

	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(&dto)
}
