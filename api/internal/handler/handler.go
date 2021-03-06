package handler

import (
	"fmt"
	"net/http"

	"github.com/badis/hackathon/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

type handler struct {
	*service.Service
}

func root(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "API works!\n")
}

// New creates an http.Handler with predefined routing.
func New(s *service.Service) http.Handler {
	h := &handler{s}

	api := chi.NewRouter()

	api.Get("/", root)
	api.Post("/login", h.login)
	api.Post("/register", h.registerPatient)

	apiWithCors := cors.Default().Handler(api)
	return apiWithCors
}
