package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/BradPreston/blog-backend/pkg/api"
)

func (s *Server) ApiStatus(w http.ResponseWriter, r *http.Request) {
	ResponseJSON(w, "blog post API is running properly", "success", http.StatusOK)
}

func (s *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post api.BlogPost

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		ResponseJSON(w, "could not read post body", "fail", http.StatusConflict)
		return
	}

	err = s.postService.New(post)

	if err != nil {
		ResponseJSON(w, "could not create blog post", "fail", http.StatusConflict)
		return
	}

	ResponseJSON(w, "post created successfully", "success", http.StatusCreated)
}

func (s *Server) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := s.postService.GetAll()

	if err != nil {
		ResponseJSON(w, "could not get all posts", "fail", http.StatusConflict)
		return
	}

	ResponseJSON(w, posts, "success", http.StatusOK)
}

func (s *Server) GetOnePost(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(uri[4])
	if err != nil {
		ResponseJSON(w, fmt.Sprintf("could not find id [%d] in URI", id), "fail", http.StatusConflict)
		return
	}

	post, err := s.postService.GetOne(id)
	if err != nil {
		ResponseJSON(w, fmt.Sprintf("could not get post by id: %d", id), "fail", http.StatusConflict)
		return
	}

	ResponseJSON(w, post, "success", http.StatusOK)
}

func (s *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(uri[4])
	if err != nil {
		ResponseJSON(w, fmt.Sprintf("could not find id [%d] in URI", id), "fail", http.StatusConflict)
		return
	}

	var updatedPost api.BlogPost

	err = json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		ResponseJSON(w, "could not read post body", "success", http.StatusOK)
		return
	}

	updatedPost.UpdatedAt = time.Now()
	updatedPost.ID = id

	if err != nil {
		ResponseJSON(w, "could not update post", "fail", http.StatusConflict)
	}

	err = s.postService.Update(&updatedPost)
	if err != nil {
		ResponseJSON(w, "could not update post", "fail", http.StatusConflict)
		return
	}

	ResponseJSON(w, "post updated successfully", "success", http.StatusOK)
}

func (s *Server) DeletePost(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(uri[4])
	if err != nil {
		ResponseJSON(w, fmt.Sprintf("could not find id [%d] in URI", id), "fail", http.StatusConflict)
		return
	}

	err = s.postService.Delete(id)
	if err != nil {
		ResponseJSON(w, fmt.Sprintf("could not delete post by id: %d", id), "fail", http.StatusConflict)
		return
	}

	ResponseJSON(w, "post deleted successfully", "success", http.StatusOK)
}
