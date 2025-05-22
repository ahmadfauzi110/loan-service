package user

import (
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	CreateUser(c echo.Context) error
	ListUsers(c echo.Context) error
}
