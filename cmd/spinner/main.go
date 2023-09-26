package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"os"
	"spinner/internal/config"
	"spinner/internal/http-server/handlers/url/save"
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

	router.Post("/url", save.New(log, storage, cfg.AliasLength))

	log.Info("server starting", slog.String("address", cfg.HttpServer.Address))

	srv := &http.Server{
		Addr:         cfg.HttpServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HttpServer.TimeOut,
		WriteTimeout: cfg.HttpServer.TimeOut,
		IdleTimeout:  cfg.HttpServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stopped")
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
