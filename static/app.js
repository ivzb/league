const summoners = document.getElementById("summoners");
const scan = document.getElementById("scan");
const info = document.getElementById("info");

function run() {
    scan.onclick = function (e) {
        info.innerHTML = "";

        const lines = summoners.value.split("\n");

        for (const i in lines) {
            const line = lines[i];
            const words = line.split("joined");

            loadSummoner(words[0].trim());
        }

        e.preventDefault();
    };
}

function loadSummoner(summonerName) {
    get("summoner?name=" + summonerName, (summoner) => {
        const overview = {};
        overview.id = summoner.id;
        overview.name = summoner.name;
        overview.level = summoner.summonerLevel;

        get("leagues?summoner_id=" + overview.id, (leagues) => {
            for (const i in leagues) {
                const league = leagues[i];

                if (league.queueType !== "RANKED_SOLO_5x5") {
                    continue
                }

                overview.league = {};
                overview.league.tier = league.tier;
                overview.league.rank = league.rank;
                overview.league.wins = league.wins;
                overview.league.losses = league.losses;
                overview.league.leaguePoints = league.leaguePoints;
                break
            }

            get("matches?account_id=" + summoner.accountId, (matchesData) => {
                const matches = matchesData.matches;

                overview.games = {};
                overview.lanes = {};
                overview.wins = {};
                overview.history = [];
                overview.out = {};

                for (const i in matches) {
                    const m = matches[i];

                    const champion = m.champion;

                    if (!overview.games[champion]) {
                        overview.games[champion] = 0;
                    }

                    overview.games[champion]++;

                    if (!overview.lanes[m.lane]) {
                        overview.lanes[m.lane] = 0;
                    }

                    overview.lanes[m.lane]++;

                    get("match?match_id=" + m.gameId, (match) => {
                        for (const j in match.participantIdentities) {
                            const participantIdentity = match.participantIdentities[j];

                            if (participantIdentity.player.summonerId !== summoner.id) {
                                continue
                            }

                            for (const k in match.participants) {
                                const participant = match.participants[k];

                                if (participant.participantId !== participantIdentity.participantId) {
                                    continue
                                }

                                const champion = participant.championId

                                if (!overview.wins[champion]) {
                                    overview.wins[champion] = 0;
                                }

                                for (const l in match.teams) {
                                    const team = match.teams[l];

                                    if (team.TeamId !== participant.TeamId) {
                                        continue
                                    }

                                    if (team.win === "Win") {
                                        overview.wins[participant.championId]++;
                                    }

                                    overview.history.push({
                                        isWin: team.win === "Win",
                                        gameCreation: match.gameCreation
                                    });

                                    break
                                }

                                break;
                            }

                            break;
                        }
                    });

                    // out := map[string]string{}
                    //
                    // for champion, games := range games {
                    //     total := float64(games) / float64(len(matchlist.Matches)) * 100
                    //     winRate := float64(wins[champion]) / float64(games) * 100
                    //     out[champs[champion]] = fmt.Sprintf("%d games (%.0f%%) %d wins (%.0f%% win rate)", games, total, wins[champion], winRate)
                    // }

                    // print.Pretty(out)
                }

                // console.log(overview);
                // console.log(matches);

                info.innerHTML += '<div class="summoner">' +
                    '<p><a href="https://eune.op.gg/summoner/userName=' + overview.name + '" target="_blank">' + overview.name + '</a></p>' +
                    '<p>Level: ' + overview.level + '</p>'
                    '</div>';
                console.log(overview);

            });
        });
    });
}

function get(url, callback) {
    const request = new XMLHttpRequest();
    request.addEventListener("load", function () {
        const response = JSON.parse(this.responseText);

        if (response.error) {
            console.error(response);
        }

        callback(response);
    });
    request.open("GET", url);
    request.send();
}

run();