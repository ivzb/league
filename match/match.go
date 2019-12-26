package match

import (
	"encoding/json"
	"fmt"

	"league/http"
)

const byAccountPath = "lol/match/v4/matchlists/by-account/%s"

type (
	Match interface {
		ByAccount(id string) (*DTO, error)
	}

	match struct {
		http http.HTTP
	}
)

func New(http http.HTTP) Match {
	return &match{
		http: http,
	}
}

func (m *match) ByAccount(id string) (*DTO, error) {
	url := fmt.Sprintf(byAccountPath, id)

	response, err := m.http.Get(url)

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
