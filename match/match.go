package match

import (
	"fmt"

	"league/http"
)

const byAccountURL = "lol/match/v4/matchlists/by-account/%s"

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
	url := fmt.Sprintf(byAccountURL, id)
	var dto *DTO

	err := m.http.Get(url, &dto)

	return dto, err
}
