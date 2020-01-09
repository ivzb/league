package league

import (
	h "league/util/http"
)

type (
	League interface {
		Repo() Repo
		Web() Web
	}

	league struct {
		repo Repo
		web  Web
	}
)

func New(http h.HTTP) League {
	repo := newRepo(http)
	web := newWeb(repo)

	return &league{
		repo: repo,
		web:  web,
	}
}

func (l *league) Repo() Repo {
	return l.repo
}

func (l *league) Web() Web {
	return l.web
}
