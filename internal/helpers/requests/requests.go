package requests

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func ParseJsonToEntity(r *http.Request, entity any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(entity)
}

func GetQueryParam(r *http.Request, paramName string) string {
	params := httprouter.ParamsFromContext(r.Context())
	return params.ByName(paramName)
}

// TODO: Should I return an error? It can be helpful if user pass a non-numerical value
func GetIdParam(r *http.Request) int64 {
	idValue := GetQueryParam(r, "id")
	id, err := strconv.ParseInt(idValue, 10, 64)

	if err != nil {
		return 0
	}

	return id
}
