package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/BradPreston/blog-backend/pkg/api"
)

type Storage interface {
	// posts
	CreatePost(post api.BlogPost) error
	GetAllPosts() ([]*api.BlogPost, error)
	GetOnePost(id int) (*api.BlogPost, error)
	UpdatePost(post *api.BlogPost) error
	DeletePost(id int) error

	// users
	CreateUser(user api.User) error
	GetAllUsers() ([]*api.User, error)
	GetOneUser(id int) (*api.User, error)
	UpdateUser(user *api.User) error
	DeleteUser(id int) error

	// comments
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

// CreatePost inserts a blog post into the database
func (s *storage) CreatePost(post api.BlogPost) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
	INSERT INTO posts (title, md_body, created_at, updated_at)
	VALUES ($1, $2, $3, $4)`

	err := s.db.QueryRowContext(ctx, stmt, post.Title, post.Body, time.Now(), time.Now()).Err()
	if err != nil {
		log.Printf("there was an error: %v", err.Error())
		return err
	}

	return nil
}

// GetAllPosts returns every blog post from the database
func (s *storage) GetAllPosts() ([]*api.BlogPost, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var posts []*api.BlogPost
	query := `SELECT * FROM posts`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("there was an error: %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var post api.BlogPost

		err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			log.Printf("there was an error: %v", err.Error())
			return nil, err
		}

		posts = append(posts, &post)
	}

	return posts, nil
}

// GetOnePost returns one blog post from the database
func (s *storage) GetOnePost(id int) (*api.BlogPost, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var post api.BlogPost

	query := `SELECT * FROM posts WHERE id = $1`

	row := s.db.QueryRowContext(ctx, query, id)

	err := row.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		log.Printf("there was an error: %v", err.Error())
		return nil, err
	}

	return &post, nil
}

// UpdatePost updates one post in the database
func (s *storage) UpdatePost(post *api.BlogPost) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
	UPDATE posts
	SET title = $1, md_body = $2, updated_at = $3
	WHERE id = $4`

	_, err := s.db.ExecContext(ctx, query, post.Title, post.Body, time.Now(), post.ID)
	if err != nil {
		log.Printf("there was an error: %v", err.Error())
		return err
	}

	return nil
}

// DeletePost deletes one post from the database
func (s *storage) DeletePost(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `DELETE FROM posts WHERE id = $1`

	_, err := s.db.ExecContext(ctx, stmt, id)
	if err != nil {
		log.Printf("there was an error: %v", err.Error())
		return err
	}

	return nil
}

// CreateUser inserts one user into the database
func (s *storage) CreateUser(user api.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
	INSERT INTO users (email, password, username, first_name, last_name, role_id, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	err := s.db.QueryRowContext(ctx, stmt, user.Email, user.Password, user.Username, user.FirstName, user.LastName, 2, time.Now(), time.Now()).Err()
	if err != nil {
		log.Printf("there was an error: %v", err.Error())
		return err
	}

	return nil
}

// GetAllUsers returns every user from the database
func (s *storage) GetAllUsers() ([]*api.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var users []*api.User

	query := `
	SELECT id, email, username, first_name, last_name, role_id
	FROM users`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("there was an error: %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user api.User

		err := rows.Scan(&user.ID, &user.Email, &user.RoleID)
		if err != nil {
			log.Printf("there was an error: %v", err.Error())
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

// GetOneUser returns one user from the database
func (s *storage) GetOneUser(id int) (*api.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user api.User

	query := `
	SELECT id, email, username, first_name, last_name, role_id
	FROM users
	WHERE id = $1`

	row := s.db.QueryRowContext(ctx, query, id)

	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.FirstName, &user.LastName, &user.RoleID)
	if err != nil {
		log.Printf("there was an error: %v", err.Error())
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates one user in the database
func (s *storage) UpdateUser(user *api.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
	UPDATE users
	SET email = $1, username = $2, first_name = $3, last_name = $4, updated_at = $5
	WHERE id = $6
	`

	_, err := s.db.ExecContext(ctx, query, user.Email, user.Username, user.FirstName, user.LastName, time.Now(), user.ID)
	if err != nil {
		log.Printf("there was an error: %v", err.Error())
		return err
	}

	return nil
}

// DeleteUser deletes one user from the database
func (s *storage) DeleteUser(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
	DELETE FROM users
	WHERE id = $1`

	_, err := s.db.ExecContext(ctx, stmt, id)
	if err != nil {
		log.Printf("there was an error: %v", err.Error())
		return err
	}

	return nil
}
