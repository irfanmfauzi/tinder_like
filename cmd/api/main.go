package main

import (
	"log/slog"
	"net/http"
	"os"
	"tinder_like/database"
	"tinder_like/internal/routes"
	"tinder_like/middleware"

	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	handler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	db, err := database.Connect()
	if err != nil {
		return
	}

	mux := routes.RegisterRoute(db)
	server := http.Server{
		Addr:    ":8000",
		Handler: middleware.LoggerMiddleware(mux),
	}

	logger.Info("Tinder Like Apps")
	server.ListenAndServe()
}
