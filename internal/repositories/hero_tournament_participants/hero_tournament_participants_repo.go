// Deprecated: use internal/repositories/hero/tournament/participant instead (package participant).
package hero_tournament_participants

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_hero_tournament_participant.sql
var insertParticipantSQL string

//go:embed sql/list_hero_tournament_participants_by_tournament.sql
var listParticipantsByTournamentSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewHeroTournamentParticipantsRepo(db *sql.DB, log *slog.Logger) *Repo {
	return &Repo{db: db, log: log}
}

type CreateArgs struct {
	HeroTournamentID int
	HeroID           int
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertParticipantSQL, a.HeroTournamentID, a.HeroID).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert hero_tournament_participant: %w", err)
	}
	return id, nil
}

type Row struct {
	ID               int
	HeroTournamentID int
	HeroID           int
	Wins             int
	Losses           int
	Draws            int
	Standing         sql.NullInt64
	Eliminated       bool
}

func (r *Repo) ListByTournamentContext(ctx context.Context, tx repositories.DbTx, tournamentID int, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listParticipantsByTournamentSQL, tournamentID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list hero_tournament_participants: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var p Row
		if err := rows.Scan(&p.ID, &p.HeroTournamentID, &p.HeroID, &p.Wins, &p.Losses, &p.Draws, &p.Standing, &p.Eliminated); err != nil {
			return nil, fmt.Errorf("scan hero_tournament_participant: %w", err)
		}
		out = append(out, p)
	}
	return out, rows.Err()
}
