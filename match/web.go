package match

import (
	"encoding/json"
	"net/http"

	"league/util/config"
	h "league/util/http"
)

type (
	Web interface {
		ByMatchID(w http.ResponseWriter, r *http.Request)
		ByAccountID(w http.ResponseWriter, r *http.Request)
	}

	web struct {
		repo   Repo
		config *config.Config
	}
)

func newWeb(repo Repo, config *config.Config) Web {
	return &web{
		repo:   repo,
		config: config,
	}
}

func (web *web) ByMatchID(w http.ResponseWriter, r *http.Request) {
	matchIds, ok := r.URL.Query()["match_id"]

	if !ok {
		h.BadRequest(w, "invalid id parameter")
		return
	}

	dto, err := web.repo.ByMatchID(matchIds[0])

	if err != nil {
		h.BadRequest(w, "couldn't load matches")
		return
	}

	body, err := json.Marshal(dto)

	if err != nil {
		h.InternalServerError(w)
		return
	}

	h.OK(w, body)
}

func (web *web) ByAccountID(w http.ResponseWriter, r *http.Request) {
	accountIds, ok := r.URL.Query()["account_id"]

	if !ok {
		h.BadRequest(w, "invalid id parameter")
		return
	}

	dto, err := web.repo.ByAccountID(accountIds[0], web.config.MatchesLimit)

	if err != nil {
		h.BadRequest(w, "couldn't load matches")
		return
	}

	body, err := json.Marshal(dto)

	if err != nil {
		h.InternalServerError(w)
		return
	}

	h.OK(w, body)
}
