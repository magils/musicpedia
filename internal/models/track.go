package models

import "time"

type Track struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	TrackNumber int       `json:"trackNumber"`
	Duration    int       `json:"duration"`
	CreatedAt   time.Time `json:"-"`
	Performer   Performer `json:"performer"`
	Album       Album     `json:"albumN"`
}
