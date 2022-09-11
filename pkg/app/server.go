package app

import (
	"log"

	"github.com/BradPreston/blog-backend/pkg/api"
	"github.com/go-chi/chi"
)

type Server struct {
	router      *chi.Router
	postService api.PostService
}

func NewServer(router *chi.Router, postService api.PostService) *Server {
	return &Server{
		router:      router,
		postService: postService,
	}
}

func (s *Server) Run() error {
	r := s.Routes()

	err := r.Run()

	if err != nil {
		log.Printf("server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
