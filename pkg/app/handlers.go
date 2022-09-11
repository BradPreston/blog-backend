package app

import (
	"encoding/json"
	"net/http"
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
