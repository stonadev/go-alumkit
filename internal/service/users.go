package service

import (
	"context"
	"time"

	"github.com/stonadev/alumkit/internal/repo"
)

// User represents a user in the service layer
type User struct {
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

// UserList represents a paginated list of users
type UserList struct {
	Users      []User
	Total      int64
	Page       int
	PerPage    int
	TotalPages int
}

// UserService handles user business logic
type UserService interface {
	List(ctx context.Context, page, perPage int) (*UserList, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
}

type userService struct {
	userRepo repo.UserRepo
}

// NewUserService creates a new UserService
func NewUserService(userRepo repo.UserRepo) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) List(ctx context.Context, page, perPage int) (*UserList, error) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 20
	}

	offset := (page - 1) * perPage
	users, total, err := s.userRepo.List(ctx, perPage, offset)
	if err != nil {
		return nil, err
	}

	result := make([]User, len(users))
	for i, u := range users {
		result[i] = User{
			ID:               u.ID,
			Name:             u.Name,
			Email:            u.Email,
			State:            u.State,
			PhotoPath:        u.PhotoPath,
			DateOfBirth:      u.DateOfBirth,
			Gender:           u.Gender,
			BloodGroup:       u.BloodGroup,
			PresentAddress:   u.PresentAddress,
			PermanentAddress: u.PermanentAddress,
			CreatedAt:        u.CreatedAt,
		}
	}

	totalPages := int(total) / perPage
	if int(total)%perPage > 0 {
		totalPages++
	}

	return &UserList{
		Users:      result,
		Total:      total,
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
	}, nil
}

func (s *userService) GetByID(ctx context.Context, id int64) (*User, error) {
	u, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:               u.ID,
		Name:             u.Name,
		Email:            u.Email,
		State:            u.State,
		PhotoPath:        u.PhotoPath,
		DateOfBirth:      u.DateOfBirth,
		Gender:           u.Gender,
		BloodGroup:       u.BloodGroup,
		PresentAddress:   u.PresentAddress,
		PermanentAddress: u.PermanentAddress,
		CreatedAt:        u.CreatedAt,
	}, nil
}

func (s *userService) Create(ctx context.Context, user *User) error {
	return s.userRepo.Create(ctx, &repo.UserWithProfile{
		Name:             user.Name,
		Email:            user.Email,
		State:            user.State,
		PhotoPath:        user.PhotoPath,
		DateOfBirth:      user.DateOfBirth,
		Gender:           user.Gender,
		BloodGroup:       user.BloodGroup,
		PresentAddress:   user.PresentAddress,
		PermanentAddress: user.PermanentAddress,
	})
}

func (s *userService) Update(ctx context.Context, user *User) error {
	return s.userRepo.Update(ctx, &repo.UserWithProfile{
		ID:               user.ID,
		Name:             user.Name,
		Email:            user.Email,
		State:            user.State,
		PhotoPath:        user.PhotoPath,
		DateOfBirth:      user.DateOfBirth,
		Gender:           user.Gender,
		BloodGroup:       user.BloodGroup,
		PresentAddress:   user.PresentAddress,
		PermanentAddress: user.PermanentAddress,
	})
}

func (s *userService) Delete(ctx context.Context, id int64) error {
	return s.userRepo.Delete(ctx, id)
}
