package routes

import (
	"numbers/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func MountNumbers(r *chi.Mux, h *handlers.Handler) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/numbers", h.AddNumber)
	})
}
