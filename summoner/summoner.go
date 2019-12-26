package summoner

import (
	"encoding/json"
	"fmt"
	nhttp "net/http"

	"league/config"
	"league/http"
)

const byNamePath = "lol/summoner/v4/summoners/by-name"

type (
	Summoner interface {
		ByName(name string) (*DTO, error)
	}

	summoner struct {
		config *config.Config
		http http.HTTP
	}
)

func New(config *config.Config, http http.HTTP) Summoner {
	return &summoner{
		config:  config,
		http: http,
	}
}

func (s *summoner) ByName(name string) (*DTO, error) {
	url := fmt.Sprintf("%s/%s/%s", s.config.BaseURL, byNamePath, name)

	headers := map[string]string{
		"X-Riot-Token": s.config.ApiKey,
	}

	response, err := s.http.Do(nhttp.MethodGet, url, nil, headers)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var dto *DTO

	if err := json.NewDecoder(response.Body).Decode(&dto); err != nil {
		return nil, err
	}

	return dto, nil
}