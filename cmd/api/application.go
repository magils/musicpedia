package main

import "com.mgil.musicpedia/internal/models"

type Application struct {
	config       Config
	repositories models.Repositories
}
