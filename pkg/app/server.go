package app

import (
	"log"
	"net/http"

	"github.com/BradPreston/blog-backend/pkg/api"
)

type Server struct {
	router      http.Handler
	postService api.PostService
}

func NewServer(router http.Handler, postService api.PostService) *Server {
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
