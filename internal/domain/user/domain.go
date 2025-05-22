package user

type User struct {
	ID       *int   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
}

type CreateUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	UserType string `json:"user_type" validate:"required,oneof=employee customer"`
}

type FilterUser struct {
	Name     string `query:"name" json:"name"`
	Email    string `query:"email" json:"email"`
	UserType string `query:"user_type" json:"user_type"`
}
