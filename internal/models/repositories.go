package models

import "database/sql"

type Repositories struct {
	Genre GenreRepository
}

func NewRepositories(db *sql.DB) Repositories {
	return Repositories{
		Genre: GenreRepository{DB: db},
	}
}
