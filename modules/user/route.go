package user

import (
	"github.com/go-chi/chi/v5"
)

func PlayerRoutes(router chi.Router) {
	router.Get("/", List)
	router.Get("/{id}", Get)
}