package telescope

import (
	"context"
	"log/slog"
)

// SlogHandler forwards records to an inner handler and mirrors them into telescope entries.
type SlogHandler struct {
	inner slog.Handler
	svc   *Service
	level slog.Leveler
}

func NewSlogHandler(inner slog.Handler, svc *Service, level slog.Leveler) *SlogHandler {
	return &SlogHandler{inner: inner, svc: svc, level: level}
}

func (h *SlogHandler) Enabled(ctx context.Context, l slog.Level) bool {
	return h.inner.Enabled(ctx, l)
}

func (h *SlogHandler) Handle(ctx context.Context, r slog.Record) error {
	// forward first
	if err := h.inner.Handle(ctx, r); err != nil {
		return err
	}
	// mirror into telescope as a log entry; errors show on index
	should := r.Level >= slog.LevelError
	content := map[string]any{"level": r.Level.String(), "message": r.Message}
	// capture attributes
	attrs := map[string]any{}
	r.Attrs(func(a slog.Attr) bool { attrs[a.Key] = a.Value.Any(); return true })
	if len(attrs) > 0 {
		content["attrs"] = attrs
	}
	_, _ = h.svc.Capture(ctx, "log", content, WithDisplayOnIndex(should))
	return nil
}

func (h *SlogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &SlogHandler{inner: h.inner.WithAttrs(attrs), svc: h.svc, level: h.level}
}

func (h *SlogHandler) WithGroup(name string) slog.Handler {
	return &SlogHandler{inner: h.inner.WithGroup(name), svc: h.svc, level: h.level}
}
