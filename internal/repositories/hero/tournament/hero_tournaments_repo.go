package tournament

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_hero_tournament.sql
var insertHeroTournamentSQL string

//go:embed sql/get_hero_tournament_by_id.sql
var getHeroTournamentByIDSQL string

//go:embed sql/list_hero_tournaments_by_round.sql
var listHeroTournamentsByRoundSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewHeroTournamentRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	RoundID   *int
	Name      string
	StartDate *string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertHeroTournamentSQL, a.RoundID, a.Name, a.StartDate).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert hero_tournament: %w", err)
	}
	return id, nil
}

type Row struct {
	ID                 int
	RoundID            sql.NullInt64
	Name               string
	CurrentRoundNumber int
	Finished           bool
	WinnerDominionID   sql.NullInt64
	StartDate          sql.NullTime
}

func (r *Repo) GetByIDContext(ctx context.Context, tx repositories.DbTx, id int) (*Row, error) {
	var t Row
	if err := tx.QueryRowContext(ctx, getHeroTournamentByIDSQL, id).
		Scan(&t.ID, &t.RoundID, &t.Name, &t.CurrentRoundNumber, &t.Finished, &t.WinnerDominionID, &t.StartDate); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get hero_tournament: %w", err)
	}
	return &t, nil
}

func (r *Repo) ListByRoundContext(ctx context.Context, tx repositories.DbTx, roundID int, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listHeroTournamentsByRoundSQL, roundID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list hero_tournaments: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var t Row
		if err := rows.Scan(&t.ID, &t.RoundID, &t.Name, &t.CurrentRoundNumber, &t.Finished, &t.WinnerDominionID, &t.StartDate); err != nil {
			return nil, fmt.Errorf("scan hero_tournament: %w", err)
		}
		out = append(out, t)
	}
	return out, rows.Err()
}
