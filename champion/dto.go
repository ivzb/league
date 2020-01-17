package champion

type (
	DTO struct {
		Type    string          `json:"type,omitempty"`
		Format  string          `json:"format,omitempty"`
		Version string          `json:"version,omitempty"`
		Data    map[string]Data `json:"data,omitempty"`
	}

	Data struct {
		Version string   `json:"version,omitempty"`
		ID      string   `json:"id,omitempty"`
		Key     string   `json:"key,omitempty"`
		Name    string   `json:"name,omitempty"`
		Title   string   `json:"title,omitempty"`
		Blurb   string   `json:"blurb,omitempty"`
		Info    Info     `json:"info,omitempty"`
		Image   Image    `json:"image,omitempty"`
		Tags    []string `json:"tags,omitempty"`
		Partype string   `json:"partype,omitempty"`
		Stats   Stats    `json:"stats,omitempty"`
	}

	Info struct {
		Attack     int `json:"attack,omitempty"`
		Defense    int `json:"defense,omitempty"`
		Magic      int `json:"magic,omitempty"`
		Difficulty int `json:"difficulty,omitempty"`
	}

	Image struct {
		Full   string `json:"full,omitempty"`
		Sprite string `json:"sprite,omitempty"`
		Group  string `json:"group,omitempty"`
		X      int    `json:"x,omitempty"`
		Y      int    `json:"y,omitempty"`
		W      int    `json:"w,omitempty"`
		H      int    `json:"h,omitempty"`
	}

	Stats struct {
		HP                   float64 `json:"hp,omitempty"`
		HPPerLevel           float64 `json:"hpperlevel,omitempty"`
		MoveSpeed            float64 `json:"movespeed,omitempty"`
		Armor                float64 `json:"armor,omitempty"`
		ArmorPerLevel        float64 `json:"armorperlevel,omitempty"`
		SpellBlock           float64 `json:"spellblock,omitempty"`
		SpellBlockPerLevel   float64 `json:"spellblockperlevel,omitempty"`
		AttackRange          float64 `json:"attachrange,omitempty"`
		HPRegen              float64 `json:"hpregen,omiempty"`
		HPRegenPerLevel      float64 `json:"hpregen,omiempty"`
		Mpregen              float64 `json:"mpregen,omiempty"`
		MpregenPerLevel      float64 `json:"mpregen,omiempty"`
		Crit                 float64 `json:"crit,omiempty"`
		CritPerLevel         float64 `json:"crit,omiempty"`
		AttackDamage         float64 `json:"attackdamage,omiempty"`
		AttackDamagePerLevel float64 `json:"attackdamageperlevel,omiempty"`
		AttackSpeed          float64 `json:"attackspeed,omiempty"`
	}
)
