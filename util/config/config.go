package config

import (
	"encoding/json"
	"fmt"

	"league/util/file"
)

type (
	Config struct {
		BaseURL          string `json:"base_url"`
		ApiKeyFile       string `json:"api_key_file"`
		SummonerName     string `json:"summoner_name"`
		MatchesLimit     int    `json:"matches_limit"`
		ParticipantsFile string `json:"participants_file"`

		ApiKey string
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

	apiKeyBytes, err := file.Read(config.ApiKeyFile)

	if err != nil {
		return nil, fmt.Errorf("could not find api key file in %s", config.ApiKeyFile)
	}

	config.ApiKey = string(apiKeyBytes)

	return config, nil
}
