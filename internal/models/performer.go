package models

import (
	"context"
	"database/sql"
	"time"
)

type Performer struct {
	Id         string       `json:"id"`
	Name       string       `json:"name"`
	Bio        string       `json:"bio"`
	Country    string       `json:"country"`
	FormedYear int32        `json:"formedYear"`
	IsBand     bool         `json:"isBand"`
	PictureUrl string       `json:"pictureUrl"`
	Members    []BandMember `json:"members"`
	CreatedAt  time.Time    `json:"-"`
	UpdatedAt  time.Time    `json:"-"`
}

type PerformerRepository struct {
	DB *sql.DB
}

func (p *PerformerRepository) GetAll() ([]Performer, error) {
	query := `
		SELECT id, name, bio, country, formed_year, is_band, picture_url
		FROM perfomers
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := p.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	perfomers := []Performer{}

	for rows.Next() {
		var perfomer Performer

		err := rows.Scan(
			&perfomer.Id,
			&perfomer.Name,
			&perfomer.Country,
			&perfomer.FormedYear,
			&perfomer.IsBand,
			&perfomer.PictureUrl,
		)

		if err != nil {
			return nil, err
		}

		perfomers = append(perfomers, perfomer)
	}

	return perfomers, nil
}
