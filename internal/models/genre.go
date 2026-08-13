package models

import (
	"context"
	"database/sql"
	"time"

	"com.mgil.musicpedia/internal/schemas"
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
			&genre.Description,
		)

		if err != nil {
			return nil, err
		}

		genres = append(genres, &genre)
	}

	return genres, nil
}

func (g *GenreRepository) Exists(id int64) (bool, error) {
	query := `
		SELECT 1 FROM genres WHERE id = ?
	`
	ctx, cancel := CreateDefaultContext()
	defer cancel()

	var exists int

	err := g.DB.QueryRowContext(ctx, query, id).Scan(&exists)

	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (g *GenreRepository) Get(id int64) (*Genre, error) {
	query := `
		SELECT id, name, description
		FROM genres
		WHERE id = ?
	`
	ctx, cancel := CreateDefaultContext()
	defer cancel()

	var genre Genre

	err := g.DB.QueryRowContext(ctx, query, id).Scan(
		&genre.Id,
		&genre.Name,
		&genre.Description,
	)

	if err != nil {
		return nil, err
	}

	return &genre, nil
}

func (g *GenreRepository) Insert(genre *Genre) error {
	query := `
		INSERT INTO genres(name, description)
		VALUES (?, ?)
	`
	args := []any{genre.Name, genre.Description}
	ctx, cancel := CreateDefaultContext()

	defer cancel()

	result, err := g.DB.ExecContext(ctx, query, args...)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	genre.Id = int(id)

	return nil
}

func (g *GenreRepository) Update(id int64, data schemas.GenreUpdateSchema) (*Genre, error) {
	existingGenre, err := g.Get(id)

	if err != nil {
		return nil, err
	}

	if data.Name != nil {
		existingGenre.Name = *data.Name
	}

	if data.Description != nil {
		existingGenre.Description = *data.Description
	}

	query := `
		UPDATE genres
		SET name = ?, description = ?
		WHERE id = ?
	`
	args := []any{existingGenre.Name, existingGenre.Description, id}
	ctx, cancel := CreateDefaultContext()
	defer cancel()

	_, dbErr := g.DB.ExecContext(ctx, query, args...)

	if dbErr != nil {
		return nil, err
	}

	return existingGenre, nil
}

func (g *GenreRepository) Delete(id int64) error {
	query := `
		DELETE FROM genres
		WHERE id = ?
	`
	ctx, cancel := CreateDefaultContext()
	defer cancel()

	_, err := g.DB.ExecContext(ctx, query, id)
	return err
}
