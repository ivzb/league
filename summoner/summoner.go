package summoner

import (
	"encoding/json"
	"fmt"

	"league/http"
)

const byNamePath = "lol/summoner/v4/summoners/by-name/%s"

type (
	Summoner interface {
		ByName(name string) (*DTO, error)
	}

	summoner struct {
		http   http.HTTP
	}
)

func New(http http.HTTP) Summoner {
	return &summoner{
		http:   http,
	}
}

func (s *summoner) ByName(name string) (*DTO, error) {
	url := fmt.Sprintf(byNamePath, name)

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
