package match

import (
	"fmt"

	"league/util/http"
)

const (
	byMatchID    = "lol/match/v4/matches/%s"
	byAccountURL = "lol/match/v4/matchlists/by-account/%s?endIndex=%d"
)

type (
	Match interface {
		ByMatchID(id string) (*MatchDto, error)
		ByAccountID(accountID string, limit int) (*MatchlistDto, error)
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

func (m *match) ByMatchID(id string) (*MatchDto, error) {
	url := fmt.Sprintf(byMatchID, id)
	var dto *MatchDto

	_, err := m.http.Get(url, &dto)

	return dto, err
}

func (m *match) ByAccountID(id string, limit int) (*MatchlistDto, error) {
	url := fmt.Sprintf(byAccountURL, id, limit)
	var dto *MatchlistDto

	_, err := m.http.Get(url, &dto)

	return dto, err
}
