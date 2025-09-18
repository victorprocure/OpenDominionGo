// Deprecated: use internal/repositories/race/perk instead (package perk).
package race_perks

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_race_perk.sql
var insertRacePerkSQL string

//go:embed sql/list_race_perks_by_race.sql
var listRacePerksByRaceSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRacePerksRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	RaceID         int
	RacePerkTypeID int
	Value          float64
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertRacePerkSQL, a.RaceID, a.RacePerkTypeID, a.Value).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert race_perk: %w", err)
	}
	return id, nil
}

type Row struct {
	ID             int
	RaceID         int
	RacePerkTypeID int
	Value          float64
}

func (r *Repo) ListByRaceContext(ctx context.Context, tx repositories.DbTx, raceID int, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listRacePerksByRaceSQL, raceID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list race_perks: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var x Row
		if err := rows.Scan(&x.ID, &x.RaceID, &x.RacePerkTypeID, &x.Value); err != nil {
			return nil, fmt.Errorf("scan race_perk: %w", err)
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
