package summoner

import (
	"encoding/json"
	"net/http"

	h "league/util/http"
)

type (
	Web interface {
		BySummonerName(w http.ResponseWriter, r *http.Request)
	}

	web struct {
		repo Repo
	}
)

func newWeb(repo Repo) Web {
	return &web{
		repo: repo,
	}
}

func (web *web) BySummonerName(w http.ResponseWriter, r *http.Request) {
	names, ok := r.URL.Query()["name"]

	if !ok {
		h.BadRequest(w, "invalid name parameter")
		return
	}

	dto, err := web.repo.BySummonerName(names[0])

	if err != nil {
		h.BadRequest(w, "couldn't load summoner")
		return
	}

	body, err := json.Marshal(dto)

	if err != nil {
		h.InternalServerError(w)
		return
	}

	h.OK(w, body)
}
