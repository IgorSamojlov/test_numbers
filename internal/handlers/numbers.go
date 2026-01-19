package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) AddNumber(w http.ResponseWriter, r *http.Request) {
	data := &numberRequest{}

	err := render.Bind(r, data)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, ErrResponse{ErrorText: err.Error()})

		return
	}

	numbers, err := h.Repository.AddNumber(r.Context(), data.Number)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, ErrResponse{ErrorText: err.Error()})

		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &HttpResponse{Numbers: numbers})
}
