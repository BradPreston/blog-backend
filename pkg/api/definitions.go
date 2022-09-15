package api

import "time"

type BlogPost struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"md_body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
