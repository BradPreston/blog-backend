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
		r.Put("/posts/{id}", s.UpdatePost)
		r.Delete("/posts/{id}", s.DeletePost)
		// user routes
		r.Post("/users", s.CreateUser)
		r.Get("/users", s.GetAllUsers)
		r.Get("/users/{id}", s.GetOneUser)
		r.Put("/users/{id}", s.UpdateUser)
		r.Put("/users/{id}/update_password", s.UpdatePassword)
		r.Delete("/users/{id}", s.DeleteUser)
	})

	return router
}
