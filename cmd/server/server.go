package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	maps := newMapping()
	maps.mapUrlsToControllers(r)
	return r
}
