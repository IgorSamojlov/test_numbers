package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	StatusText string `json:"status"`
	ErrorText  string `json:"error,omitempty"`
}

func (e ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusBadRequest)
	return nil
}

type HttpResponse struct {
	Error   string `json:"error,omitempty"`
	Numbers any    `json:"record,omitempty"`
}

type numberRequest struct {
	Number int
}

type Repository interface {
	AddNumber(ctx context.Context, number int) ([]int, error)
}

type Handler struct {
	Repository Repository
}

func (o *HttpResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func New(r Repository) *Handler {
	return &Handler{
		Repository: r,
	}
}
func (nr *numberRequest) Bind(r *http.Request) error {
	return nil
}
