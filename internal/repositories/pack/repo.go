package pack

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_pack.sql
var insertPackSQL string

//go:embed sql/get_pack_by_round_and_name.sql
var getPackByRoundAndNameSQL string

//go:embed sql/update_pack.sql
var updatePackSQL string

//go:embed sql/delete_pack.sql
var deletePackSQL string

//go:embed sql/list_packs_by_round.sql
var listPacksByRoundSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	RoundID  int
	RealmID  *int
	Name     string
	Password string
	Size     int
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertPackSQL, a.RoundID, a.RealmID, a.Name, a.Password, a.Size).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert pack: %w", err)
	}
	return id, nil
}

type Row struct {
	ID      int
	RoundID int
	RealmID *int
	Name    string
	Size    int
}

func (r *Repo) GetByRoundAndNameContext(ctx context.Context, tx repositories.DbTx, roundID, name string) (*Row, error) {
	var p Row
	if err := tx.QueryRowContext(ctx, getPackByRoundAndNameSQL, roundID, name).
		Scan(&p.ID, &p.RoundID, &p.RealmID, &p.Name, &p.Size); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get pack: %w", err)
	}
	return &p, nil
}

type UpdateArgs struct {
	ID       int
	RealmID  *int
	Name     string
	Password string
	Size     int
}

func (r *Repo) UpdateContext(ctx context.Context, tx repositories.DbTx, a UpdateArgs) error {
	if _, err := tx.ExecContext(ctx, updatePackSQL, a.ID, a.RealmID, a.Name, a.Password, a.Size); err != nil {
		return fmt.Errorf("update pack: %w", err)
	}
	return nil
}

func (r *Repo) DeleteByIDContext(ctx context.Context, tx repositories.DbTx, id int) error {
	if _, err := tx.ExecContext(ctx, deletePackSQL, id); err != nil {
		return fmt.Errorf("delete pack: %w", err)
	}
	return nil
}

func (r *Repo) ListByRoundContext(ctx context.Context, tx repositories.DbTx, roundID, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listPacksByRoundSQL, roundID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list packs: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var p Row
		if err := rows.Scan(&p.ID, &p.RoundID, &p.RealmID, &p.Name, &p.Size); err != nil {
			return nil, fmt.Errorf("scan pack: %w", err)
		}
		out = append(out, p)
	}
	return out, rows.Err()
}
