package main

import (
	"context"
	"log"

	"numbers/cmd/http/app"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	ctx := context.Background()

	db, err := pgxpool.New(ctx, cfg.DBDSL)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// init logger slog

	app.Run(
		r,
		db,
		// cfg,
		// logger
	)
}
