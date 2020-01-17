package champion

import (
	h "league/util/http"
)

type (
	Champion interface {
		Repo() Repo
		Web() Web
	}

	champion struct {
		repo Repo
		web  Web
	}
)

func New(http h.HTTP) Champion {
	repo := newRepo(http)
	web := newWeb(repo)

	return &champion{
		repo: repo,
		web:  web,
	}
}

func (c *champion) Repo() Repo {
	return c.repo
}

func (c *champion) Web() Web {
	return c.web
}
