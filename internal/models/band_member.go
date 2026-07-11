package models

type BandMember struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Role     string `json:"role"`
}
