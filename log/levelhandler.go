package log

import "golang.org/x/exp/slog"

var _ slog.Handler = (*LevelHandler)(nil)

type LevelHandler struct {
	level   slog.Level
	handler slog.Handler
}

// NewLevelHandler implements slog.Handler interface. It is used to wrap a
// handler with a new log level. As log level cannot be modified within a hanlder
// through the interface of slog, you can use this to wrap a handler with a new
// log level.
func NewLevelHandler(level slog.Level, h slog.Handler) *LevelHandler {
	return &LevelHandler{level, h}
}

// Enabled reports whether the handler handles records at the given level.
// The handler ignores records whose level is lower.
// Enabled is called early, before any arguments are processed,
// to save effort if the log event should be discarded.
func (h *LevelHandler) Enabled(level slog.Level) bool {
	return level >= h.level
}

// Handle handles the Record.
// It will only be called if Enabled returns true.
// Handle methods that produce output should observe the following rules:
//   - If r.Time is the zero time, ignore the time.
//   - If an Attr's key is the empty string, ignore the Attr.
func (h *LevelHandler) Handle(r slog.Record) error {
	return h.handler.Handle(r)
}

// WithAttrs returns a new Handler whose attributes consist of
// both the receiver's attributes and the arguments.
// The Handler owns the slice: it may retain, modify or discard it.
func (h *LevelHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewLevelHandler(h.level, h.handler.WithAttrs(attrs))
}

// WithGroup returns a new Handler with the given group appended to
// the receiver's existing groups.
// The keys of all subsequent attributes, whether added by With or in a
// Record, should be qualified by the sequence of group names.
func (h *LevelHandler) WithGroup(name string) slog.Handler {
	return NewLevelHandler(h.level, h.handler.WithGroup(name))
}
