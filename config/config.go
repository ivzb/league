package config

import (
	"encoding/json"
	"fmt"

	"league/file"
)

type (
	Config struct {
		BaseURL      string `json:"base_url"`
		ApiKey       string `json:"api_key"`
		SummonerName string `json:"summoner_name"`
		MatchesLimit int    `json:"matches_limit"`
	}
)

func New(file file.File, path string) (*Config, error) {
	confBytes, err := file.Read(path)

	if err != nil {
		return nil, err
	}

	config := &Config{}

	if err := json.Unmarshal(confBytes, &config); err != nil {
		return nil, fmt.Errorf("could not parse config: %v", err)
	}

	return config, nil
}
