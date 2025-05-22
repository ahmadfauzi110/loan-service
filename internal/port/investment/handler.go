package investment

import (
	"github.com/labstack/echo/v4"
)

type InvestmentHandler interface {
	CreateInvestment(c echo.Context) error
	ListInvestments(c echo.Context) error
}
