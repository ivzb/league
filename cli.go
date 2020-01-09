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

type cli struct {
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

func newCli(configPath string) (*cli, error) {
	cli := &cli{}
	var err error

	cli.file = file.New()
	cli.config, err = config.New(cli.file, configPath)

	if err != nil {
		return nil, err
	}

	cli.http = http.New(cli.config)
	cli.champion = champion.New(cli.http)
	cli.summoner = summoner.New(cli.http)
	cli.spectator = spectator.New(cli.http)
	cli.match = match.New(cli.http, cli.config)
	cli.timeline = timeline.New(cli.http)
	cli.league = league.New(cli.http)

	return cli, nil
}

func (cli *cli) run() error {
	champs, err := cli.champion.Map()

	if err != nil {
		return err
	}

	summonerNames, err := participant.New(cli.file, cli.config.ParticipantsFile)

	if err != nil {
		return err
	}

	for _, summonerName := range summonerNames {
		summoner, err := cli.summoner.Repo().BySummonerName(summonerName)

		if err != nil {
			return err
		}

		cli.leagues(summoner.ID)

		matchlist, err := cli.match.Repo().ByAccountID(summoner.AccountID, cli.config.MatchesLimit)

		if err != nil {
			return err
		}

		cli.champions(matchlist, summoner, champs)
	}

	return nil
}

func (cli *cli) champions(matchlist *match.MatchlistDto, summoner *summoner.DTO, champs map[int]string) error {
	games := map[int]int{}
	lanes := map[string]int{}
	wins, diffs, err := cli.wins(matchlist, summoner)

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

func (cli *cli) wins(matchlist *match.MatchlistDto, me *summoner.DTO) (map[int]int, []string, error) {
	wins := map[int]int{}
	diff := ""
	todaysDiff := ""

	for _, m := range matchlist.Matches {
		match, err := cli.match.Repo().ByMatchID(strconv.FormatInt(m.GameId, 10))

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

func (cli *cli) leagues(summonerID string) error {
	leagues, err := cli.league.Repo().BySummonerId(summonerID)

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

func (cli *cli) getTimeline(match *match.MatchlistDto) (*timeline.DTO, error) {
	return cli.timeline.ByMatch(match.Matches[0].GameId)
}

func (cli *cli) spectate(me *summoner.DTO) (map[string]*summoner.DTO, error) {
	game, err := cli.spectator.BySummoner(me.ID)

	if err != nil {
		return nil, err
	}

	participants := map[string]*summoner.DTO{}

	for _, participant := range game.Participants {
		summoner, err := cli.summoner.Repo().BySummonerName(participant.SummonerName)

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
