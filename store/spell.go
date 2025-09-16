package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/victorprocure/opendominiongo/helpers"
)

type Spell struct {
	Id           int
	Key          string
	Name         string
	Category     string
	CostMana     float64
	CostStrength float64
	Duration     int
	Cooldown     int
	Active       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Races        []*Race
	Perks        []*SpellPerk
}

var getSpellSqlTemplate = `
		SELECT
            s.id, s.key, s.name, s.category, s.cost_mana, s.cost_strength,
            s.duration, s.cooldown, s.active, s.created_at, s.updated_at,
            COALESCE(
                (
                 SELECT json_agg(json_build_object(
				 		'id', r.id, 
				 		'key', r.key,
						'name', r.name,
						'alignment', r.alignment,
						'home_land_type', r.home_land_type,
						'description', r.description,
						'playable', r.playable,
						'attacker_difficulty', r.attacker_difficulty,
						'explorer_difficulty', r.explorer_difficulty,
						'converter_difficulty', r.converter_difficulty,
						'overall_difficulty', r.overall_difficulty,
                        'created_at', r.created_at, 
						'updated_at', r.updated_at))
                 FROM races r
                 %s
                ), '[]'::json
            ) AS races_json,
            COALESCE(
                json_agg(
                  json_build_object(
                    'id', sp.id, 
					'value', sp.value,
					'created_at', sp.created_at,
					'updated_at', sp.updated_at,
                    'perk_type', json_build_object(
									'id', spt.id,
									'key', spt.key,
                                    'created_at', spt.created_at,
									'updated_at', spt.updated_at)
                  )
                ) FILTER (WHERE sp.id IS NOT NULL),
                '[]'::json
            ) AS perks_json
        FROM spells s
        LEFT JOIN spell_perks sp ON sp.spell_id = s.id
        LEFT JOIN spell_perk_types spt ON spt.id = sp.spell_perk_type_id
        %s
        GROUP BY s.id`

var defaultGetSpellSqlTemplate = fmt.Sprintf(getSpellSqlTemplate, "WHERE r.key = ANY (array_remove(regexp_split_to_array(COALESCE(s.races, ''), '\\s*,\\s*'), ''))", "%s")

