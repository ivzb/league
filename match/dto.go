package match

type (
	MatchDto struct {
		SeasonID              int                      `json:"seasonId,omitempty"`
		QueueID               int                      `json:"queueId,omitempty"`
		GameID                int64                    `json:"gameId,omitempty"`
		ParticipantIdentities []ParticipantIdentityDto `json:"participantIdentities,omitempty"`
		GameVersion           string                   `json:"gameVersion,omitempty"`
		PlatformID            string                   `json:"platformId,omitempty"`
		GameMode              string                   `json:"gameMode,omitempty"`
		MapID                 int                      `json:"mapId,omitempty"`
		GameType              string                   `json:"gameType,omitempty"`
		Teams                 []TeamStatsDto           `json:"teams,omitempty"`
		Participants          []ParticipantDto         `json:"participants,omitempty"`
		GameDuration          int64                    `json:"gameDuration,omitempty"`
		GameCreation          int64                    `json:"gameCreation,omitempty"`
	}

	ParticipantIdentityDto struct {
		Player        PlayerDto `json:"player,omitempty"`
		ParticipantId int       `json:"participantId,omitempty"`
	}

	PlayerDto struct {
		CurrentPlatformId string `json:"currentPlatformId,omitempty"`
		SummonerName      string `json:"summonerName,omitempty"`
		MatchHistoryUri   string `json:"matchHistoryUri,omitempty"`
		PlatformId        string `json:"platformId,omitempty"`
		CurrentAccountId  string `json:"currentAccountId,omitempty"`
		ProfileIcon       int    `json:"profileIcon,omitempty"`
		SummonerId        string `json:"summonerId,omitempty"`
		AccountId         string `json:"accountId,omitempty"`
	}

	TeamStatsDto struct {
		FirstDragon          bool          `json:"firstDragon"`
		FirstInhibitor       bool          `json:"firstInhibitor"`
		Bans                 []TeamBansDto `json:"bans"`
		BaronKills           int           `json:"baronKills"`
		FirstRiftHerald      bool          `json:"firstRiftHerald"`
		FirstBaron           bool          `json:"firstBaron"`
		RiftHeraldKills      int           `json:"riftHeraldKills"`
		FirstBlood           bool          `json:"firstBlood"`
		TeamId               int           `json:"teamId"`
		FirstTower           bool          `json:"firstTower"`
		VilemawKills         int           `json:"vilemawKills"`
		InhibitorKills       int           `json:"inhibitorKills"`
		TowerKills           int           `json:"towerKills"`
		DominionVictoryScore int           `json:"dominionVictoryScore"`
		Win                  string        `json:"win"`
		DragonKills          int           `json:"dragonKills"`
	}

	TeamBansDto struct {
		PickTurn   int `json:"pickTurn"`
		ChampionId int `json:"championId"`
	}

	ParticipantDto struct {
		Stats                     ParticipantStatsDto    `json:"stats"`
		ParticipantId             int                    `json:"participantId"`
		Runes                     []RuneDto              `json:"runes"`
		Timeline                  ParticipantTimelineDto `json:"timeline"`
		TeamId                    int                    `json:"teamId"`
		Spell2Id                  int                    `json:"spell2Id"`
		Masteries                 []MasteryDto           `json:"masteries"`
		HighestAchievedSeasonTier string                 `json:"highestAchievedSeasonTier"`
		Spell1Id                  int                    `json:"spell1Id"`
		ChampionId                int                    `json:"championId"`
	}

	ParticipantStatsDto struct {
		FirstBloodAssist                bool  `json:"firstBloodAssist"`
		VisionScore                     int64 `json:"visionScore"`
		MagicDamageDealtToChampions     int64 `json:"magicDamageDealtToChampions"`
		DamageDealtToObjectives         int64 `json:"damageDealtToObjectives"`
		TotalTimeCrowdControlDealt      int   `json:"totalTimeCrowdControlDealt"`
		LongestTimeSpentLiving          int   `json:"longestTimeSpentLiving"`
		Perk1Var1                       int   `json:"perk1Var1"`
		Perk1Var3                       int   `json:"perk1Var3"`
		Perk1Var2                       int   `json:"perk1Var2"`
		TripleKills                     int   `json:"tripleKills"`
		Perk3Var3                       int   `json:"perk3Var3"`
		NodeNeutralizeAssist            int   `json:"nodeNeutralizeAssist"`
		Perk3Var2                       int   `json:"perk3Var2"`
		PlayerScore9                    int   `json:"playerScore9"`
		PlayerScore8                    int   `json:"playerScore8"`
		Kills                           int   `json:"kills"`
		PlayerScore1                    int   `json:"playerScore1"`
		PlayerScore0                    int   `json:"playerScore0"`
		PlayerScore3                    int   `json:"playerScore3"`
		PlayerScore2                    int   `json:"playerScore2"`
		PlayerScore5                    int   `json:"playerScore5"`
		PlayerScore4                    int   `json:"playerScore4"`
		PlayerScore7                    int   `json:"playerScore7"`
		PlayerScore6                    int   `json:"playerScore6"`
		Perk5Var1                       int   `json:"perk5Var1"`
		Perk5Var3                       int   `json:"perk5Var3"`
		Perk5Var2                       int   `json:"perk5Var2"`
		TotalScoreRank                  int   `json:"totalScoreRank"`
		NeutralMinionsKilled            int   `json:"neutralMinionsKilled"`
		DamageDealtToTurrets            int64 `json:"damageDealtToTurrets"`
		PhysicalDamageDealtToChampions  int64 `json:"physicalDamageDealtToChampions"`
		NodeCapture                     int   `json:"nodeCapture"`
		LargestMultiKill                int   `json:"largestMultiKill"`
		Perk2Var2                       int   `json:"perk2Var2"`
		Perk2Var3                       int   `json:"perk2Var3"`
		TotalUnitsHealed                int   `json:"totalUnitsHealed"`
		Perk2Var1                       int   `json:"perk2Var1"`
		Perk4Var1                       int   `json:"perk4Var1"`
		Perk4Var2                       int   `json:"perk4Var2"`
		Perk4Var3                       int   `json:"perk4Var3"`
		WardsKilled                     int   `json:"wardsKilled"`
		LargestCriticalStrike           int   `json:"largestCriticalStrike"`
		LargestKillingSpree             int   `json:"largestKillingSpree"`
		QuadraKills                     int   `json:"quadraKills"`
		TeamObjective                   int   `json:"teamObjective"`
		MagicDamageDealt                int64 `json:"magicDamageDealt"`
		Item2                           int   `json:"item2"`
		Item3                           int   `json:"item3"`
		Item0                           int   `json:"item0"`
		NeutralMinionsKilledTeamJungle  int   `json:"neutralMinionsKilledTeamJungle"`
		Item6                           int   `json:"item6"`
		Item4                           int   `json:"item4"`
		Item5                           int   `json:"item5"`
		Perk1                           int   `json:"perk1"`
		Perk0                           int   `json:"perk0"`
		Perk3                           int   `json:"perk3"`
		Perk2                           int   `json:"perk2"`
		Perk5                           int   `json:"perk5"`
		Perk4                           int   `json:"perk4"`
		Perk3Var1                       int   `json:"perk3Var1"`
		DamageSelfMitigated             int64 `json:"damageSelfMitigated"`
		MagicalDamageTaken              int64 `json:"magicalDamageTaken"`
		FirstInhibitorKill              bool  `json:"firstInhibitorKill"`
		TrueDamageTaken                 int64 `json:"trueDamageTaken"`
		NodeNeutralize                  int   `json:"nodeNeutralize"`
		Assists                         int   `json:"assists"`
		CombatPlayerScore               int   `json:"combatPlayerScore"`
		PerkPrimaryStyle                int   `json:"perkPrimaryStyle"`
		GoldSpent                       int   `json:"goldSpent"`
		TrueDamageDealt                 int64 `json:"trueDamageDealt"`
		ParticipantId                   int   `json:"participantId"`
		TotalDamageTaken                int64 `json:"totalDamageTaken"`
		PhysicalDamageDealt             int64 `json:"physicalDamageDealt"`
		SightWardsBoughtInGame          int   `json:"sightWardsBoughtInGame"`
		TotalDamageDealtToChampions     int64 `json:"totalDamageDealtToChampions"`
		PhysicalDamageTaken             int64 `json:"physicalDamageTaken"`
		TotalPlayerScore                int   `json:"totalPlayerScore"`
		Win                             bool  `json:"win"`
		ObjectivePlayerScore            int   `json:"objectivePlayerScore"`
		TotalDamageDealt                int64 `json:"totalDamageDealt"`
		Item1                           int   `json:"item1"`
		NeutralMinionsKilledEnemyJungle int   `json:"neutralMinionsKilledEnemyJungle"`
		Deaths                          int   `json:"deaths"`
		WardsPlaced                     int   `json:"wardsPlaced"`
		PerkSubStyle                    int   `json:"perkSubStyle"`
		TurretKills                     int   `json:"turretKills"`
		FirstBloodKill                  bool  `json:"firstBloodKill"`
		TrueDamageDealtToChampions      int64 `json:"trueDamageDealtToChampions"`
		GoldEarned                      int   `json:"goldEarned"`
		KillingSprees                   int   `json:"killingSprees"`
		UnrealKills                     int   `json:"unrealKills"`
		AltarsCaptured                  int   `json:"altarsCaptured"`
		FirstTowerAssist                bool  `json:"firstTowerAssist"`
		FirstTowerKill                  bool  `json:"firstTowerKill"`
		ChampLevel                      int   `json:"champLevel"`
		DoubleKills                     int   `json:"doubleKills"`
		NodeCaptureAssist               int   `json:"nodeCaptureAssist"`
		InhibitorKills                  int   `json:"inhibitorKills"`
		FirstInhibitorAssist            bool  `json:"firstInhibitorAssist"`
		Perk0Var1                       int   `json:"perk0Var1"`
		Perk0Var2                       int   `json:"perk0Var2"`
		Perk0Var3                       int   `json:"perk0Var3"`
		VisionWardsBoughtInGame         int   `json:"visionWardsBoughtInGame"`
		AltarsNeutralized               int   `json:"altarsNeutralized"`
		PentaKills                      int   `json:"pentaKills"`
		TotalHeal                       int64 `json:"totalHeal"`
		TotalMinionsKilled              int   `json:"totalMinionsKilled"`
		TimeCCingOthers                 int64 `json:"timeCCingOthers"`
	}

	RuneDto struct {
		RuneId int `json:"runeId"`
		Rank   int `json:"rank"`
	}

	ParticipantTimelineDto struct {
		Lane                        string             `json:"lane"`
		ParticipantId               int                `json:"participantId"`
		CsDiffPerMinDeltas          map[string]float64 `json:"csDiffPerMinDeltas"`
		GoldPerMinDeltas            map[string]float64 `json:"goldPerMinDeltas"`
		XpDiffPerMinDeltas          map[string]float64 `json:"xpDiffPerMinDeltas"`
		CreepsPerMinDeltas          map[string]float64 `json:"creepsPerMinDeltas"`
		XpPerMinDeltas              map[string]float64 `json:"xpPerMinDeltas"`
		Role                        string             `json:"role"`
		DamageTakenDiffPerMinDeltas map[string]float64 `json:"damageTakenDiffPerMinDeltas"`
		DamageTakenPerMinDeltas     map[string]float64 `json:"damageTakenPerMinDeltas"`
	}

	MasteryDto struct {
		MasteryId int `json:"masteryId"`
		Rank      int `json:"rank"`
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
