package timeline

import (
	"encoding/json"
	"fmt"

	"league/http"
)

const byMatchPath = "lol/match/v4/timelines/by-match/%d"

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
	url := fmt.Sprintf(byMatchPath, id)

	response, err := t.http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var dto *DTO

	if err := json.NewDecoder(response.Body).Decode(&dto); err != nil {
		return nil, err
	}

	return dto, nil
}
