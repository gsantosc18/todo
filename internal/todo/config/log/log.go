package config

import (
	"log/slog"
	"os"
)

func init() {
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	slog.SetDefault(jsonLogger)
}
