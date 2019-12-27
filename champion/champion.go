package champion

import (
	"encoding/json"
	"strconv"

	"league/http"
)

// get last version from https://ddragon.leagueoflegends.com/api/versions.json
const championsPath = "http://ddragon.leagueoflegends.com/cdn/9.24.2/data/en_US/champion.json"

type (
	Champion interface {
		All() (*DTO, error)
		Map() (map[int]string, error)
	}

	champion struct {
		http http.HTTP
	}
)

func New(http http.HTTP) Champion {
	return &champion{
		http: http,
	}
}

func (c *champion) All() (*DTO, error) {
	url := championsPath

	response, err := c.http.Get(url)

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

func (c *champion) Map() (map[int]string, error) {
	all, err := c.All()

	if err != nil {
		return nil, err
	}

	champions := map[int]string{}

	for name, champion := range all.Data {
		key, err := strconv.Atoi(champion.Key)

		if err != nil {
			return nil, err
		}

		champions[key] = name
	}

	return champions, nil
}
