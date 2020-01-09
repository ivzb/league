package main

import (
	"net/http"

	"league/champion"
	"league/league"
	"league/match"
	"league/spectator"
	"league/summoner"
	"league/timeline"
	"league/util/config"
	"league/util/file"
	h "league/util/http"
)

type web struct {
	file   file.File
	config *config.Config
	http   h.HTTP

	champion  champion.Champion
	summoner  summoner.Summoner
	spectator spectator.Spectator
	match     match.Match
	timeline  timeline.Timeline
	league    league.League
}

func newWeb(configPath string) (*web, error) {
	web := &web{}
	var err error

	web.file = file.New()
	web.config, err = config.New(web.file, configPath)

	if err != nil {
		return nil, err
	}

	web.http = h.New(web.config)
	web.champion = champion.New(web.http)
	web.summoner = summoner.New(web.http)
	web.spectator = spectator.New(web.http)
	web.match = match.New(web.http, web.config)
	web.timeline = timeline.New(web.http)
	web.league = league.New(web.http)

	return web, nil
}

func (web *web) run() error {
	http.HandleFunc("/summoner", web.summoner.Web().BySummonerName)
	http.HandleFunc("/leagues", web.league.Web().BySummonerId)
	http.HandleFunc("/match", web.match.Web().ByMatchID)
	http.HandleFunc("/matches", web.match.Web().ByAccountID)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	return http.ListenAndServe(":8080", nil)
}
