package responses

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, data any, status int, headers http.Header, isOutputIndented bool) error {
	var jsonData []byte
	var err error

	if isOutputIndented {
		jsonData, err = json.MarshalIndent(data, "", " ")
	} else {
		jsonData, err = json.Marshal(data)
	}

	if err != nil {
		return err
	}

	jsonData = append(jsonData, '\n')

	for key, values := range headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)

	return nil
}

func WriteJsonResponse(w http.ResponseWriter, data any, status int, headers http.Header) error {
	return WriteJson(w, data, status, headers, false)
}

func WriteEnvelopedJsonResponse(w http.ResponseWriter, envelopName string, data any, status int, headers http.Header) error {
	envelopData := map[string]any{
		envelopName: data,
	}
	return WriteJson(w, envelopData, status, headers, false)
}

func WriteOkResponse(w http.ResponseWriter, data any) error {
	return WriteJson(w, data, http.StatusOK, nil, false)
}

func WriteCreatedResponse(w http.ResponseWriter, data any) error {
	return WriteJson(w, data, http.StatusCreated, nil, false)
}

func WriteNotFoundResponse(w http.ResponseWriter) error {
	data := map[string]string{
		"error": "Entity not found",
	}
	return WriteJson(w, data, http.StatusNotFound, nil, false)
}

func WriteNoContentResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func WriteErrorResponse(w http.ResponseWriter, status int, payload any) {
	wrapper := map[string]any{"error": payload}
	err := WriteJson(w, wrapper, status, nil, false)

	if err != nil {
		w.WriteHeader(500)
	}
}

func WriteBadRequestErrorResponse(w http.ResponseWriter, err error) {
	WriteErrorResponse(w, http.StatusBadRequest, err)
}

func WriteInternalServerErrorResponse(w http.ResponseWriter, err error) {
	WriteErrorResponse(w, http.StatusInternalServerError, err)
}
