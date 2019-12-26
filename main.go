package main

import (
	"encoding/json"
	"fmt"

	"league/summoner"
)

var configPath = "config.json"

func main() {
	di, err := newDI(configPath)

	if err != nil {
		panic(err)
	}

	me, err := di.summoner.ByName(di.config.SummonerName)

	if err != nil {
		panic(err)
	}

	match, err := di.match.ByAccount(me.AccountID)

	if err != nil {
		panic(err)
	}

	timeline, err := di.timeline.ByMatch(match.Matches[0].GameId)
	prettyPrint(timeline)

	if true {
		return
	}

	game, err := di.spectator.BySummoner(me.ID)

	if err != nil {
		panic(err)
	}

	participants := map[string]*summoner.DTO{}

	for _, participant := range game.Participants {
		it, err := di.summoner.ByName(participant.SummonerName)

		if err != nil {
			panic(err)
		}

		participants[participant.SummonerName] = it
	}

	//prettyPrint(game)
	//prettyPrint(participants)
}

func prettyPrint(v interface{}) {
	bytes, err := json.MarshalIndent(v, "", "    ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
