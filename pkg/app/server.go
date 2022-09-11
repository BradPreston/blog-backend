package app

import (
	"fmt"
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

	fmt.Println("application is running on port 8080")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		return err
	}

	return nil
}
