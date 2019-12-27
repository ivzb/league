package spectator

import (
	"encoding/json"
	"fmt"

	"league/http"
)

const bySummonerPath = "lol/spectator/v4/active-games/by-summoner/%s"

type (
	Spectator interface {
		BySummoner(id string) (*DTO, error)
	}

	spectator struct {
		http http.HTTP
	}
)

func New(http http.HTTP) Spectator {
	return &spectator{
		http: http,
	}
}

func (s *spectator) BySummoner(id string) (*DTO, error) {
	url := fmt.Sprintf(bySummonerPath, id)

	response, err := s.http.Get(url)

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
