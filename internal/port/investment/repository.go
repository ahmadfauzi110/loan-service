package investment

import domain "github.com/ahmadfauzi110/loan-service/internal/domain/investment"

type InvestmentRepository interface {
	CreateInvestment(investment *domain.Investment) (*int, error)
	UpdateInvestment(investment *domain.Investment) error
	GetAllInvestmentByLoanID(loanID int) ([]*domain.Investment, error)
	ListInvestments(filter *domain.FilterInvestment) ([]domain.Investment, error)
}
