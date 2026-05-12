package responses

import (
	"encoding/json"
	"net/http"
)

func WriteJsonResponse(w http.ResponseWriter, data any, httpStatus int) error {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(jsonData)

	return nil
}
