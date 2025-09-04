package tint

import (
	"log/slog"
	"os"
	"testing"
	"time"
)

func TestMultiline(t *testing.T) {
	w := os.Stderr
	logger := slog.New(NewHandler(w, &Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.Kitchen,
	}))
	logger.Info("Starting server", "multi", `some
like
it
	hot`, "last", 2001)
}
