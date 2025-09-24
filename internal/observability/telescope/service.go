package telescope

import (
	"database/sql"
	"encoding/json"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/victorprocure/opendominiongo/internal/config"
	telentry "github.com/victorprocure/opendominiongo/internal/repositories/telescope/entry"
	teltag "github.com/victorprocure/opendominiongo/internal/repositories/telescope/entry/tag"
)

type service struct {
	db      *sql.DB
	entries *telentry.Repo
	tags    *teltag.Repo
	log     *slog.Logger
}

func NewService(db *sql.DB, cfg *config.AppConfig) Service {
	return &service{
		db:      db,
		entries: telentry.NewRepo(db, cfg.Log),
		tags:    teltag.NewRepo(db, cfg.Log),
		log:     cfg.Log,
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
func (s *service) Capture(ctx *gin.Context, typ string, content any, opts ...Option) (int64, error) {
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
	sequence, err := s.entries.CreateContext(ctx, s.db, telentry.CreateArgs{
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
		_ = s.tags.CreateContext(ctx, s.db, teltag.CreateArgs{EntryUUID: entryUUID, Tag: tag})
	}
	return sequence, nil
}
