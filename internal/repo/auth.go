package repo

import (
	"context"
	"database/sql"
	"time"
)

// User represents a user in the repository layer
type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	State     string
	CreatedAt time.Time
}

// AuthRepo handles authentication data access
type AuthRepo interface {
	GetUserByID(ctx context.Context, id int64) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateSession(ctx context.Context, userID int64, expiresAt time.Time) (string, error)
	GetSession(ctx context.Context, sessionID string) (*int64, error)
	DeleteSession(ctx context.Context, sessionID string) error
	DeleteExpiredSessions(ctx context.Context) error
}

type authRepo struct {
	db *sql.DB
}

// NewAuthRepo creates a new AuthRepo
func NewAuthRepo(db *sql.DB) AuthRepo {
	return &authRepo{db: db}
}

func (r *authRepo) GetUserByID(ctx context.Context, id int64) (*User, error) {
	return nil, sql.ErrNoRows
}

func (r *authRepo) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	return nil, sql.ErrNoRows
}

func (r *authRepo) CreateSession(ctx context.Context, userID int64, expiresAt time.Time) (string, error) {
	return "", sql.ErrNoRows
}

func (r *authRepo) GetSession(ctx context.Context, sessionID string) (*int64, error) {
	return nil, sql.ErrNoRows
}

func (r *authRepo) DeleteSession(ctx context.Context, sessionID string) error {
	return nil
}

func (r *authRepo) DeleteExpiredSessions(ctx context.Context) error {
	return nil
}
