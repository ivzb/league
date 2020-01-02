package league

import (
	"fmt"

	"league/util/http"
)

const bySummonerIdURL = "lol/league/v4/entries/by-summoner/%s"

type (
	League interface {
		BySummonerId(id string) ([]*LeagueDTO, error)
	}

	league struct {
		http http.HTTP
	}
)

func New(http http.HTTP) League {
	return &league{
		http: http,
	}
}

func (l *league) BySummonerId(id string) ([]*LeagueDTO, error) {
	url := fmt.Sprintf(bySummonerIdURL, id)
	var dto []*LeagueDTO

	_, err := l.http.Get(url, &dto)

	return dto, err
}
