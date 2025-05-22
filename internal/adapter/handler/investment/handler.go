package investment

import (
	"net/http"

	port "github.com/ahmadfauzi110/loan-service/internal/port/investment"
	"github.com/ahmadfauzi110/loan-service/util"
	"github.com/labstack/echo/v4"

	domain "github.com/ahmadfauzi110/loan-service/internal/domain/investment"
)

type investmentHandler struct {
	investmentService port.InvestmentService
}

func NewInvestmentHandler(investmentService port.InvestmentService) port.InvestmentHandler {
	return &investmentHandler{
		investmentService: investmentService,
	}
}

func (h *investmentHandler) CreateInvestment(c echo.Context) error {
	var investment domain.CreateInvestment
	if err := c.Bind(&investment); err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}

	if err := c.Validate(investment); err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "validation error",
		})
	}

	result, err := h.investmentService.CreateInvestment(&investment)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}

	return c.JSON(http.StatusCreated, util.Response{
		Data:    result,
		Message: "investment created successfully",
	})
}

func (h *investmentHandler) ListInvestments(c echo.Context) error {
	var filter domain.FilterInvestment
	if err := c.Bind(&filter); err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}

	if err := c.Validate(filter); err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "validation error",
		})
	}

	result, err := h.investmentService.ListInvestments(&filter)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}

	return c.JSON(http.StatusOK, util.Response{
		Data:    result,
		Message: "success",
	})
}
