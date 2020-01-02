package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	nhttp "net/http"
	"strconv"
	"strings"
	"time"

	"league/util/config"
)

const (
	fmtBaseURL = "%s/%s"
	xRiotToken = "X-Riot-Token"
	retryAfter = "Retry-After"
)

type (
	HTTP interface {
		Get(url string, response interface{}) ([]byte, error)
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

func (h *http) Get(url string, dto interface{}) ([]byte, error) {
	if !strings.Contains(url, "http") {
		url = fmt.Sprintf(fmtBaseURL, h.config.BaseURL, url)
	}

	request, err := nhttp.NewRequest(nhttp.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Add(xRiotToken, h.config.ApiKey)

	client := &nhttp.Client{}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		var error *Error
		err := json.Unmarshal(bytes, &error)

		if err != nil {
			return bytes, fmt.Errorf("couldn't parse error respnse: %v", err)
		}

		if error.Status != nil {
			if error.Status.StatusCode == 429 {
				retryAfterSeconds, err := strconv.Atoi(response.Header.Get(retryAfter))

				if err != nil {
					return bytes, fmt.Errorf("couldn't parse retry after header: %v", err)
				}

				fmt.Printf("hit rate limit, waiting for %d seconds\n", retryAfterSeconds)

				<-time.After(time.Duration(retryAfterSeconds) * time.Second)

				return h.Get(url, &dto)
			}

			return bytes, fmt.Errorf("URL: %s\nStatus code: %d\nBody: %s", url, response.StatusCode, string(bytes))
		}
	}

	return bytes, json.Unmarshal(bytes, &dto)
}
