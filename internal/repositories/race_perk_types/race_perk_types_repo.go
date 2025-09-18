// Deprecated: use internal/repositories/race/perk_type instead (package perk_type).
package race_perk_types

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_race_perk_type.sql
var insertRacePerkTypeSQL string

//go:embed sql/get_race_perk_type_by_key.sql
var getRacePerkTypeByKeySQL string

//go:embed sql/list_race_perk_types.sql
var listRacePerkTypesSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRacePerkTypesRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct{ Key string }

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, key string) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertRacePerkTypeSQL, key).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert race_perk_type: %w", err)
	}
	return id, nil
}

type Row struct {
	ID  int
	Key string
}

func (r *Repo) GetByKeyContext(ctx context.Context, tx repositories.DbTx, key string) (*Row, error) {
	var x Row
	if err := tx.QueryRowContext(ctx, getRacePerkTypeByKeySQL, key).Scan(&x.ID, &x.Key); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get race_perk_type: %w", err)
	}
	return &x, nil
}

func (r *Repo) ListContext(ctx context.Context, tx repositories.DbTx, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listRacePerkTypesSQL, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list race_perk_types: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var x Row
		if err := rows.Scan(&x.ID, &x.Key); err != nil {
			return nil, fmt.Errorf("scan race_perk_type: %w", err)
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
