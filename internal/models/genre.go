package models

import (
	"context"
	"database/sql"
	"time"
)

type Genre struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GenreRepository struct {
	DB *sql.DB
}

func (g *GenreRepository) GetAll() ([]*Genre, error) {
	query := `
		SELECT id, name, description
		FROM genres
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	rows, err := g.DB.QueryContext(ctx, query)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	genres := []*Genre{}

	for rows.Next() {
		var genre Genre

		err := rows.Scan(
			&genre.Id,
			&genre.Name,
		)

		if err != nil {
			return nil, err
		}

		genres = append(genres, &genre)
	}

	return genres, nil
}
