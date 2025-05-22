package user

import (
	"net/http"

	port "github.com/ahmadfauzi110/loan-service/internal/port/user"
	"github.com/ahmadfauzi110/loan-service/util"
	"github.com/labstack/echo/v4"

	domain "github.com/ahmadfauzi110/loan-service/internal/domain/user"
)

type userHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) port.UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) CreateUser(c echo.Context) error {
	var user domain.CreateUser
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "validation error",
		})
	}

	result := h.userService.CreateUser(&user)

	return c.JSON(http.StatusCreated, util.Response{
		Data:    result,
		Message: "user created successfully",
	})
}

func (h *userHandler) ListUsers(c echo.Context) error {
	var filter domain.FilterUser
	if err := c.Bind(filter); err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}

	result, err := h.userService.ListUsers(&filter)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}

	return c.JSON(http.StatusCreated, util.Response{
		Data:    result,
		Message: "user created successfully",
	})
}
