package spells

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/lib/pq"
	"github.com/victorprocure/opendominiongo/internal/domain"
	"github.com/victorprocure/opendominiongo/internal/helpers"
	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/get_spell_no_group.sql
var getSpellNoGroupSQL string

//go:embed sql/upsert_spell.sql
var upsertSpellSQL string

const defaultRaceWhereClause = `WHERE r.key = ANY (COALESCE(s.races, ARRAY[]::text[]))`

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewSpellsRepository(db *sql.DB, log *slog.Logger) *Repo {
	return &Repo{db: db, log: log}
}

func (r *Repo) GetSpellByKeyContext(ctx context.Context, tx repositories.DbTx, key string) (*domain.Spell, error) {
	query := fmt.Sprintf(getSpellNoGroupSQL, defaultRaceWhereClause, "WHERE s.key = $1")
	spell, err := scanOneSpellRow(tx.QueryRowContext(ctx, query, key))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("unable to get spell by key: %s, error: %w", key, err)
	}

	return spell, nil
}

func (r *Repo) GetSpellsByRaceKeyContext(ctx context.Context, tx repositories.DbTx, raceKey string) ([]*domain.Spell, error) {
	query := fmt.Sprintf(getSpellNoGroupSQL, defaultRaceWhereClause, `
		WHERE (
  			COALESCE(cardinality(s.races), 0) = 0           -- NULL or empty array
  			OR s.races @> ARRAY[$1]::text[]                 -- contains raceKey
		)`)
	rows, err := tx.QueryContext(ctx, query, raceKey)
	if err != nil {
		return nil, fmt.Errorf("unable to get spells: %w", err)
	}
	defer rows.Close()

	var spells []*domain.Spell
	for rows.Next() {
		spell, err := scanOneSpellRow(rows)
		if err != nil {
			return nil, fmt.Errorf("unable to scan spell row: %w", err)
		}
		spells = append(spells, spell)
	}

	return spells, rows.Err()
}

// UpsertArgs is the normalized contract for upserting a spell.
type UpsertArgs struct {
	Key          string
	Name         string
	Category     string
	ManaCost     float64
	StrengthCost float64
	Duration     int
	Cooldown     int
	Active       bool
	Races        []string
	Perks        map[string]string
}

func (r *Repo) UpsertFromSyncContext(ctx context.Context, tx repositories.DbTx, a UpsertArgs) error {
	var perksJSON []byte
	if len(a.Perks) > 0 {
		var err error
		perksJSON, err = helpers.MarshalPerksAsJSONArrayFromMap(a.Perks)
		if err != nil {
			return fmt.Errorf("unable to marshal perks: %w", err)
		}
	}

	var id int
	// Log the input JSON for debugging if available
	if r.log != nil {
		r.log.Info("upsert spell input", "key", a.Key, "perksJSON", string(perksJSON))
	}

	if err := tx.QueryRowContext(ctx, upsertSpellSQL,
		a.Key, a.Name, a.Category, a.ManaCost, a.StrengthCost,
		a.Duration, a.Cooldown, a.Active, pq.Array(a.Races), perksJSON,
	).Scan(&id); err != nil {
		return fmt.Errorf("unable to prepare upsert spell statement for key %s: %w", a.Key, err)
	}

	return nil
}

func scanOneSpellRow(s repositories.RowScanner) (*domain.Spell, error) {
	var (
		sp    spellRow
		races []byte
		perks []byte
	)
	if err := s.Scan(&sp.ID, &sp.Key, &sp.Name, &sp.Category,
		&sp.CostMana, &sp.CostStrength, &sp.Duration, &sp.Cooldown,
		&sp.Active, &races, &perks); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("unable to scan spell row: %w", err)
	}

	spell, err := toDomain(&sp, races, perks)
	if err != nil {
		return nil, fmt.Errorf("unable to convert spell row to domain: %w", err)
	}
	return spell, nil
}

func toDomain(sr *spellRow, r, p []byte) (*domain.Spell, error) {
	spell := domain.Spell{
		ID:           sr.ID,
		Key:          sr.Key,
		Name:         sr.Name,
		Category:     sr.Category,
		CostMana:     sr.CostMana,
		CostStrength: sr.CostStrength,
		Duration:     sr.Duration,
		Cooldown:     sr.Cooldown,
		Active:       sr.Active,
	}

	var racesJSON []struct {
		Key string `json:"key"`
	}
	if err := json.Unmarshal(r, &racesJSON); err != nil {
		return nil, fmt.Errorf("unable to unmarshal races: %w", err)
	}

	if len(racesJSON) > 0 {
		spell.RaceKeys = make([]string, len(racesJSON))
		for i, r := range racesJSON {
			spell.RaceKeys[i] = r.Key
		}
	}

	var perksJSON []struct {
		PerkType struct {
			Key string `json:"key"`
		} `json:"perk_type"`
		Value string `json:"value"`
	}
	if err := json.Unmarshal(p, &perksJSON); err != nil {
		return nil, fmt.Errorf("unable to unmarshal perks: %w", err)
	}

	if len(perksJSON) > 0 {
		spell.Perks = make([]domain.SpellPerk, len(perksJSON))
		for i, p := range perksJSON {
			spell.Perks[i] = domain.SpellPerk{
				TypeKey: p.PerkType.Key,
				Value:   p.Value,
			}
		}
	}

	return &spell, nil
}
