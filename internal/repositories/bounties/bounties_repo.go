package bounty

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_bounty.sql
var insertBountySQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewBountiesRepo(db *sql.DB, log *slog.Logger) *Repo {
	return &Repo{db: db, log: log}
}

type InsertArgs struct {
	RoundID             int
	SourceRealmID       int
	SourceDominionID    int
	TargetDominionID    int
	CollectedByDominion *int // nullable
	Type                string
	Reward              bool
}

func (r *Repo) InsertContext(ctx context.Context, tx repositories.DbTx, a InsertArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertBountySQL,
		a.RoundID,
		a.SourceRealmID,
		a.SourceDominionID,
		a.TargetDominionID,
		a.CollectedByDominion,
		a.Type,
		a.Reward,
	).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert bounty: %w", err)
	}
	return id, nil
}
