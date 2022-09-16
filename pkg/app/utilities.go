package app

import (
	"encoding/json"
	"net/http"
)

func ErrorJSON(w http.ResponseWriter, message string, statusCode int) {
	response := map[string]string{
		"status": "fail",
		"data":   message,
	}
	resJSON, _ := json.MarshalIndent(response, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resJSON)
}

func SuccessJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	response := make(map[string]interface{})
	response["status"] = "success"
	response["data"] = data
	resJSON, _ := json.MarshalIndent(response, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resJSON)
}
