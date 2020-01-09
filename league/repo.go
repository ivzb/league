package league

import (
	"fmt"

	"league/util/http"
)

const bySummonerIdURL = "lol/league/v4/entries/by-summoner/%s"

type (
	Repo interface {
		BySummonerId(id string) ([]*LeagueDTO, error)
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

func (r *repo) BySummonerId(id string) ([]*LeagueDTO, error) {
	url := fmt.Sprintf(bySummonerIdURL, id)
	var dto []*LeagueDTO

	_, err := r.http.Get(url, &dto)

	return dto, err
}
