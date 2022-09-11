package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ErrorJSON(w http.ResponseWriter, err error, message string, statusCode int) {
	response := map[string]string{
		"status": "fail",
		"error":  fmt.Sprint(err),
		"data":   message,
	}
	resJSON, _ := json.MarshalIndent(response, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resJSON)
}

func SuccessJSON(w http.ResponseWriter, message string, statusCode int) {
	response := map[string]string{
		"status": "success",
		"data":   message,
	}
	resJSON, _ := json.MarshalIndent(response, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resJSON)
}
