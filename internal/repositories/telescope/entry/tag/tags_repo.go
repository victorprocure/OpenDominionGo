package tag

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_telescope_entry_tag.sql
var insertTelescopeEntryTagSQL string

//go:embed sql/list_telescope_entry_tags_by_entry.sql
var listTelescopeEntryTagsByEntrySQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	EntryUUID uuid.UUID
	Tag       string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) error {
	if _, err := tx.ExecContext(ctx, insertTelescopeEntryTagSQL, a.EntryUUID, a.Tag); err != nil {
		return fmt.Errorf("insert telescope_entry_tag: %w", err)
	}
	return nil
}

type Row struct {
	EntryUUID uuid.UUID
	Tag       string
}

func (r *Repo) ListByEntryContext(ctx context.Context, tx repositories.DbTx, entryUUID uuid.UUID, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listTelescopeEntryTagsByEntrySQL, entryUUID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list telescope_entry_tags: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var x Row
		if err := rows.Scan(&x.EntryUUID, &x.Tag); err != nil {
			return nil, fmt.Errorf("scan telescope_entry_tag: %w", err)
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
