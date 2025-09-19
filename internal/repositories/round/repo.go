package round

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"
	"time"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

//go:embed sql/get_round_by_id.sql
var getRoundByIDSQL string

//go:embed sql/get_latest_round.sql
var getLatestRoundSQL string

//go:embed sql/list_active_rounds.sql
var listActiveRoundsSQL string

type Row struct {
	ID                           int
	RoundLeagueID                int
	Number                       int
	Name                         string
	StartDate                    time.Time
	EndDate                      time.Time
	CreatedAt                    sql.NullTime
	UpdatedAt                    sql.NullTime
	RealmSize                    int
	PackSize                     int
	PlayersPerRace               int
	MixedAlignment               bool
	OffensiveActionsProhibitedAt sql.NullTime
	DiscordGuildID               sql.NullString
	TechVersion                  int
	LargestHit                   int
	AssignmentComplete           bool
}

func scanRow(s scanner) (Row, error) {
	var x Row
	err := s.Scan(
		&x.ID,
		&x.RoundLeagueID,
		&x.Number,
		&x.Name,
		&x.StartDate,
		&x.EndDate,
		&x.CreatedAt,
		&x.UpdatedAt,
		&x.RealmSize,
		&x.PackSize,
		&x.PlayersPerRace,
		&x.MixedAlignment,
		&x.OffensiveActionsProhibitedAt,
		&x.DiscordGuildID,
		&x.TechVersion,
		&x.LargestHit,
		&x.AssignmentComplete,
	)
	if err != nil {
		return Row{}, err
	}
	return x, nil
}

type scanner interface{ Scan(dest ...any) error }

func (r *Repo) GetByIDContext(ctx context.Context, tx repositories.DbTx, id int) (*Row, error) {
	var x Row
	row := tx.QueryRowContext(ctx, getRoundByIDSQL, id)
	xr, err := scanRow(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get round by id: %w", err)
	}
	x = xr
	return &x, nil
}

func (r *Repo) GetLatestContext(ctx context.Context, tx repositories.DbTx) (*Row, error) {
	row := tx.QueryRowContext(ctx, getLatestRoundSQL)
	xr, err := scanRow(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get latest round: %w", err)
	}
	return &xr, nil
}

func (r *Repo) ListActiveContext(ctx context.Context, tx repositories.DbTx, now time.Time) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listActiveRoundsSQL, now)
	if err != nil {
		return nil, fmt.Errorf("list active rounds: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		xr, err := scanRow(rows)
		if err != nil {
			return nil, fmt.Errorf("scan round: %w", err)
		}
		out = append(out, xr)
	}
	return out, rows.Err()
}
