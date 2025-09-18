package datasync

import (
	"context"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

type Syncer interface {
	Name() string
	PerformDataSync(ctx context.Context, tx repositories.DbTx) error
}