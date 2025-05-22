package user

import (
	port "github.com/ahmadfauzi110/loan-service/internal/port/user"

	domain "github.com/ahmadfauzi110/loan-service/internal/domain/user"
)

type userService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) port.UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(request *domain.CreateUser) error {
	user := &domain.User{
		Name:     request.Name,
		Email:    request.Email,
		UserType: request.UserType,
	}

	return s.repo.CreateUser(user)
}

func (s *userService) ListUsers(filter *domain.FilterUser) ([]domain.User, error) {
	return s.repo.ListUsers(*filter)
}
