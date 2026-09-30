package repo

import (
	"context"
	"database/sql"
	"time"
)

// UserWithProfile represents a user with profile data
type UserWithProfile struct {
	ID               int64
	Name             string
	Email            string
	State            string
	PhotoPath        *string
	DateOfBirth      *time.Time
	Gender           *string
	BloodGroup       *string
	PresentAddress   *string
	PermanentAddress *string
	CreatedAt        time.Time
}

// UserRepo handles user data access
type UserRepo interface {
	List(ctx context.Context, limit, offset int) ([]UserWithProfile, int64, error)
	GetByID(ctx context.Context, id int64) (*UserWithProfile, error)
	Create(ctx context.Context, user *UserWithProfile) error
	Update(ctx context.Context, user *UserWithProfile) error
	Delete(ctx context.Context, id int64) error
}

type userRepo struct {
	db *sql.DB
}

// NewUserRepo creates a new UserRepo
func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) List(ctx context.Context, limit, offset int) ([]UserWithProfile, int64, error) {
	return nil, 0, nil
}

func (r *userRepo) GetByID(ctx context.Context, id int64) (*UserWithProfile, error) {
	return nil, sql.ErrNoRows
}

func (r *userRepo) Create(ctx context.Context, user *UserWithProfile) error {
	return nil
}

func (r *userRepo) Update(ctx context.Context, user *UserWithProfile) error {
	return nil
}

func (r *userRepo) Delete(ctx context.Context, id int64) error {
	return nil
}
