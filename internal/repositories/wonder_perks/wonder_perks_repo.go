// Deprecated: use internal/repositories/wonder/perk instead.
package wonder_perks

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_wonder_perk.sql
var insertWonderPerkSQL string

//go:embed sql/list_wonder_perks_by_wonder.sql
var listWonderPerksByWonderSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewWonderPerksRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	WonderID         int
	WonderPerkTypeID int
	Value            string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertWonderPerkSQL, a.WonderID, a.WonderPerkTypeID, a.Value).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert wonder_perk: %w", err)
	}
	return id, nil
}

type Row struct {
	ID               int
	WonderID         int
	WonderPerkTypeID int
	Value            string
}

func (r *Repo) ListByWonderContext(ctx context.Context, tx repositories.DbTx, wonderID int, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listWonderPerksByWonderSQL, wonderID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list wonder_perks: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var x Row
		if err := rows.Scan(&x.ID, &x.WonderID, &x.WonderPerkTypeID, &x.Value); err != nil {
			return nil, fmt.Errorf("scan wonder_perk: %w", err)
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
