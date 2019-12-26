package timeline

type (
	DTO struct {
		Frames        []Frame `json:"frames,omitempty"`
		FrameInterval int64   `json:"frameInterval,omitempty"`
	}

	Frame struct {
		Timestamp         int64                       `json:"timestamp,omitempty"`
		ParticipantFrames map[string]ParticipantFrame `json:"participantFrames,omitempty"`
		Events            []Event                     `json:"events,omitempty"`
	}

	ParticipantFrame struct {
		TotalGold           int      `json:"totalGold,omitempty"`
		TeamScore           int      `json:"teamScore,omitempty"`
		ParticipantId       int      `json:"participantId,omitempty"`
		Level               int      `json:"level,omitempty"`
		CurrentGold         int      `json:"currentGold,omitempty"`
		MinionsKilled       int      `json:"minionsKilled,omitempty"`
		DominionScore       int      `json:"dominionScore,omitempty"`
		Position            Position `json:"position,omitempty"`
		XP                  int      `json:"xp,omitempty"`
		JungleMinionsKilled int      `json:"jungleMinionsKilled,omitempty"`
	}

	Position struct {
		X int `json:"x,omitempty"`
		Y int `json:"y,omitempty"`
	}

	// type (Legal values: CHAMPION_KILL, WARD_PLACED, WARD_KILL, BUILDING_KILL, ELITE_MONSTER_KILL, ITEM_PURCHASED, ITEM_SOLD, ITEM_DESTROYED, ITEM_UNDO, SKILL_LEVEL_UP, ASCENDED_EVENT, CAPTURE_POINT, PORO_KING_SUMMON)
	Event struct {
		EventType               string   `json:"eventType,omitempty"`
		TowerType               string   `json:"towerType,omitempty"`
		TeamId                  int      `json:"teamId,omitempty"`
		AscendedType            string   `json:"ascendedType,omitempty"`
		KillerId                int      `json:"killerId,omitempty"`
		LevelUpType             string   `json:"levelUpType,omitempty"`
		PointCaptured           string   `json:"pointCaptured,omitempty"`
		AssistingParticipantIds []int    `json:"assistingParticipantIds,omitempty"`
		WardType                string   `json:"wardType,omitempty"`
		MonsterType             string   `json:"monsterType,omitempty"`
		Type                    string   `json:",omitempty"`
		SkillSlot               int      `json:"skillSlot,omitempty"`
		VictimId                int      `json:"victimId,omitempty"`
		Timestamp               int64    `json:"timestamp,omitempty"`
		AfterId                 int      `json:"afterId,omitempty"`
		MonsterSubType          string   `json:"monsterSubType,omitempty"`
		LaneType                string   `json:"laneType,omitempty"`
		ItemId                  int      `json:"itemId,omitempty"`
		ParticipantId           int      `json:"participantId,omitempty"`
		BuildingType            string   `json:"buildingType,omitempty"`
		CreatorId               int      `json:"creatorId,omitempty"`
		Position                Position `json:"position,omitempty"`
		BeforeId                int      `json:"beforeId,omitempty"`
	}
)
