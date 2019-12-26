package spectator

type (
	DTO struct {
		GameID            int64                    `json:"gameId,omitempty"`
		GameStartTime     int64                    `json:"gameStartTime,omitempty"`
		PlatformID        string                   `json:"platformId,omitempty"`
		GameMode          string                   `json:"gameMode,omitempty"`
		MapID             int64                    `json:"mapId,omitempty"`
		GameType          string                   `json:"gameType,omitempty"`
		BannedChampions   []BannedChampion         `json:"bannedChampions,omitempty"`
		Observers         Observer                 `json:"observers,omitempty"`
		Participants      []CurrentGameParticipant `json:"participants,omitempty"`
		GameLength        int64                    `json:"gameLength,omitempty"`
		GameQueueConfigId int64                    `json:"gameQueueConfigId,omitempty"`
	}

	BannedChampion struct {
		PickTurn   int   `json:"pickTurn,omitempty"`
		ChampionId int64 `json:"championId,omitempty"`
		TeamID     int64 `json:"teamID,omitempty"`
	}

	Observer struct {
		EncryptionKey string `json:"encryptionKey,omitempty"`
	}

	CurrentGameParticipant struct {
		ProfileIconId            int64                     `json:"profileIconId,omitempty"`
		ChampionId               int64                     `json:"championId,omitempty"`
		SummonerName             string                    `json:"summonerName,omitempty"`
		GameCustomizationObjects []GameCustomizationObject `json:"gameCustomizationObjects,omitempty"`
		Bot                      bool                      `json:"bot,omitempty"`
		Perks                    Perks                     `json:"perks,omitempty"`
		Spell1ID                 int64                     `json:"spell1Id,omitempty"`
		Spell2ID                 int64                     `json:"spell2Id,omitempty"`
		TeamID                   int64                     `json:"teamId,omitempty"`
		SummonerID               string                    `json:"summonerId,omitempty"`
	}

	GameCustomizationObject struct {
		Category string `json:"category,omitempty"`
		Content  string `json:"content,omitempty"`
	}

	Perks struct {
		PerkStyle    int64   `json:"perkStyle,omitempty"`
		PerkIds      []int64 `json:"perkIds,omitempty"`
		PerkSubStyle int64   `json:"perkSubStyle,omitempty"`
	}
)
