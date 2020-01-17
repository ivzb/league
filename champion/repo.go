package champion

import (
	"fmt"

	"league/util/http"
)

// get last version from https://ddragon.leagueoflegends.com/api/versions.json
const championsURL = "http://ddragon.leagueoflegends.com/cdn/10.1.1/data/en_US/champion.json"
const spriteURL = "http://ddragon.leagueoflegends.com/cdn/10.1.1/img/sprite/%s"


type (
	Repo interface {
		All() (*DTO, error)
		Sprite(id string) ([]byte, error)
	}

	repo struct {
		http http.HTTP

		dtoCache *DTO
		spriteCache map[string][]byte
	}
)

func newRepo(http http.HTTP) Repo {
	return &repo{
		http: http,

		spriteCache: map[string][]byte{},
	}
}

func (r *repo) All() (*DTO, error) {
	if r.dtoCache != nil {
		return r.dtoCache, nil
	}

	var dto *DTO
	_, err := r.http.Get(championsURL, &dto)

	if err == nil {
		r.dtoCache = dto
	}

	return dto, err
}

func (r *repo) Sprite(id string) ([]byte, error) {
	if sprite, ok := r.spriteCache[id]; ok {
		return sprite, nil
	}

	sprite, err := r.http.Get(fmt.Sprintf(spriteURL, id), nil)

	if err == nil {
		r.spriteCache[id] = sprite
	}

	return sprite, err
}
