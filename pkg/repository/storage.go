package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/BradPreston/blog-backend/pkg/api"
)

type Storage interface {
	CreatePost(request api.BlogPost) error
	GetAllPosts() ([]*api.BlogPost, error)
	GetOnePost(id int) (*api.BlogPost, error)
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

func (s *storage) GetAllPosts() ([]*api.BlogPost, error) {
	var posts []*api.BlogPost
	query := `SELECT * FROM posts`

	rows, err := s.db.Query(query)
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
	var post api.BlogPost

	query := `SELECT * FROM posts WHERE id = $1`

	row := s.db.QueryRow(query, id)

	err := row.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		log.Printf("there was an error: %v", err.Error())
		return nil, err
	}

	return &post, nil
}
