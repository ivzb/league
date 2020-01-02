package league

type (
	LeagueDTO struct {
		QueueType    string        `json:"queueType"`
		SummonerName string        `json:"summonerName"`
		HotStreak    bool          `json:"hotStreak"`
		MiniSeries   MiniSeriesDTO `json:"miniSeries"`
		Wins         int           `json:"wins"`
		Veteran      bool          `json:"veteran"`
		Losses       int           `json:"losses"`
		Rank         string        `json:"rank"`
		LeagueId     string        `json:"leagueId"`
		Inactive     bool          `json:"inactive"`
		FreshBlood   bool          `json:"freshBlood"`
		Tier         string        `json:"tier"`
		SummonerId   string        `json:"summonerId"`
		LeaguePoints int           `json:"leaguePoints"`
	}

	MiniSeriesDTO struct {
		Progress string `json:"progress"`
		Losses   int    `json:"losses"`
		Target   int    `json:"target"`
		Wins     int    `json:"wins"`
	}
)
