package app

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (s *Server) Routes() http.Handler {
	router := chi.NewRouter()

	router.Route("/v1/api", func(r chi.Router) {
		// status routes
		r.Get("/status", s.ApiStatus)
		// blog post routes
		r.Post("/posts", s.CreatePost)
		r.Get("/posts", s.GetAllPosts)
		r.Get("/posts/{id}", s.GetOnePost)
	})

	return router
}
