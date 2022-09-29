package app

import (
	"fmt"
	"net/http"

	"github.com/BradPreston/blog-backend/pkg/api"
)

type Server struct {
	router      http.Handler
	postService api.PostService
	userService api.UserService
}

func NewServer(router http.Handler, postService api.PostService, userService api.UserService) *Server {
	return &Server{
		router:      router,
		postService: postService,
		userService: userService,
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
