package match

type (
	DTO struct {
		Matches    []MatchReference `json:"matches"`
		TotalGames int              `json:"totalGames"`
		StartIndex int              `json:"startIndex"`
		EndIndex   int              `json:"endIndex"`
	}

	MatchReference struct {
		Lane       string `json:"lane"`
		GameId     int64  `json:"gameId"`
		Champion   int    `json:"champion"`
		PlatformId string `json:"platformId"`
		Season     int    `json:"season"`
		Queue      int    `json:"queue"`
		Role       string `json:"role"`
		Timestamp  int64  `json:"timestamp"`
	}
)
