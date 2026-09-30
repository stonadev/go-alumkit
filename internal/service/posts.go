package service

import (
	"context"
	"time"

	"github.com/stonadev/alumkit/internal/repo"
)

// Post represents a post in the service layer
type Post struct {
	ID          int64
	Title       string
	Body        string
	Slug        string
	Thumbnail   *string
	PublishedAt *time.Time
	CreatedAt   time.Time
}

// PostList represents a paginated list of posts
type PostList struct {
	Posts      []Post
	Total      int64
	Page       int
	PerPage    int
	TotalPages int
}

// PostService handles post business logic
type PostService interface {
	ListPublished(ctx context.Context, page, perPage int) (*PostList, error)
	ListAll(ctx context.Context, page, perPage int) (*PostList, error)
	GetByID(ctx context.Context, id int64) (*Post, error)
	GetBySlug(ctx context.Context, slug string) (*Post, error)
	Create(ctx context.Context, post *Post) error
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id int64) error
}

type postService struct {
	postRepo repo.PostRepo
}

// NewPostService creates a new PostService
func NewPostService(postRepo repo.PostRepo) PostService {
	return &postService{postRepo: postRepo}
}

func (s *postService) ListPublished(ctx context.Context, page, perPage int) (*PostList, error) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 20
	}

	offset := (page - 1) * perPage
	posts, total, err := s.postRepo.List(ctx, perPage, offset)
	if err != nil {
		return nil, err
	}

	// Filter only published posts
	var published []Post
	for _, p := range posts {
		if p.PublishedAt != nil {
			published = append(published, s.toPost(p))
		}
	}

	return &PostList{
		Posts:      published,
		Total:      total,
		Page:       page,
		PerPage:    perPage,
		TotalPages: int(total) / perPage,
	}, nil
}

func (s *postService) ListAll(ctx context.Context, page, perPage int) (*PostList, error) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 20
	}

	offset := (page - 1) * perPage
	posts, total, err := s.postRepo.List(ctx, perPage, offset)
	if err != nil {
		return nil, err
	}

	result := make([]Post, len(posts))
	for i, p := range posts {
		result[i] = s.toPost(p)
	}

	totalPages := int(total) / perPage
	if int(total)%perPage > 0 {
		totalPages++
	}

	return &PostList{
		Posts:      result,
		Total:      total,
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
	}, nil
}

func (s *postService) GetByID(ctx context.Context, id int64) (*Post, error) {
	p, err := s.postRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	post := s.toPost(*p)
	return &post, nil
}

func (s *postService) GetBySlug(ctx context.Context, slug string) (*Post, error) {
	p, err := s.postRepo.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	post := s.toPost(*p)
	return &post, nil
}

func (s *postService) Create(ctx context.Context, post *Post) error {
	return s.postRepo.Create(ctx, &repo.Post{
		Title:     post.Title,
		Body:      post.Body,
		Slug:      post.Slug,
		Thumbnail: post.Thumbnail,
	})
}

func (s *postService) Update(ctx context.Context, post *Post) error {
	return s.postRepo.Update(ctx, &repo.Post{
		ID:          post.ID,
		Title:       post.Title,
		Body:        post.Body,
		Slug:        post.Slug,
		Thumbnail:   post.Thumbnail,
		PublishedAt: post.PublishedAt,
	})
}

func (s *postService) Delete(ctx context.Context, id int64) error {
	return s.postRepo.Delete(ctx, id)
}

func (s *postService) toPost(p repo.Post) Post {
	return Post{
		ID:          p.ID,
		Title:       p.Title,
		Body:        p.Body,
		Slug:        p.Slug,
		Thumbnail:   p.Thumbnail,
		PublishedAt: p.PublishedAt,
		CreatedAt:   p.CreatedAt,
	}
}
