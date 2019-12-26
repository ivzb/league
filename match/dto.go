package match

type (
	DTO struct {
		Matches    []MatchReference `json:"matches,omitempty"`
		TotalGames int              `json:"totalGames,omitempty"`
		StartIndex int              `json:"startIndex,omitempty"`
		EndIndex   int              `json:"endIndex,omitempty"`
	}

	MatchReference struct {
		Lane       string `json:"lane,omitempty"`
		GameId     int64  `json:"gameId,omitempty"`
		Champion   int    `json:"champion,omitempty"`
		PlatformId string `json:"platformId,omitempty"`
		Season     int    `json:"season,omitempty"`
		Queue      int    `json:"queue,omitempty"`
		Role       string `json:"role,omitempty"`
		Timestamp  int64  `json:"timestamp,omitempty"`
	}
)
