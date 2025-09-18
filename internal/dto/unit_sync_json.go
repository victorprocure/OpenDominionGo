package dto

type UnitSyncJSON struct {
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	NeedBoat     bool      `json:"need_boat"`
	CostPlatinum int       `json:"cost_platinum"`
	CostOre      int       `json:"cost_ore"`
	CostLumber   int       `json:"cost_lumber"`
	CostGems     int       `json:"cost_gems"`
	CostMana     int       `json:"cost_mana"`
	PowerOffense int       `json:"power_offense"`
	PowerDefense int       `json:"power_defense"`
	Perks        KeyValues `json:"perks"`
}
