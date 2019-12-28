package match

type (
	MatchDto struct {
		SeasonID int `json:"seasonId,omitempty"`
		QueueID int `json:"queueId,omitempty"`
		GameID int64 `json:"gameId,omitempty"`
		ParticipantIdentities []ParticipantIdentityDto `json:"participantIdentities,omitempty"`
		GameVersion string `json:"gameVersion,omitempty"`
		PlatformID string `json:"platformId,omitempty"`
		GameMode string `json:"gameMode,omitempty"`
		MapID int `json:"mapId,omitempty"`
		GameType string `json:"gameType,omitempty"`
		Teams []TeamStatsDto `json:"teams,omitempty"`
		Participants []ParticipantDto `json:"participants,omitempty"`
		GameDuration int64 `json:"gameDuration,omitempty"`
		GameCreation int64 `json:"gameCreation,omitempty"`
	}

	MatchlistDto struct {
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
