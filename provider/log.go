package provider

import (
	"context"
	"log/slog"
	"os"
)

const LogKey = "log"

func WithLog(ctx context.Context) context.Context {
	l := slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level:     slog.LevelDebug,
				AddSource: true,
			},
		),
	)

	return context.WithValue(ctx, LogKey, l)
}
