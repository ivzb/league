package main

import (
	"encoding/json"
	"fmt"
	"time"

	"league/match"
	"league/summoner"
	"league/timeline"
)

const (
	configPath = "config.json"
	day        = 86400
)

func main() {
	di, err := newDI(configPath)

	if err != nil {
		panic(err)
	}

	champs, err := di.champion.Map()

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

	champions := map[string]int{}

	for _, match := range match.Matches {
		timestamp := time.Unix(match.Timestamp/1000, 0)
		now := time.Now().Unix()
		limit := now - (now % day)

		if timestamp.Unix() >= limit {
			fmt.Println(timestamp)
		}

		champion := champs[match.Champion]

		if _, ok := champions[champion]; !ok {
			champions[champion] = 0
		}

		champions[champion]++
	}

	out := map[string]string{}

	// print wins-loses, and types of games

	for champion, games := range champions {
		out[champion] = fmt.Sprintf("%d games (%.0f%%)", games, float64(games) / float64(len(match.Matches)) * 100)
	}

	prettyPrint(out)
}

func prettyPrint(v interface{}) {
	bytes, err := json.MarshalIndent(v, "", "    ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func getTimeline(di *di, match *match.MatchlistDto) *timeline.DTO {
	timeline, err := di.timeline.ByMatch(match.Matches[0].GameId)

	if err != nil {
		panic(err)
	}

	return timeline
}

func spectate(di *di, me *summoner.DTO) map[string]*summoner.DTO {
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

	return participants
}
