package summoner

type DTO struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	PuuID         string `json:"puuid"`
	SummonerLevel int64  `json:"summonerLevel"`
	RevisionDate  int64  `json:"revisionDate"`
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
}
