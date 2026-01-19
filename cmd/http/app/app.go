package app

import (
	"net/http"

	"numbers/cmd/http/routes"
	"numbers/internal/handlers"
	"numbers/internal/repositories"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(r *chi.Mux, db *pgxpool.Pool) {
	rNumbers := repositories.New(db)

	hNumbers := handlers.New(rNumbers)

	routes := routes.NRoutes{ChiRouter: r, HNumbers: hNumbers}
	routes.Mount()

	http.ListenAndServe(":8000", r)
}
