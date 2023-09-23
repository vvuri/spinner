package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"os"
	"spinner/internal/config"
	"spinner/internal/http-server/middleware/logger"
	"spinner/internal/lib/logger/sl"
	"spinner/internal/storage/sqlite"
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

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	//err = storage.SaveURL("http://google.com", "google")
	//if err != nil {
	//	log.Error("failed to init storage", sl.Err(err))
	//	os.Exit(1)
	//}

	_ = storage
	//err := storage.
	query, err := storage.GetURL("google")
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	log.Info(query)

	router := chi.NewRouter()

	// добавляем к кождому запросу request-id
	router.Use(middleware.RequestID)
	// router.Use(middleware.Logger)
	router.Use(logger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

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
