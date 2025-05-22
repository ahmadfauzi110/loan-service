package loan

import (
	"github.com/labstack/echo/v4"
)

type LoanHandler interface {
	CreateLoan(c echo.Context) error
	ApproveLoan(c echo.Context) error
	DisburseLoan(c echo.Context) error
	ListLoans(c echo.Context) error
}
