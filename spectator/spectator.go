package spectator

import (
	"encoding/json"
	"fmt"
	nhttp "net/http"

	"league/config"
	"league/http"
)

const bySummonerPath = "lol/spectator/v4/active-games/by-summoner"

type (
	Spectator interface {
		BySummoner(id string) (*DTO, error)
	}

	spectator struct {
		config *config.Config
		http   http.HTTP
	}
)

func New(config *config.Config, http http.HTTP) Spectator {
	return &spectator{
		config: config,
		http:   http,
	}
}

func (s *spectator) BySummoner(id string) (*DTO, error) {
	url := fmt.Sprintf("%s/%s/%s", s.config.BaseURL, bySummonerPath, id)

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
