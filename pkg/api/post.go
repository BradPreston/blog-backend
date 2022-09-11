package api

type PostService interface{}

type PostRepository interface{}

type postService struct {
	storage PostRepository
}

func NewPostService(postRepo PostRepository) PostService {
	return &postService{
		storage: postRepo,
	}
}
