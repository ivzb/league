package league

import (
	"encoding/json"
	"net/http"

	h "league/util/http"
)

type (
	Web interface {
		BySummonerId(w http.ResponseWriter, r *http.Request)
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

func (web *web) BySummonerId(w http.ResponseWriter, r *http.Request) {
	summonerIds, ok := r.URL.Query()["summoner_id"]

	if !ok {
		h.BadRequest(w, "invalid id parameter")
		return
	}

	dto, err := web.repo.BySummonerId(summonerIds[0])

	if err != nil {
		h.BadRequest(w, "couldn't load league")
		return
	}

	body, err := json.Marshal(dto)

	if err != nil {
		h.InternalServerError(w)
		return
	}

	h.OK(w, body)
}
