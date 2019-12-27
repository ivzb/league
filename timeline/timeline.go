package timeline

import (
	"fmt"

	"league/http"
)

const byMatchURL = "lol/match/v4/timelines/by-match/%d"

type (
	Timeline interface {
		ByMatch(id int64) (*DTO, error)
	}

	timeline struct {
		http http.HTTP
	}
)

func New(http http.HTTP) Timeline {
	return &timeline{
		http: http,
	}
}

func (t *timeline) ByMatch(id int64) (*DTO, error) {
	url := fmt.Sprintf(byMatchURL, id)
	var dto *DTO

	err := t.http.Get(url, &dto)

	return dto, err
}
