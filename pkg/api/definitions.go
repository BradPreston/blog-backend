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
	ID        int       `json:"id"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	RoleID    int       `json:"role_id"`
	Username  string    `json:"username,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Comment struct {
	ID          int       `json:"id"`
	CommentBody string    `json:"comment_body"`
	UserID      int       `json:"user_id"`
	PostID      int       `json:"post_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Role struct {
	ID       int    `json:"id"`
	RoleName string `json:"role_name"`
}
