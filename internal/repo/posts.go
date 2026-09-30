package repo

import (
	"context"
	"database/sql"
	"time"
)

// Post represents a blog post in the repository layer
type Post struct {
	ID          int64
	Title       string
	Body        string
	Slug        string
	Thumbnail   *string
	PublishedAt *time.Time
	CreatedAt   time.Time
}

// PostRepo handles post data access
type PostRepo interface {
	List(ctx context.Context, limit, offset int) ([]Post, int64, error)
	GetByID(ctx context.Context, id int64) (*Post, error)
	GetBySlug(ctx context.Context, slug string) (*Post, error)
	Create(ctx context.Context, post *Post) error
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id int64) error
}

type postRepo struct {
	db *sql.DB
}

// NewPostRepo creates a new PostRepo
func NewPostRepo(db *sql.DB) PostRepo {
	return &postRepo{db: db}
}

func (r *postRepo) List(ctx context.Context, limit, offset int) ([]Post, int64, error) {
	return nil, 0, nil
}

func (r *postRepo) GetByID(ctx context.Context, id int64) (*Post, error) {
	return nil, sql.ErrNoRows
}

func (r *postRepo) GetBySlug(ctx context.Context, slug string) (*Post, error) {
	return nil, sql.ErrNoRows
}

func (r *postRepo) Create(ctx context.Context, post *Post) error {
	return nil
}

func (r *postRepo) Update(ctx context.Context, post *Post) error {
	return nil
}

func (r *postRepo) Delete(ctx context.Context, id int64) error {
	return nil
}
