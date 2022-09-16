package app

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, data interface{}, statusMessaage string, statusCode int) {
	response := make(map[string]interface{})
	response["status"] = statusMessaage
	response["data"] = data
	resJSON, _ := json.MarshalIndent(response, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resJSON)
}
