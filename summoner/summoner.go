package summoner

import (
	h "league/util/http"
)

type (
	Summoner interface {
		Repo() Repo
		Web() Web
	}

	summoner struct {
		repo Repo
		web  Web
	}
)

func New(http h.HTTP) Summoner {
	repo := newRepo(http)
	web := newWeb(repo)

	return &summoner{
		repo: repo,
		web:  web,
	}
}

func (s *summoner) Repo() Repo {
	return s.repo
}

func (s *summoner) Web() Web {
	return s.web
}
