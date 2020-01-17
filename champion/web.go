package champion

import (
	"encoding/json"
	"net/http"

	h "league/util/http"
)

type (
	Web interface {
		All(w http.ResponseWriter, r *http.Request)
		Sprite(w http.ResponseWriter, r *http.Request)
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

func (web *web) All(w http.ResponseWriter, r *http.Request) {
	dto, err := web.repo.All()

	if err != nil {
		h.BadRequest(w, "couldn't load champions")
		return
	}

	body, err := json.Marshal(dto)

	if err != nil {
		h.InternalServerError(w)
		return
	}

	h.OK(w, body)
}

func (web *web) Sprite(w http.ResponseWriter, r *http.Request) {
	ids, ok := r.URL.Query()["id"]

	if !ok {
		h.BadRequest(w, "invalid id parameter")
		return
	}

	sprite, err := web.repo.Sprite(ids[0])

	if err != nil {
		h.BadRequest(w, "couldn't load sprite")
		return
	}

	h.OK(w, sprite)
}
