package api

import "time"

type BlogPost struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"md_body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
}
