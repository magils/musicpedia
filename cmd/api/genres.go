package main

import (
	"net/http"

	"com.mgil.musicpedia/internal/helpers/requests"
	"com.mgil.musicpedia/internal/helpers/responses"
	"com.mgil.musicpedia/internal/schemas"

	"com.mgil.musicpedia/internal/models"
)

func (app *Application) ListGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.repositories.Genre.GetAll()

	if err != nil {
		responses.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses.WriteEnvelopedJsonResponse(w, "genres", genres, http.StatusOK, nil)
}

func (app *Application) CreateGenre(w http.ResponseWriter, r *http.Request) {
	var genre models.Genre
	err := requests.ParseJsonToEntity(r, &genre)

	if err != nil {
		responses.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = app.repositories.Genre.Insert(&genre)
	responses.WriteCreatedResponse(w, genre)
}

func (app *Application) UpdateGenre(w http.ResponseWriter, r *http.Request) {
	var genreUpdateData schemas.GenreUpdateSchema
	err := requests.ParseJsonToEntity(r, &genreUpdateData)
	entityId := requests.GetIdParam(r)

	if err != nil {
		responses.WriteBadRequestErrorResponse(w, err)
		return
	}

	updatedGenre, err := app.repositories.Genre.Update(entityId, genreUpdateData)

	if err != nil {
		responses.WriteInternalServerErrorResponse(w, err)
		return
	}

	responses.WriteOkResponse(w, updatedGenre)
}

func (app *Application) DeleteGenre(w http.ResponseWriter, r *http.Request) {
	entityId := requests.GetIdParam(r)
	entityExists, err := app.repositories.Genre.Exists(entityId)

	// TODO: Improve check to handle errors better
	if !entityExists || err != nil {
		responses.WriteNotFoundResponse(w)
		return
	}

	err = app.repositories.Genre.Delete(entityId)

	if err != nil {
		responses.WriteInternalServerErrorResponse(w, err)
		return
	}

	responses.WriteNoContentResponse(w)
}
