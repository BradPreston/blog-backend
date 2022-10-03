package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func (s *Server) Routes() http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "*"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

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
		r.Patch("/users/{id}/update_password", s.UpdatePassword)
		r.Delete("/users/{id}", s.DeleteUser)
	})

	return router
}
