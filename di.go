package main

import (
	"league/config"
	"league/file"
	"league/http"
	"league/match"
	"league/spectator"
	"league/summoner"
)

type di struct {
	config *config.Config
	http   http.HTTP

	summoner  summoner.Summoner
	spectator spectator.Spectator
	match match.Match
}

func newDI(configPath string) (*di, error) {
	di := &di{}
	var err error

	file := file.New()

	di.config, err = config.New(file, configPath)

	if err != nil {
		return nil, err
	}

	di.http = http.New(di.config)
	di.summoner = summoner.New(di.http)
	di.spectator = spectator.New(di.http)
	di.match = match.New(di.http)

	return di, nil
}
