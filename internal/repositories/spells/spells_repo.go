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
	"github.com/victorprocure/opendominiongo/internal/dto"
	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/get_spell_no_group.sql
var getSpellNoGroupSQL string

//go:embed sql/upsert_spell.sql
var upsertSpellSQL string

const defaultRaceWhereClause = `WHERE r.key = ANY (COALESCE(s.races, ARRAY[]::text[]))`

type SpellsRepo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewSpellsRepository(db *sql.DB, log *slog.Logger) *SpellsRepo {
	return &SpellsRepo{db: db, log: log}
}

func (r *SpellsRepo) GetSpellByKeyContext(key string, ctx context.Context, tx repositories.DbTx) (*domain.Spell, error) {
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

func (r *SpellsRepo) GetSpellsByRaceKeyContext(raceKey string, ctx context.Context, tx repositories.DbTx) ([]*domain.Spell, error) {

	query := fmt.Sprintf(getSpellNoGroupSQL, defaultRaceWhereClause, `
		WHERE (
  			COALESCE(cardinality(s.races), 0) = 0                -- NULL or empty array
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

func (r *SpellsRepo) UpsertSpellForSyncContext(sp *dto.SpellYaml, ctx context.Context, tx repositories.DbTx) error {
	var perksJSON []byte
	if len(sp.Perks) > 0 {
		b, err := json.Marshal(sp.Perks)
		if err != nil {
			return fmt.Errorf("unable to marshal perks: %w", err)
		}
		perksJSON = b
	}

	var id int
	if err := tx.QueryRowContext(ctx, upsertSpellSQL,
		sp.Key, sp.Name, sp.Category, sp.ManaCost, sp.StrengthCost,
		sp.Duration, sp.Cooldown, sp.Active, pq.Array(sp.Races), perksJSON,
	).Scan(&id); err != nil {
		return fmt.Errorf("unable to prepare upsert spell statement: %w", err)
	}

	return nil
}

func scanOneSpellRow(s repositories.RowScanner) (*domain.Spell, error) {
	var (
		sp    spellRow
		races []byte
		perks []byte
	)
	if err := s.Scan(&sp.Id, &sp.Key, &sp.Name, &sp.Category,
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

func toDomain(sr *spellRow, r []byte, p []byte) (*domain.Spell, error) {
	spell := domain.Spell{
		Id:           sr.Id,
		Key:          sr.Key,
		Name:         sr.Name,
		Category:     sr.Category,
		CostMana:     sr.CostMana,
		CostStrength: sr.CostStrength,
		Duration:     sr.Duration,
		Cooldown:     sr.Cooldown,
		Active:       sr.Active,
	}

	var racesJson []dto.RaceJSON
	if err := json.Unmarshal(r, &racesJson); err != nil {
		return nil, fmt.Errorf("unable to unmarshal races: %w", err)
	}

	if len(racesJson) > 0 {
		spell.RaceKeys = make([]string, len(racesJson))
		for i, r := range racesJson {
			spell.RaceKeys[i] = r.Key
		}
	}

	var perksJson []dto.PerkJSON
	if err := json.Unmarshal(p, &perksJson); err != nil {
		return nil, fmt.Errorf("unable to unmarshal perks: %w", err)
	}

	if len(perksJson) > 0 {
		spell.Perks = make([]domain.SpellPerk, len(perksJson))
		for i, p := range perksJson {
			spell.Perks[i] = domain.SpellPerk{
				TypeKey: p.PerkType.Key,
				Value:   p.Value,
			}
		}
	}

	return &spell, nil
}
