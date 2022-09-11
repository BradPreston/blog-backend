package app

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (s *Server) Routes() http.Handler {
	router := chi.NewRouter()

	router.Route("/v1/api", func(r chi.Router) {
		r.Get("/status", s.ApiStatus())
	})

	return router
}
