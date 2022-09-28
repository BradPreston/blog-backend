package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/BradPreston/blog-backend/pkg/api"
)

type Storage interface {
	CreatePost(request api.BlogPost) error
	GetAllPosts() ([]*api.BlogPost, error)
	GetOnePost(id int) (*api.BlogPost, error)
	UpdatePost(post *api.BlogPost) error
	DeletePost(id int) error
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) CreatePost(request api.BlogPost) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	NewBlogPostStmt := `INSERT INTO posts (title, md_body, created_at, updated_at) VALUES ($1, $2, $3, $4)`

	err := s.db.QueryRowContext(ctx, NewBlogPostStmt, request.Title, request.Body, time.Now(), time.Now()).Err()
	if err != nil {
		log.Printf("there was an error: %v", err.Error())
		return err
	}

	return nil
}

func (s *storage) GetAllPosts() ([]*api.BlogPost, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var posts []*api.BlogPost
	query := `SELECT * FROM posts`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("there was an error: %v", err.Error())
	}

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

func (s *storage) UpdatePost(post *api.BlogPost) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE posts SET title = $1, md_body = $2, updated_at = $3 WHERE id = $4`

	_, err := s.db.ExecContext(ctx, query, post.Title, post.Body, time.Now(), post.ID)
	if err != nil {
		log.Printf("there was an error: %v", err.Error())
		return err
	}

	return nil
}

func (s *storage) DeletePost(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM posts WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("there was an error: %v", err.Error())
		return err
	}

	return nil
}
