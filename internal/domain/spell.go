package domain

type Spell struct {
	ID           int
	Key          string
	Name         string
	Category     string
	CostMana     float64
	CostStrength float64
	Duration     int
	Cooldown     int
	Active       bool
	RaceKeys     []string
	Perks        []SpellPerk
}

type SpellPerk struct {
	TypeKey string
	Value   string
}

func NewSpell(key, name string) *Spell {
	return &Spell{Key: key, Name: name, Active: true}
}

func (s *Spell) IsGlobal() bool { return len(s.RaceKeys) == 0 }
