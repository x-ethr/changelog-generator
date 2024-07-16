package git

import (
	"context"
	"log/slog"
	"testing"
)

func TestVerify(t *testing.T) {
	ctx := context.Background()
	slog.SetLogLoggerLevel(slog.LevelDebug)
	t.Run("Verify", func(t *testing.T) {
		e := Verify(ctx)
		if e != nil {
			t.Fatal("Unexpected Error", e)
		}

		t.Log("Successfully Verified `git` Executable")
	})
}
