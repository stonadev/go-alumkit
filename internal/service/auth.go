package service

import (
	"context"
	"errors"

	"github.com/stonadev/alumkit/internal/repo"
)

// Errors
var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserInactive       = errors.New("user account is inactive")
)

// AuthUser represents a user in the service layer
type AuthUser struct {
	ID    int64
	Name  string
	Email string
	State string
}

// AuthService handles authentication business logic
type AuthService interface {
	Login(ctx context.Context, email, password string) (*AuthUser, error)
	GetUserByID(ctx context.Context, id int64) (*AuthUser, error)
}

type authService struct {
	authRepo repo.AuthRepo
}

// NewAuthService creates a new AuthService
func NewAuthService(authRepo repo.AuthRepo) AuthService {
	return &authService{authRepo: authRepo}
}

func (s *authService) Login(ctx context.Context, email, password string) (*AuthUser, error) {
	user, err := s.authRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// TODO: Verify password hash
	if user.Password != password {
		return nil, ErrInvalidCredentials
	}

	if user.State != "active" {
		return nil, ErrUserInactive
	}

	return &AuthUser{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		State: user.State,
	}, nil
}

func (s *authService) GetUserByID(ctx context.Context, id int64) (*AuthUser, error) {
	user, err := s.authRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &AuthUser{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		State: user.State,
	}, nil
}
