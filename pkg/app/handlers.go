package app

import (
	"encoding/json"
	"net/http"

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
		ErrorJSON(w, err, "could not read post body", http.StatusConflict)
		return
	}

	err = s.postService.New(post)

	if err != nil {
		ErrorJSON(w, err, "could not create blog post", http.StatusConflict)
		return
	}

	SuccessJSON(w, "post created successfully", http.StatusOK)
}
