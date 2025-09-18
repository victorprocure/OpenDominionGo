// Deprecated: use internal/repositories/hero/tournament/battle instead (package battle).
package hero_tournament_battles

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_hero_tournament_battle.sql
var insertTournamentBattleSQL string

//go:embed sql/list_hero_tournament_battles_by_tournament.sql
var listTournamentBattlesByTournamentSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewHeroTournamentBattlesRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	HeroTournamentID int
	HeroBattleID     int
	RoundNumber      int
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertTournamentBattleSQL, a.HeroTournamentID, a.HeroBattleID, a.RoundNumber).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert hero_tournament_battle: %w", err)
	}
	return id, nil
}

type Row struct {
	ID               int
	HeroTournamentID int
	HeroBattleID     int
	RoundNumber      int
}

func (r *Repo) ListByTournamentContext(ctx context.Context, tx repositories.DbTx, tournamentID int, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listTournamentBattlesByTournamentSQL, tournamentID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list hero_tournament_battles: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var b Row
		if err := rows.Scan(&b.ID, &b.HeroTournamentID, &b.HeroBattleID, &b.RoundNumber); err != nil {
			return nil, fmt.Errorf("scan hero_tournament_battle: %w", err)
		}
		out = append(out, b)
	}
	return out, rows.Err()
}
