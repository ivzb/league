package match

import (
	"league/util/config"
	h "league/util/http"
)

type (
	Match interface {
		Repo() Repo
		Web() Web
	}

	match struct {
		repo Repo
		web  Web
	}
)

func New(http h.HTTP, config *config.Config) Match {
	repo := newRepo(http)
	web := newWeb(repo, config)

	return &match{
		repo: repo,
		web:  web,
	}
}

func (m *match) Repo() Repo {
	return m.repo
}

func (m *match) Web() Web {
	return m.web
}
