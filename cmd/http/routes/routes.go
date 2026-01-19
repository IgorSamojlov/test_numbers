package routes

import (
	"net/http"
	"numbers/internal/handlers"

	"github.com/go-chi/chi/v5"
)

type NRoutes struct {
	HNumbers  *handlers.Handler
	ChiRouter *chi.Mux
}

func (r *NRoutes) Mount() {
	MountNumbers(r.ChiRouter, r.HNumbers)
	r.ChiRouter.Get("/helthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I am gooood"))
	})
}
