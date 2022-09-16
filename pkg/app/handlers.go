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
	response := map[string]string{
		"status": "success",
		"data":   "blog post API is running properly",
	}

	resJSON, _ := json.MarshalIndent(response, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}

func (s *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post api.BlogPost

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		ErrorJSON(w, "could not read post body", http.StatusConflict)
		return
	}

	err = s.postService.New(post)

	if err != nil {
		ErrorJSON(w, "could not create blog post", http.StatusConflict)
		return
	}

	SuccessJSON(w, "post created successfully", http.StatusCreated)
}

func (s *Server) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := s.postService.GetAll()

	if err != nil {
		ErrorJSON(w, "could not get all posts", http.StatusConflict)
		return
	}

	SuccessJSON(w, posts, http.StatusOK)
}

func (s *Server) GetOnePost(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(uri[4])
	if err != nil {
		ErrorJSON(w, fmt.Sprintf("could not find id [%d] in URI", id), http.StatusConflict)
		return
	}

	post, err := s.postService.GetOne(id)
	if err != nil {
		ErrorJSON(w, fmt.Sprintf("could not get post by id: %d", id), http.StatusConflict)
		return
	}

	SuccessJSON(w, post, http.StatusOK)
}

func (s *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(uri[4])
	if err != nil {
		ErrorJSON(w, fmt.Sprintf("could not find id [%d] in URI", id), http.StatusConflict)
		return
	}

	var updatedPost api.BlogPost

	err = json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		ErrorJSON(w, "could not read post body", http.StatusOK)
		return
	}

	updatedPost.UpdatedAt = time.Now()
	updatedPost.ID = id

	if err != nil {
		ErrorJSON(w, "could not update post", http.StatusConflict)
	}

	err = s.postService.Update(&updatedPost)
	if err != nil {
		ErrorJSON(w, "could not update post", http.StatusConflict)
		return
	}

	SuccessJSON(w, "post updated successfully", http.StatusOK)
}

func (s *Server) DeletePost(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(uri[4])
	if err != nil {
		ErrorJSON(w, fmt.Sprintf("could not find id [%d] in URI", id), http.StatusConflict)
		return
	}

	err = s.postService.Delete(id)
	if err != nil {
		ErrorJSON(w, fmt.Sprintf("could not delete post by id: %d", id), http.StatusConflict)
		return
	}

	SuccessJSON(w, "post deleted successfully", http.StatusOK)
}
