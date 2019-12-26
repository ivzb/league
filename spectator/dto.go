package spectator

type (
	DTO struct {
		GameID            int64                    `json:"gameId"`
		GameStartTime     int64                    `json:"gameStartTime"`
		PlatformID        string                   `json:"platformId"`
		GameMode          string                   `json:"gameMode"`
		MapID             int64                    `json:"mapId"`
		GameType          string                   `json:"gameType"`
		BannedChampions   []BannedChampion         `json:"bannedChampions"`
		Observers         Observer                 `json:"observers"`
		Participants      []CurrentGameParticipant `json:"participants"`
		GameLength        int64                    `json:"gameLength"`
		GameQueueConfigId int64                    `json:"gameQueueConfigId"`
	}

	BannedChampion struct {
		PickTurn   int   `json:"pickTurn"`
		ChampionId int64 `json:"championId"`
		TeamID     int64 `json:"teamID"`
	}

	Observer struct {
		EncryptionKey string `json:"encryptionKey"`
	}

	CurrentGameParticipant struct {
		ProfileIconId            int64                     `json:"profileIconId"`
		ChampionId               int64                     `json:"championId"`
		SummonerName             string                    `json:"summonerName"`
		GameCustomizationObjects []GameCustomizationObject `json:"gameCustomizationObjects"`
		Bot                      bool                      `json:"bot"`
		Perks                    Perks                     `json:"perks"`
		Spell1ID                 int64                     `json:"spell1Id"`
		Spell2ID                 int64                     `json:"spell2Id"`
		TeamID                   int64                     `json:"teamId"`
		SummonerID               string                    `json:"summonerId"`
	}

	GameCustomizationObject struct {
		Category string `json:"category"`
		Content  string `json:"content"`
	}

	Perks struct {
		PerkStyle    int64   `json:"perkStyle"`
		PerkIds      []int64 `json:"perkIds"`
		PerkSubStyle int64   `json:"perkSubStyle"`
	}
)
