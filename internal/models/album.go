package models

type Album struct {
	Id              string    `json:"id"`
	Title           string    `json:"title"`
	ReleaseYear     int       `json:"releaseYear"`
	RecordLabel     string    `json:"label"`
	Performer       Performer `json:"performer"`
	Genres          []Genre   `json:"genres"`
	CoverPictureUrl string    `json:"coverPictureUrl"`
	TotalTracks     int       `json:"totalTracks"`
}
