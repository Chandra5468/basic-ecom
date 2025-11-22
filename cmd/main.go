package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Chandra5468/basic-ecom/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			// dsn: "postgres://postgres:user@localhost/cfq?sslmode=disable",
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=user dbname=postgres sslmode=disable"),
		},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)) // we can also use json handler for slog.NewJsonHandler
	slog.SetDefault(logger)
	// Database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
		// os.Exit(1)
	}
	defer conn.Close(ctx)

	logger.Info("connected to database", "dsn", cfg.db.dsn)
	api := &application{
		config: cfg,
		db:     conn,
	}

	// passing this as a dependency. So all files make logs to this logger
	if err := api.run(api.mount()); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
