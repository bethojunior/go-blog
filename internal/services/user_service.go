package services

import (
	"errors"

	"github.com/bethojunior/go-blog/internal/models"
	"github.com/bethojunior/go-blog/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(user *models.User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}

	return s.repo.Create(user)
}

type ListUserResponse struct {
	Users []models.User `json:"users"`
	Total int64         `json:"total"`
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
}

func (s *UserService) ListUser(page, limit int) (*ListUserResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	users, total, err := s.repo.List(page, limit)
	if err != nil {
		return nil, err
	}

	return &ListUserResponse{
		Users: users,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
	if id == 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.repo.GetByID(id)
}
