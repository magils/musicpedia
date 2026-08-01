package main

import (
	"encoding/json"
	"net/http"

	"com.mgil.musicpedia/internal/helpers/responses"
)

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	// TODO: Load Config data here
	statusResponse := map[string]string{
		"status": "OK",
	}

	err := responses.WriteJson(w, statusResponse, http.StatusOK, nil, true)

	if err != nil {
		errorResponse := map[string]string{
			"error": err.Error(),
		}
		jsonResponse, _ := json.Marshal(errorResponse)
		http.Error(w, string(jsonResponse), http.StatusInternalServerError)
	}
}
