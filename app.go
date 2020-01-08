package main

import (
	"fmt"
	"strconv"
	s "strings"

	"league/champion"
	"league/league"
	"league/match"
	"league/spectator"
	"league/summoner"
	"league/timeline"
	"league/util/config"
	"league/util/file"
	"league/util/http"
	"league/util/participant"
	"league/util/print"
	"league/util/strings"
	"league/util/time"
)

const (
	soloQ  = "RANKED_SOLO_5x5"
	win    = "Win"
	top    = "TOP"
	jungle = "JUNGLE"
	mid    = "MID"
	bottom = "BOTTOM"
	none   = "NONE"
)

type app struct {
	file   file.File
	config *config.Config
	http   http.HTTP

	champion  champion.Champion
	summoner  summoner.Summoner
	spectator spectator.Spectator
	match     match.Match
	timeline  timeline.Timeline
	league    league.League
}

func newApp(configPath string) (*app, error) {
	app := &app{}
	var err error

	app.file = file.New()
	app.config, err = config.New(app.file, configPath)

	if err != nil {
		return nil, err
	}

	app.http = http.New(app.config)
	app.champion = champion.New(app.http)
	app.summoner = summoner.New(app.http)
	app.spectator = spectator.New(app.http)
	app.match = match.New(app.http)
	app.timeline = timeline.New(app.http)
	app.league = league.New(app.http)

	return app, nil
}

func (app *app) run() error {
	champs, err := app.champion.Map()

	if err != nil {
		return err
	}

	summonerNames, err := participant.New(app.file, app.config.ParticipantsFile)

	if err != nil {
		return err
	}

	for _, summonerName := range summonerNames {
		summoner, err := app.summoner.ByName(summonerName)

		if err != nil {
			return err
		}

		app.leagues(summoner.ID)

		matchlist, err := app.match.ByAccountID(summoner.AccountID, app.config.MatchesLimit)

		if err != nil {
			return err
		}

		app.champions(matchlist, summoner, champs)
	}

	return nil
}

func (app *app) champions(matchlist *match.MatchlistDto, summoner *summoner.DTO, champs map[int]string) error {
	games := map[int]int{}
	lanes := map[string]int{}
	wins, diffs, err := app.wins(matchlist, summoner)

	if err != nil {
		return err
	}

	for _, match := range matchlist.Matches {
		champion := match.Champion

		if _, ok := games[champion]; !ok {
			games[champion] = 0
		}

		games[champion]++

		if _, ok := lanes[match.Lane]; !ok {
			lanes[match.Lane] = 0
		}

		lanes[match.Lane]++
	}

	fmt.Printf("lanes: %s\nlast %d: %s\ntoday's: %s\n", formatLanes(lanes), len(matchlist.Matches), diffs[0], diffs[1])

	out := map[string]string{}

	for champion, games := range games {
		total := float64(games) / float64(len(matchlist.Matches)) * 100
		winRate := float64(wins[champion]) / float64(games) * 100
		out[champs[champion]] = fmt.Sprintf("%d games (%.0f%%) %d wins (%.0f%% win rate)", games, total, wins[champion], winRate)
	}

	return print.Pretty(out)
}

func (app *app) wins(matchlist *match.MatchlistDto, me *summoner.DTO) (map[int]int, []string, error) {
	wins := map[int]int{}
	diff := ""
	todaysDiff := ""

	for _, m := range matchlist.Matches {
		match, err := app.match.ByMatchID(strconv.FormatInt(m.GameId, 10))

		if err != nil {
			return nil, nil, err
		}

		for _, participantIdentity := range match.ParticipantIdentities {
			if participantIdentity.Player.SummonerId != me.ID {
				continue
			}

			for _, participant := range match.Participants {
				if participant.ParticipantId != participantIdentity.ParticipantId {
					continue
				}

				champion := participant.ChampionId

				if _, ok := wins[champion]; !ok {
					wins[champion] = 0
				}

				for _, team := range match.Teams {
					if team.TeamId != participant.TeamId {
						continue
					}

					char := "-"

					if team.Win == win {
						wins[participant.ChampionId]++
						char = "+"
					}

					diff = diff + char

					if time.IsToday(match.GameCreation) {
						todaysDiff = todaysDiff + char
					}

					break
				}

				break
			}

			break
		}
	}

	diff = strings.Reverse(diff)
	todaysDiff = strings.Reverse(todaysDiff)

	return wins, []string{diff, todaysDiff}, nil
}

func (app *app) leagues(summonerID string) error {
	leagues, err := app.league.BySummonerId(summonerID)

	if err != nil {
		return err
	}

	for _, league := range leagues {
		if league.QueueType != soloQ {
			continue
		}

		fmt.Printf("%s: %s %s wins(%d) losses(%d) winrate(%.0f%%) hotStreak: %t\n", league.SummonerName, league.Tier, league.Rank, league.Wins, league.Losses, float64(league.Wins)/float64(league.Wins+league.Losses)*100, league.HotStreak)
		break
	}

	return nil
}

func (app *app) getTimeline(match *match.MatchlistDto) (*timeline.DTO, error) {
	return app.timeline.ByMatch(match.Matches[0].GameId)
}

func (app *app) spectate(me *summoner.DTO) (map[string]*summoner.DTO, error) {
	game, err := app.spectator.BySummoner(me.ID)

	if err != nil {
		return nil, err
	}

	participants := map[string]*summoner.DTO{}

	for _, participant := range game.Participants {
		summoner, err := app.summoner.ByName(participant.SummonerName)

		if err != nil {
			return nil, err
		}

		participants[participant.SummonerName] = summoner
	}

	return participants, nil
}

func formatLanes(lanes map[string]int) string {
	laneNames := []string{top, jungle, mid, bottom, none}
	laneNamesMap := map[string]bool{}
	total := 0

	for _, laneName := range laneNames {
		laneNamesMap[laneName] = true
		total += lanes[laneName]
	}

	for lane := range lanes {
		if _, ok := laneNamesMap[lane]; !ok {
			panic(fmt.Sprintf("lane %s is not defined\n", lane))
		}
	}

	result := ""

	for _, laneName := range laneNames {
		if _, ok := lanes[laneName]; !ok {
			continue
		}

		result = fmt.Sprintf("%s, %s: %d (%.0f%%)", result, laneName, lanes[laneName], float64(lanes[laneName])/float64(total)*100)
	}

	return s.Trim(result, ", ")
}
