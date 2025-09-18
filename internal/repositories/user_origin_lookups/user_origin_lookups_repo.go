package user_origin_lookups

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_user_origin_lookup.sql
var insertUserOriginLookupSQL string

//go:embed sql/get_user_origin_lookup_by_ip.sql
var getUserOriginLookupByIPSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewUserOriginLookupsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	IPAddress    string
	ISP          *string
	Organization *string
	Country      *string
	Region       *string
	City         *string
	VPN          *bool
	Score        *float64
	Data         *string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) error {
	if _, err := tx.ExecContext(ctx, insertUserOriginLookupSQL, a.IPAddress, a.ISP, a.Organization, a.Country, a.Region, a.City, a.VPN, a.Score, a.Data); err != nil {
		return fmt.Errorf("insert user_origin_lookup: %w", err)
	}
	return nil
}

type Row struct {
	ID           int
	IPAddress    string
	ISP          sql.NullString
	Organization sql.NullString
	Country      sql.NullString
	Region       sql.NullString
	City         sql.NullString
	VPN          sql.NullBool
	Score        sql.NullFloat64
	Data         sql.NullString
}

func (r *Repo) GetByIPContext(ctx context.Context, tx repositories.DbTx, ip string) (*Row, error) {
	var x Row
	if err := tx.QueryRowContext(ctx, getUserOriginLookupByIPSQL, ip).
		Scan(&x.ID, &x.IPAddress, &x.ISP, &x.Organization, &x.Country, &x.Region, &x.City, &x.VPN, &x.Score, &x.Data); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get user_origin_lookup: %w", err)
	}
	return &x, nil
}
