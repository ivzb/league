package summoner

type DTO struct {
	ProfileIconID int    `json:"profileIconId,omitempty"`
	Name          string `json:"name,omitempty"`
	PuuID         string `json:"puuid,omitempty"`
	SummonerLevel int64  `json:"summonerLevel,omitempty"`
	RevisionDate  int64  `json:"revisionDate,omitempty"`
	ID            string `json:"id,omitempty"`
	AccountID     string `json:"accountId,omitempty"`
}
