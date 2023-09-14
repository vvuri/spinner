package main

import (
	"log/slog"
	"os"
	"spinner/internal/config"
)

const (
	envLocal   = "local"
	envStaging = "stage"
	envProd    = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("Start logger")

	// TODO: init storage: sqlite

	// TODO: init router: chi, chi render

	// TODO: init server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug}))
	case envStaging:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo}))
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelWarn}))
	}

	return log
}
