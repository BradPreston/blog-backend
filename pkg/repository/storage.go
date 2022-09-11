package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/BradPreston/blog-backend/pkg/api"
)

type Storage interface {
	CreatePost(request api.BlogPost) error
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
	NewBlogPostStmt := `INSERT INTO posts (title, md_body, created_at, updated_at) VALUES ($1, $2, $3, $4)`

	err := s.db.QueryRow(NewBlogPostStmt, request.Title, request.Body, time.Now(), time.Now()).Err()
	if err != nil {
		log.Printf("there was an error: %v", err.Error())
		return err
	}

	return nil
}
