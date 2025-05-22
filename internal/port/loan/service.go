package loan

import domain "github.com/ahmadfauzi110/loan-service/internal/domain/loan"

type LoanService interface {
	CreateLoan(loan *domain.CreateLoan) (*int, error)
	ApproveLoan(loan *domain.ApproveLoan, fileUrl string) error
	DisburseLoan(loan *domain.DisburseLoan, fileUrl string) error
	ListLoans(filter *domain.FilterLoan) ([]domain.LoanList, error)
}
