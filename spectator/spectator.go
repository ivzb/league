package spectator

import (
	"fmt"

	"league/http"
)

const bySummonerURL = "lol/spectator/v4/active-games/by-summoner/%s"

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
	url := fmt.Sprintf(bySummonerURL, id)
	var dto *DTO

	err := s.http.Get(url, &dto)

	return dto, err
}
