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
	Repo interface {
		ByMatchID(matchID string) (*MatchDto, error)
		ByAccountID(accountID string, limit int) (*MatchlistDto, error)
	}

	repo struct {
		http http.HTTP
	}
)

func newRepo(http http.HTTP) Repo {
	return &repo{
		http: http,
	}
}

func (r *repo) ByMatchID(matchID string) (*MatchDto, error) {
	url := fmt.Sprintf(byMatchID, matchID)
	var dto *MatchDto

	_, err := r.http.Get(url, &dto)

	return dto, err
}

func (r *repo) ByAccountID(id string, limit int) (*MatchlistDto, error) {
	url := fmt.Sprintf(byAccountURL, id, limit)
	var dto *MatchlistDto

	_, err := r.http.Get(url, &dto)

	return dto, err
}
