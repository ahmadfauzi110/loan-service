package investment

import (
	"time"

	loanDomain "github.com/ahmadfauzi110/loan-service/internal/domain/loan"
	userDomain "github.com/ahmadfauzi110/loan-service/internal/domain/user"
)

type Investment struct {
	ID              *int             `json:"id"`
	LoanID          int              `json:"loan_id"`
	Loan            *loanDomain.Loan `json:"loan"`
	InvestorID      int              `json:"investor_id"`
	Investor        *userDomain.User `json:"investor"`
	Amount          int              `json:"amount"`
	Roi             int              `json:"rate"`
	AggrementLetter *string          `json:"aggrement_letter"`
	Date            *time.Time       `json:"date"`
}

type CreateInvestment struct {
	LoanID     int `json:"loan_id" validate:"required"`
	InvestorID int `json:"investor_id" validate:"required"`
	Amount     int `json:"amount" validate:"required"`
}

var AggrementLetterBody = "Dear Mr/Ms%s, Below are the details of your investment:\nLoan ID : %d\nAmount : %d\nAggreement letter : %s \nThank you"

type FilterInvestment struct {
	InvestorID *int `json:"investor_id" query:"investor_id" validate:"required"`
}
