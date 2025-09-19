package valor

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_valor.sql
var insertValorSQL string

//go:embed sql/list_valor_by_dominion.sql
var listValorByDominionSQL string

//go:embed sql/sum_valor_by_dominion.sql
var sumValorByDominionSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type InsertArgs struct {
	RoundID    int
	RealmID    int
	DominionID int
	Source     string
	Amount     float64
}

func (r *Repo) InsertContext(ctx context.Context, tx repositories.DbTx, a InsertArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertValorSQL, a.RoundID, a.RealmID, a.DominionID, a.Source, a.Amount).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert valor: %w", err)
	}
	return id, nil
}

type Row struct {
	ID         int
	RoundID    int
	RealmID    int
	DominionID int
	Source     string
	Amount     float64
}

func (r *Repo) ListByDominionContext(ctx context.Context, tx repositories.DbTx, dominionID, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listValorByDominionSQL, dominionID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list valor: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var v Row
		if err := rows.Scan(&v.ID, &v.RoundID, &v.RealmID, &v.DominionID, &v.Source, &v.Amount); err != nil {
			return nil, fmt.Errorf("scan valor: %w", err)
		}
		out = append(out, v)
	}
	return out, rows.Err()
}

func (r *Repo) SumByDominionContext(ctx context.Context, tx repositories.DbTx, dominionID int) (float64, error) {
	var total float64
	if err := tx.QueryRowContext(ctx, sumValorByDominionSQL, dominionID).Scan(&total); err != nil {
		return 0, fmt.Errorf("sum valor: %w", err)
	}
	return total, nil
}
