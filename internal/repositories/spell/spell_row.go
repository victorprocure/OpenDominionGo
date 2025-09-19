package spell

import "time"

type spellRow struct {
	ID           int       `db:"id"`
	Key          string    `db:"key"`
	Name         string    `db:"name"`
	Category     string    `db:"category"`
	CostMana     float64   `db:"cost_mana"`
	CostStrength float64   `db:"cost_strength"`
	Duration     int       `db:"duration"`
	Cooldown     int       `db:"cooldown"`
	Active       bool      `db:"active"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	Races        []string  `db:"-"`
}
