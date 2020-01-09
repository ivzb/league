package summoner

import (
	"fmt"

	"league/util/http"
)

const bySummonerNameURL = "lol/summoner/v4/summoners/by-name/%s"

type (
	Repo interface {
		BySummonerName(name string) (*DTO, error)
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

func (repo *repo) BySummonerName(name string) (*DTO, error) {
	url := fmt.Sprintf(bySummonerNameURL, name)
	var dto *DTO

	_, err := repo.http.Get(url, &dto)

	return dto, err
}
