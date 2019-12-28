package match

import (
	"fmt"

	"league/http"
)

const byAccountURL = "lol/match/v4/matchlists/by-account/%s"

type (
	Match interface {
		ByID(id string)
		ByAccount(accountID string) (*MatchlistDto, error)
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

func (m *match) ByAccount(id string) (*MatchlistDto, error) {
	url := fmt.Sprintf(byAccountURL, id)
	var dto *MatchlistDto

	err := m.http.Get(url, &dto)

	return dto, err
}
