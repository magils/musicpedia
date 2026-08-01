package schemas

type GenreUpdateSchema struct {
	Name        *string `json:"title"`
	Description *string `json:"description"`
}
