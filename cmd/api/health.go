package main

import (
	"encoding/json"
	"net/http"

	responses "com.mgil.musicpedia/internal"
)

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	// TODO: Load Config data here
	statusResponse := map[string]string{
		"status": "OK",
	}

	err := responses.WriteJsonResponse(w, statusResponse, http.StatusOK)

	if err != nil {
		errorResponse := map[string]string{
			"error": err.Error(),
		}
		jsonResponse, _ := json.Marshal(errorResponse)
		http.Error(w, string(jsonResponse), http.StatusInternalServerError)
	}
}
