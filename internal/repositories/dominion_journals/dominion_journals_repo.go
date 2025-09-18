package dominion_journals

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_dominion_journal.sql
var insertDominionJournalSQL string

//go:embed sql/list_dominion_journals_by_dominion.sql
var listDominionJournalsByDominionSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewDominionJournalsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	DominionID int
	Content    string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertDominionJournalSQL, a.DominionID, a.Content).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert dominion_journal: %w", err)
	}
	return id, nil
}

type Row struct {
	ID         int
	DominionID int
	Content    string
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
}

func (r *Repo) ListByDominionContext(ctx context.Context, tx repositories.DbTx, dominionID int, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listDominionJournalsByDominionSQL, dominionID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list dominion_journals: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var j Row
		if err := rows.Scan(&j.ID, &j.DominionID, &j.Content, &j.CreatedAt, &j.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan dominion_journal: %w", err)
		}
		out = append(out, j)
	}
	return out, rows.Err()
}
