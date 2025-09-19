package infoops

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_info_op.sql
var insertInfoOpSQL string

//go:embed sql/list_info_ops_by_target_dominion.sql
var listInfoOpsByTargetDominionSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewInfoOpsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	SourceRealmID    int
	SourceDominionID int
	TargetDominionID int
	TargetRealmID    *int
	Type             string
	Data             string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertInfoOpSQL, a.SourceRealmID, a.SourceDominionID, a.TargetDominionID, a.Type, a.Data, a.TargetRealmID).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert info_op: %w", err)
	}
	return id, nil
}

type Row struct {
	ID               int
	SourceRealmID    int
	SourceDominionID int
	TargetDominionID int
	TargetRealmID    *int
	Type             string
	Data             string
}

func (r *Repo) ListByTargetDominionContext(ctx context.Context, tx repositories.DbTx, dominionID int, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listInfoOpsByTargetDominionSQL, dominionID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list info_ops: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var io Row
		if err := rows.Scan(&io.ID, &io.SourceRealmID, &io.SourceDominionID, &io.TargetDominionID, &io.TargetRealmID, &io.Type, &io.Data); err != nil {
			return nil, fmt.Errorf("scan info_op: %w", err)
		}
		out = append(out, io)
	}
	return out, rows.Err()
}
