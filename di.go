package main

import (
	"league/config"
	"league/file"
	"league/http"
	"league/spectator"
	"league/summoner"
)

type di struct {
	config *config.Config
	http   http.HTTP

	summoner summoner.Summoner
	spectator spectator.Spectator
}

func newDI(configPath string) (*di, error) {
	di := &di{}
	var err error

	file := file.New()

	di.config, err = config.New(file, configPath)

	if err != nil {
		return nil, err
	}

	di.http = http.New()
	di.summoner = summoner.New(di.config, di.http)
	di.spectator= spectator.New(di.config, di.http)

	return di, nil
}
