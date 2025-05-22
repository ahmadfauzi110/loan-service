package user

import (
	port "github.com/ahmadfauzi110/loan-service/internal/port/user"
	"gorm.io/gorm"

	domain "github.com/ahmadfauzi110/loan-service/internal/domain/user"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) CreateUser(user *domain.User) error {
	model := Model{
		Name:     user.Name,
		Email:    user.Email,
		UserType: user.UserType,
	}

	return r.DB.Create(&model).Error
}

func (r *userRepository) GetUserByID(id int) (*domain.User, error) {
	var model Model

	if err := r.DB.First(&model, id).Error; err != nil {
		return nil, err
	}

	return &domain.User{
		ID:       model.ID,
		Name:     model.Name,
		Email:    model.Email,
		UserType: model.UserType,
	}, nil
}

func (r *userRepository) ListUsers(filter domain.FilterUser) ([]domain.User, error) {
	var models []Model

	if filter.Email != "" {
		r.DB.Where("email = ?", filter.Email)
	}

	if filter.Name != "" {
		r.DB.Where("name = ?", filter.Name)
	}

	if filter.UserType != "" {
		r.DB.Where("user_type = ?", filter.UserType)
	}

	if err := r.DB.Find(&models).Error; err != nil {
		return nil, err
	}

	var result []domain.User
	for _, v := range models {
		result = append(result, domain.User{
			ID:       v.ID,
			Name:     v.Name,
			Email:    v.Email,
			UserType: v.UserType,
		})
	}

	return result, nil
}
