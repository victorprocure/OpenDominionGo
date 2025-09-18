package telescope

import (
	"context"
	"database/sql"
	"encoding/json"
	"log/slog"

	"github.com/google/uuid"
	telentries "github.com/victorprocure/opendominiongo/internal/repositories/telescope/entries"
	teltags "github.com/victorprocure/opendominiongo/internal/repositories/telescope/entries/tags"
)

type Service struct {
	db      *sql.DB
	entries *telentries.Repo
	tags    *teltags.Repo
	log     *slog.Logger
}

func NewService(db *sql.DB, log *slog.Logger) *Service {
	return &Service{
		db:      db,
		entries: telentries.NewRepo(db, log),
		tags:    teltags.NewRepo(db, log),
		log:     log,
	}
}

type CaptureOptions struct {
	BatchID              uuid.UUID
	FamilyHash           *string
	ShouldDisplayOnIndex bool
	Tags                 []string
}

type Option func(*CaptureOptions)

func WithBatchID(id uuid.UUID) Option  { return func(o *CaptureOptions) { o.BatchID = id } }
func WithFamilyHash(h string) Option   { return func(o *CaptureOptions) { o.FamilyHash = &h } }
func WithDisplayOnIndex(v bool) Option { return func(o *CaptureOptions) { o.ShouldDisplayOnIndex = v } }
func WithTags(tags ...string) Option {
	return func(o *CaptureOptions) { o.Tags = append(o.Tags, tags...) }
}

// Capture serializes content as JSON and stores a telescope entry with optional tags.
func (s *Service) Capture(ctx context.Context, typ string, content any, opts ...Option) (int64, error) {
	var co CaptureOptions
	for _, fn := range opts {
		fn(&co)
	}
	// Ensure batch id
	if co.BatchID == uuid.Nil {
		co.BatchID = uuid.New()
	}
	// JSON encode content
	b, err := json.Marshal(content)
	if err != nil {
		// fallback minimal content
		b = []byte(`{"error":"marshal failed"}`)
	}

	// Insert entry
	entryUUID := uuid.New()
	sequence, err := s.entries.CreateContext(ctx, s.db, telentries.CreateArgs{
		UUID:                 entryUUID,
		BatchID:              co.BatchID,
		FamilyHash:           co.FamilyHash,
		ShouldDisplayOnIndex: co.ShouldDisplayOnIndex,
		Type:                 typ,
		Content:              string(b),
	})
	if err != nil {
		return 0, err
	}
	// Add tags
	for _, tag := range co.Tags {
		_ = s.tags.CreateContext(ctx, s.db, teltags.CreateArgs{EntryUUID: entryUUID, Tag: tag})
	}
	return sequence, nil
}
