package summoner

import (
	"fmt"

	"league/util/http"
)

const byNameURL = "lol/summoner/v4/summoners/by-name/%s"

type (
	Summoner interface {
		ByName(name string) (*DTO, error)
	}

	summoner struct {
		http http.HTTP
	}
)

func New(http http.HTTP) Summoner {
	return &summoner{
		http: http,
	}
}

func (s *summoner) ByName(name string) (*DTO, error) {
	url := fmt.Sprintf(byNameURL, name)
	var dto *DTO

	_, err := s.http.Get(url, &dto)

	return dto, err
}
