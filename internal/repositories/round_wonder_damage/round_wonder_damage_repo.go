// deprecated: moved to internal/repositories/round/wonder/damage
package round_wonder_damage

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_round_wonder_damage.sql
var insertRoundWonderDamageSQL string

//go:embed sql/list_round_wonder_damage_by_round_wonder.sql
var listByRoundWonderSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRoundWonderDamageRepo(db *sql.DB, log *slog.Logger) *Repo {
	return &Repo{db: db, log: log}
}

type InsertArgs struct {
	RoundWonderID int
	RealmID       int
	DominionID    int
	Damage        int
	Source        *string
}

func (r *Repo) InsertContext(ctx context.Context, tx repositories.DbTx, a InsertArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertRoundWonderDamageSQL,
		a.RoundWonderID, a.RealmID, a.DominionID, a.Damage, a.Source,
	).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert round_wonder_damage: %w", err)
	}
	return id, nil
}

type Row struct {
	ID            int
	RoundWonderID int
	RealmID       int
	DominionID    int
	Damage        int
	Source        *string
}

func (r *Repo) ListByRoundWonderContext(ctx context.Context, tx repositories.DbTx, roundWonderID int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listByRoundWonderSQL, roundWonderID)
	if err != nil {
		return nil, fmt.Errorf("list round_wonder_damage: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var rw Row
		if err := rows.Scan(&rw.ID, &rw.RoundWonderID, &rw.RealmID, &rw.DominionID, &rw.Damage, &rw.Source); err != nil {
			return nil, fmt.Errorf("scan round_wonder_damage: %w", err)
		}
		out = append(out, rw)
	}
	return out, rows.Err()
}
