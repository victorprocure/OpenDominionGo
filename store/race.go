package store

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log/slog"
	"time"

	"github.com/victorprocure/opendominiongo/helpers"
)

//go:embed data/races
var racesFS embed.FS

type Race struct {
	Id                  int
	Name                string `yaml:"name"`
	Alignment           string `yaml:"alignment"`
	HomeLandType        string `yaml:"home_land_type"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Description         string                  `yaml:"description"`
	Playable            helpers.BoolDefaultTrue `yaml:"playable"`
	AttackerDifficulty  int                     `yaml:"attacker_difficulty"`
	ExplorerDifficulty  int                     `yaml:"explorer_difficulty"`
	ConverterDifficulty int                     `yaml:"converter_difficulty,omitempty"`
	OverallDifficulty   int                     `yaml:"overall_difficulty,omitempty"`
	Key                 string                  `yaml:"key"`
	Perks               []*RacePerk             `yaml:"-"`
	Units               []*Unit                 `yaml:"-"`
}

func (r Race) String() string {
	return fmt.Sprintf(
		"Name: %s\nAlignment: %s\nHomeLandType: %s\nPlayable: %t\nAttackerDifficulty: %d\nExplorerDifficulty: %d\nConverterDifficulty: %d\nOverallDifficulty: %d\nKey: %s\nDescription: %s\n",
		r.Name, r.Alignment, r.HomeLandType, r.Playable.OrDefault(),
		r.AttackerDifficulty, r.ExplorerDifficulty, r.ConverterDifficulty, r.OverallDifficulty,
		r.Key, r.Description,
	)
}

func (s *Storage) CreateOrUpdateRaceContext(r *Race, ctx context.Context, db DbTx) error {
	err := db.QueryRowContext(ctx, `INSERT INTO Races
		(name, alignment, home_land_type, description, playable, attacker_difficulty,
		 explorer_difficulty, converter_difficulty, overall_difficulty, key) 
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		 ON CONFLICT (key) DO UPDATE
		   SET name = $1,
		       alignment = $2,
		       home_land_type = $3,
			   description = $4,
			   playable = $5,
			   attacker_difficulty = $6,
			   explorer_difficulty = $7,
			   converter_difficulty = $8,
			   overall_difficulty = $9,
			   key = $10
		 RETURNING id`,
		r.Name,
		r.Alignment,
		r.HomeLandType,
		r.Description,
		r.Playable.OrDefault(),
		r.AttackerDifficulty,
		r.ExplorerDifficulty,
		r.ConverterDifficulty,
		r.OverallDifficulty,
		r.Key,
	).Scan(&r.Id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) CreateOrUpdateRace(r *Race) error {
	return s.CreateOrUpdateRaceContext(r, context.Background(), s.db)
}

func (s *Storage) GetRaceByKeyContext(key string, ctx context.Context, db DbTx) (*Race, error) {
	var r Race
	err := db.QueryRowContext(ctx, "SELECT id, name, alignment, home_land_type, description, playable, attacker_difficulty, explorer_difficulty, converter_difficulty, overall_difficulty, key, created_at, updated_at FROM Races WHERE key = $1", key).Scan(
		&r.Id,
		&r.Name,
		&r.Alignment,
		&r.HomeLandType,
		&r.Description,
		&r.Playable,
		&r.AttackerDifficulty,
		&r.ExplorerDifficulty,
		&r.ConverterDifficulty,
		&r.OverallDifficulty,
		&r.Key,
		&r.CreatedAt,
		&r.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		s.Log.Error("Unable to query for race by key", slog.String("key", key), slog.Any("error", err))
		return nil, err
	}

	perks, err := s.GetRacePerksByRaceId(&r, ctx, db)
	if err != nil {
		return nil, err
	}
	r.Perks = perks

	return &r, nil
}

func (s *Storage) GetRaceByKey(key string) (*Race, error) {
	return s.GetRaceByKeyContext(key, context.Background(), s.db)
}

type RacePerk struct {
	Id           int
	Race         *Race
	RacePerkType *RacePerkType
	Value        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (s *Storage) CreateOrUpdateRacePerkContext(rp *RacePerk, ctx context.Context, db DbTx) error {
	err := db.QueryRowContext(ctx, `INSERT INTO race_perks 
		(race_id, race_perk_type_id, value) 
		VALUES ($1, $2, $3)
		ON CONFLICT (race_id, race_perk_type_id) DO UPDATE SET
		  value = $3
		RETURNING id`,
		rp.Race.Id,
		rp.RacePerkType.Id,
		rp.Value).Scan(&rp.Id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) CreateOrUpdateRacePerk(rp *RacePerk) error {
	return s.CreateOrUpdateRacePerkContext(rp, context.Background(), s.db)
}

func (s *Storage) GetRacePerksByRaceId(r *Race, ctx context.Context, tx DbTx) ([]*RacePerk, error) {
	rows, err := tx.QueryContext(ctx, `
	SELECT
		rp.id,
		rp.value,
		rp.created_at,
		rp.updated_at,
		rpt.id AS rpt_id,
		rpt.key AS rpt_key,
		rpt.created_at as rpt_created_at,
		rpt.updated_at as rpt_updated_at
	FROM race_perks rp
	JOIN race_perk_types rpt ON rp.race_perk_type_id = rpt.id
	WHERE rp.race_id = $1`, r.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return []*RacePerk{}, nil
		}

		return nil, err
	}
	defer rows.Close()

	perks := make([]*RacePerk, 0)
	for rows.Next() {
		var (
			rpID       int
			value      float64
			rpCreated  time.Time
			rpUpdated  time.Time
			rptID      int
			rptKey     string
			rptCreated time.Time
			rptUpdated time.Time
		)
		if err := rows.Scan(
			&rpID, &value, &rpCreated, &rpUpdated,
			&rptID, &rptKey, &rptCreated, &rptUpdated,
		); err != nil {
			return nil, err
		}

		perks = append(perks, &RacePerk{
			Id:        rpID,
			Race:      r,
			Value:     value,
			CreatedAt: rpCreated,
			UpdatedAt: rpUpdated,
			RacePerkType: &RacePerkType{
				Id:        rptID,
				Key:       rptKey,
				CreatedAt: rptCreated,
				UpdatedAt: rptUpdated,
			},
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return perks, nil
}

func (s *Storage) GetRacePerksByRace(r *Race) ([]*RacePerk, error) {
	return s.GetRacePerksByRaceId(r, context.Background(), s.db)
}

type RacePerkType struct {
	Id        int
	Key       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Storage) CreateOrUpdateRacePerkTypeContext(rpt *RacePerkType, ctx context.Context, db DbTx) error {
	err := db.QueryRowContext(ctx, `INSERT INTO race_perk_types
		(key) 
		VALUES ($1)
		ON CONFLICT (key) DO UPDATE SET
		  key = $1
		RETURNING id`,
		rpt.Key,
	).Scan(&rpt.Id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) CreateOrUpdateRacePerkType(rpt *RacePerkType) error {
	return s.CreateOrUpdateRacePerkTypeContext(rpt, context.Background(), s.db)
}

func (s *Storage) GetRacePerkTypeByKeyContext(key string, ctx context.Context, db DbTx) (*RacePerkType, error) {
	var rpt RacePerkType
	err := db.QueryRowContext(ctx,
		"SELECT id, key, created_at, updated_at FROM race_perk_types WHERE key = $1",
		key,
	).Scan(
		&rpt.Id,
		&rpt.Key,
		&rpt.CreatedAt,
		&rpt.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &rpt, nil
}
