package entity

import "time"

type Documentation struct {
	name            string
	article         string
	event           Event
	createdAt       *time.Time
	updatedAt       *time.Time
	createdBy       string //siapa
	updatedBy       string //siapa
	photo_cover_url string
	gallery         *[]Photo
}

type Photo struct {
	url  string
	name string
	id   string
}
