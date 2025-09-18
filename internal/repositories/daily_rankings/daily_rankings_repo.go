package daily_rankings

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/upsert_daily_ranking.sql
var upsertDailyRankingSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewDailyRankingsRepo(db *sql.DB, log *slog.Logger) *Repo {
	return &Repo{db: db, log: log}
}

// UpsertArgs captures the fields needed to upsert a daily ranking row.
// The database enforces a UNIQUE (dominion_id, key), which we use as the conflict target.
type UpsertArgs struct {
	RoundID      int
	DominionID   int
	DominionName string
	RaceName     string
	RealmNumber  int
	RealmName    string
	Key          string
	Value        int
	Rank         *int // nullable
	PreviousRank *int // nullable
}

// UpsertFromSyncContext upserts a daily ranking row and returns the row id.
func (r *Repo) UpsertFromSyncContext(ctx context.Context, tx repositories.DbTx, a UpsertArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, upsertDailyRankingSQL,
		a.RoundID,
		a.DominionID,
		a.DominionName,
		a.RaceName,
		a.RealmNumber,
		a.RealmName,
		a.Key,
		a.Value,
		a.Rank,
		a.PreviousRank,
	).Scan(&id); err != nil {
		return 0, fmt.Errorf("upsert daily_ranking (%d,%s): %w", a.DominionID, a.Key, err)
	}
	return id, nil
}
