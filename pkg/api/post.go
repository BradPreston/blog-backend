package api

import (
	"errors"
	"strings"
)

type PostService interface {
	New(post BlogPost) error
	GetAll() ([]*BlogPost, error)
	GetOne(id int) (*BlogPost, error)
}

type PostRepository interface {
	CreatePost(BlogPost) error
	GetAllPosts() ([]*BlogPost, error)
	GetOnePost(int) (*BlogPost, error)
}

type postService struct {
	storage PostRepository
}

func NewPostService(postRepo PostRepository) PostService {
	return &postService{
		storage: postRepo,
	}
}

func (p *postService) New(post BlogPost) error {
	if post.Title == "" {
		return errors.New("blog post - title is required")
	}

	if post.Body == "" {
		return errors.New("blog post - body is required")
	}

	post.Title = strings.ToLower(post.Title)

	err := p.storage.CreatePost(post)
	if err != nil {
		return err
	}

	return nil
}

func (p *postService) GetAll() ([]*BlogPost, error) {
	posts, err := p.storage.GetAllPosts()

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *postService) GetOne(id int) (*BlogPost, error) {
	post, err := p.storage.GetOnePost(id)
	if err != nil {
		return nil, err
	}

	return post, nil
}
