package user

import domain "github.com/ahmadfauzi110/loan-service/internal/domain/user"

type UserService interface {
	CreateUser(user *domain.CreateUser) error
	ListUsers(filter *domain.FilterUser) ([]domain.User, error)
}
