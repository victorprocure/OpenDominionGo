package entries

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_telescope_entry.sql
var insertTelescopeEntrySQL string

//go:embed sql/get_telescope_entry_by_sequence.sql
var getTelescopeEntryBySequenceSQL string

//go:embed sql/list_telescope_entries_by_type.sql
var listTelescopeEntriesByTypeSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	UUID                 uuid.UUID
	BatchID              uuid.UUID
	FamilyHash           *string
	ShouldDisplayOnIndex bool
	Type                 string
	Content              string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int64, error) {
	var sequence int64
	if err := tx.QueryRowContext(ctx, insertTelescopeEntrySQL, a.UUID, a.BatchID, a.FamilyHash, a.ShouldDisplayOnIndex, a.Type, a.Content).Scan(&sequence); err != nil {
		return 0, fmt.Errorf("insert telescope_entry: %w", err)
	}
	return sequence, nil
}

type Row struct {
	Sequence             int64
	UUID                 uuid.UUID
	BatchID              uuid.UUID
	FamilyHash           sql.NullString
	ShouldDisplayOnIndex bool
	Type                 string
	Content              string
}

func (r *Repo) GetBySequenceContext(ctx context.Context, tx repositories.DbTx, sequence int64) (*Row, error) {
	var x Row
	if err := tx.QueryRowContext(ctx, getTelescopeEntryBySequenceSQL, sequence).
		Scan(&x.Sequence, &x.UUID, &x.BatchID, &x.FamilyHash, &x.ShouldDisplayOnIndex, &x.Type, &x.Content); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get telescope_entry: %w", err)
	}
	return &x, nil
}

func (r *Repo) ListByTypeContext(ctx context.Context, tx repositories.DbTx, typ string, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listTelescopeEntriesByTypeSQL, typ, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list telescope_entries: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var x Row
		if err := rows.Scan(&x.Sequence, &x.UUID, &x.BatchID, &x.FamilyHash, &x.ShouldDisplayOnIndex, &x.Type, &x.Content); err != nil {
			return nil, fmt.Errorf("scan telescope_entry: %w", err)
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
