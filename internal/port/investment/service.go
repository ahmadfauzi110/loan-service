package investment

import domain "github.com/ahmadfauzi110/loan-service/internal/domain/investment"

type InvestmentService interface {
	CreateInvestment(investment *domain.CreateInvestment) (*int, error)
	ListInvestments(filter *domain.FilterInvestment) ([]domain.Investment, error)
}
