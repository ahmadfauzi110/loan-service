package user

import domain "github.com/ahmadfauzi110/loan-service/internal/domain/user"

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByID(id int) (*domain.User, error)
	ListUsers(domain.FilterUser) ([]domain.User, error)
}
