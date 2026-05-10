package models

import "time"

type Artist struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Bio        string    `json:"bio"`
	Country    string    `json:"country"`
	FormedYear int32     `json:formedYear`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}
