package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) Routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/health", app.HealthCheck)
	router.HandlerFunc(http.MethodGet, "/v1/genres", app.ListGenres)
	router.HandlerFunc(http.MethodPost, "/v1/genres", app.CreateGenre)
	router.HandlerFunc(http.MethodPut, "/v1/genres/:id", app.UpdateGenre)
	router.HandlerFunc(http.MethodDelete, "/v1/genres/:id", app.DeleteGenre)

	return router
}
