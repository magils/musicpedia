package main

import (
	"encoding/json"
	"net/http"
)

func (app *Application) writeJsonResponse(w http.ResponseWriter, data any, statusCode int, headers http.Header) error {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	for key, values := range headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(statusCode)
	w.Write(jsonData)

	return nil
}
