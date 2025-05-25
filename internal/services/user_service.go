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

func (s *UserService) ListUser() ([]models.User, error) {
	return s.repo.List()
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
	if id == 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.repo.GetByID(id)
}
