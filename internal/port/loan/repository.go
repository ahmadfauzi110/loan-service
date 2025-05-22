package loan

import domain "github.com/ahmadfauzi110/loan-service/internal/domain/loan"

type LoanRepository interface {
	CreateLoan(loan *domain.Loan) (*int, error)
	UpdateLoan(loan *domain.Loan) error
	ApproveLoan(loan *domain.Loan) (*domain.Loan, error)
	DisburseLoan(loan *domain.Loan) (*domain.Loan, error)
	GetLoanByID(id int) (*domain.Loan, error)
	ListLoans(filter domain.FilterLoan) ([]domain.LoanList, error)
}