type perkJSON struct {
	Id        int       `json:"id"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PerkType  struct {
		Id        int       `json:"id"`
		Key       string    `json:"key"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"perk_type"`
}

type raceJSON struct {
	Id                  int       `json:"id"`
	Key                 string    `json:"key"`
	Name                string    `json:"name"`
	Alignment           string    `json:"alignment"`
	HomeLandType        string    `json:"home_land_type"`
	Description         string    `json:"description"`
	Playable            bool      `json:"playable"`
	AttackerDifficulty  int       `json:"attacker_difficulty"`
	ExplorerDifficulty  int       `json:"explorer_difficulty"`
	ConverterDifficulty int       `json:"converter_difficulty"`
	OverallDifficulty   int       `json:"overall_difficulty"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func (s *Storage) CreateOrUpdateSpellContext(spell *Spell, ctx context.Context, db DbTx) error {
	races := sql.NullString{Valid: false}
	if len(spell.Races) > 0 {
		raceKeys := make([]string, 0, len(spell.Races))
		for _, r := range spell.Races {
			raceKeys = append(raceKeys, r.Key)
		}
		races = sql.NullString{String: strings.Join(raceKeys, ", "), Valid: true}
	}

	err := db.QueryRowContext(ctx, `
		INSERT INTO spells (key, name, category, cost_mana, cost_strength, duration, cooldown, races, active)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (key) DO UPDATE SET
			name = $2,
			category = $3,
			cost_mana = $4,
			cost_strength = $5,
			duration = $6,
			cooldown = $7,
			races = $8,
			active = $9
			RETURNING id`,
		spell.Key,
		spell.Name,
		spell.Category,
		spell.CostMana,
		spell.CostStrength,
		spell.Duration,
		spell.Cooldown,
		races,
		spell.Active,
	).Scan(&spell.Id)

	if err != nil {
		return fmt.Errorf("unable to create or update spell: %s, error: %w", spell.Key, err)
	}
	return nil
}
func (s *Storage) CreateOrUpdateSpell(spell *Spell) error {
	return s.CreateOrUpdateSpellContext(spell, context.Background(), s.db)
}

func (s *Storage) GetSpellsContext(ctx context.Context, tx DbTx) ([]*Spell, error) {
	sqlQuery := fmt.Sprintf(defaultGetSpellSqlTemplate, "")
	rows, err := tx.QueryContext(ctx, sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("unable to get spells: %w", err)
	}
	defer rows.Close()

	spells := make([]*Spell, 0)
	for rows.Next() {
		spell, err := getSpellFromJsonRow(rows)
		if err != nil {
			return nil, fmt.Errorf("unable to scan spell row: %w", err)
		}
		spells = append(spells, spell)
	}

	return spells, nil
}
func (s *Storage) GetSpells() ([]*Spell, error) {
	return s.GetSpellsContext(context.Background(), s.db)
}

func (s *Storage) GetSpellByKeyContext(key string, ctx context.Context, tx DbTx) (*Spell, error) {
	sqlQuery := fmt.Sprintf(defaultGetSpellSqlTemplate, "WHERE s.key = $1")
	spell, err := getSpellFromJsonRow(tx.QueryRowContext(ctx, sqlQuery, key))
	if err != nil {
		return nil, fmt.Errorf("unable to get spell by key: %s, error: %w", key, err)
	}

	return spell, nil
}
func (s *Storage) GetSpellByKey(key string) (*Spell, error) {
	return s.GetSpellByKeyContext(key, context.Background(), s.db)
}

func (s *Storage) GetSpellByIdContext(id int, ctx context.Context, tx DbTx) (*Spell, error) {
	sqlQuery := fmt.Sprintf(defaultGetSpellSqlTemplate, "WHERE s.id = $1")
	spell, err := getSpellFromJsonRow(tx.QueryRowContext(ctx, sqlQuery, id))
	if err != nil {
		return nil, fmt.Errorf("unable to get spell by id: %d, error: %w", id, err)
	}
	return spell, nil
}
func (s *Storage) GetSpellById(id int) (*Spell, error) {
	return s.GetSpellByIdContext(id, context.Background(), s.db)
}

func (s *Storage) GetSpellsByRaceKeyContext(raceKey string, ctx context.Context, tx DbTx) ([]*Spell, error) {
	sqlQuery := fmt.Sprintf(getSpellSqlTemplate, "WHERE (COALESCE(s.races, '') = '' OR $1 = ANY(array_remove(regexp_split_to_array(COALESCE(s.races, ''), '\\s*,\\s*'), '')))", "")
	rows, err := tx.QueryContext(ctx, sqlQuery, raceKey)
	if err != nil {
		return nil, fmt.Errorf("unable to get spells: %w", err)
	}
	defer rows.Close()

	spells := make([]*Spell, 0)
	for rows.Next() {
		spell, err := getSpellFromJsonRow(rows)
		if err != nil {
			return nil, fmt.Errorf("unable to scan spell row: %w", err)
		}
		spells = append(spells, spell)
	}

	return spells, nil
}
func (s *Storage) GetSpellsByRaceKey(raceKey string) ([]*Spell, error) {
	return s.GetSpellsByRaceKeyContext(raceKey, context.Background(), s.db)
}

func getSpellFromJsonRow(scanner RowScanner) (*Spell, error) {
	var spell Spell
	var racesJSON, perksJSON []byte
	err := scanner.Scan(&spell.Id, &spell.Key, &spell.Name, &spell.Category, &spell.CostMana, &spell.CostStrength,
		&spell.Duration, &spell.Cooldown, &spell.Active, &spell.CreatedAt, &spell.UpdatedAt,
		&racesJSON,
		&perksJSON,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to get spell from json row: %w", err)
	}

	var races []raceJSON
	if err := json.Unmarshal(racesJSON, &races); err != nil {
		return nil, fmt.Errorf("unable to unmarshal races: %w", err)
	}
	spell.Races = make([]*Race, 0, len(races))
	for _, r := range races {
		spell.Races = append(spell.Races, &Race{
			Id:                  r.Id,
			Key:                 r.Key,
			Name:                r.Name,
			Alignment:           r.Alignment,
			HomeLandType:        r.HomeLandType,
			Description:         r.Description,
			Playable:            helpers.NewBoolDefaultTrue(r.Playable),
			AttackerDifficulty:  r.AttackerDifficulty,
			ExplorerDifficulty:  r.ExplorerDifficulty,
			ConverterDifficulty: r.ConverterDifficulty,
			OverallDifficulty:   r.OverallDifficulty,
			CreatedAt:           r.CreatedAt,
			UpdatedAt:           r.UpdatedAt,
		})
	}

	var perks []perkJSON
	if err := json.Unmarshal(perksJSON, &perks); err != nil {
		return nil, fmt.Errorf("unable to unmarshal perks: %w", err)
	}

	spell.Perks = make([]*SpellPerk, 0, len(perks))
	for _, p := range perks {
		spell.Perks = append(spell.Perks, &SpellPerk{
			Id:    p.Id,
			Spell: &spell,
			SpellPerkType: &SpellPerkType{
				Id:        p.PerkType.Id,
				Key:       p.PerkType.Key,
				CreatedAt: p.PerkType.CreatedAt,
				UpdatedAt: p.PerkType.UpdatedAt,
			},
			Value:     p.Value,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	return &spell, nil
}

type SpellPerk struct {
	Id            int
	Spell         *Spell `json:"-" yaml:"-"`
	SpellPerkType *SpellPerkType
	Value         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (s *Storage) CreateOrUpdateSpellPerkContext(sp *SpellPerk, ctx context.Context, db DbTx) error {
	err := db.QueryRowContext(ctx, `
		INSERT INTO spell_perks (spell_id, spell_perk_type_id, value)
		VALUES ($1, $2, $3)
		ON CONFLICT (spell_id, spell_perk_type_id) DO UPDATE SET
			value = $3
		RETURNING id`,
		sp.Spell.Id,
		sp.SpellPerkType.Id,
		sp.Value,
	).Scan(&sp.Id)

	if err != nil {
		return fmt.Errorf("unable to create or update spell perk: spell_id %d, spell_perk_type_id %d, error: %w", sp.Spell.Id, sp.SpellPerkType.Id, err)
	}
	return nil
}
func (s *Storage) CreateOrUpdateSpellPerk(sp *SpellPerk) error {
	return s.CreateOrUpdateSpellPerkContext(sp, context.Background(), s.db)
}

type SpellPerkType struct {
	Id        int
	Key       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Storage) CreateOrUpdateSpellPerkTypeContext(spt *SpellPerkType, ctx context.Context, db DbTx) error {
	err := db.QueryRowContext(ctx, `
		INSERT INTO spell_perk_types (key) VALUES ($1)
		ON CONFLICT (key) DO UPDATE SET key = $1
		RETURNING id`,
		spt.Key,
	).Scan(&spt.Id)

	if err != nil {
		return fmt.Errorf("unable to create or update spell perk type: %s, error: %w", spt.Key, err)
	}
	return nil
}
func (s *Storage) CreateOrUpdateSpellPerkType(spt *SpellPerkType) error {
	return s.CreateOrUpdateSpellPerkTypeContext(spt, context.Background(), s.db)
}
