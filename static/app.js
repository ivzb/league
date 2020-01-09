function run() {
    loadSummoner("ivzb")
}

function loadSummoner(summonerName) {
    get("summoner?name=" + summonerName, (summoner) => {
        const overview = {};
        overview.id = summoner.id;
        overview.name = summoner.name;
        overview.level = summoner.level;

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

                for (const i in matches) {
                    const m = matches[i];

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

                                    var char = "-";

                                    if (team.win === "Win") {
                                        overview.wins[participant.championId]++;
                                        char = "+"
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
                }

                // console.log(overview);
                // console.log(matches);

                console.log(overview);
            });
        });
    });
}

function get(url, callback) {
    const request = new XMLHttpRequest();
    request.addEventListener("load", function() {
        const response = JSON.parse(this.responseText);
        callback(response);
    });
    request.open("GET", url);
    request.send();
}

run();