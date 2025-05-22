package investment

import (
	port "github.com/ahmadfauzi110/loan-service/internal/port/investment"
	"gorm.io/gorm"

	domain "github.com/ahmadfauzi110/loan-service/internal/domain/investment"

	loanDomain "github.com/ahmadfauzi110/loan-service/internal/domain/loan"
	userDomain "github.com/ahmadfauzi110/loan-service/internal/domain/user"
)

type investmentRepository struct {
	DB *gorm.DB
}

func NewInvestmentRepository(db *gorm.DB) port.InvestmentRepository {
	return &investmentRepository{
		DB: db,
	}
}

func (r *investmentRepository) CreateInvestment(investment *domain.Investment) (*int, error) {
	model := Model{
		LoanID:     investment.LoanID,
		InvestorID: investment.InvestorID,
		Amount:     investment.Amount,
		Roi:        investment.Roi,
		Date:       investment.Date,
	}

	if err := r.DB.Create(&model).Error; err != nil {
		return nil, err
	}

	return model.ID, nil
}

func (r *investmentRepository) UpdateInvestment(investment *domain.Investment) error {
	model := Model{
		ID:              investment.ID,
		LoanID:          investment.LoanID,
		InvestorID:      investment.InvestorID,
		Amount:          investment.Amount,
		Roi:             investment.Roi,
		Date:            investment.Date,
		AggrementLetter: investment.AggrementLetter,
	}

	return r.DB.Save(&model).Error
}

func (r *investmentRepository) GetAllInvestmentByLoanID(loanID int) ([]*domain.Investment, error) {

	model := []Model{}
	if err := r.DB.Where("loan_id = ?", loanID).
		Preload("Loan").
		Preload("Investor").
		Find(&model).Error; err != nil {
		return nil, err
	}

	var result []*domain.Investment
	for _, v := range model {
		var investor userDomain.User
		var loan loanDomain.Loan
		if v.Investor != nil {
			investor = userDomain.User{
				ID:    v.Investor.ID,
				Name:  v.Investor.Name,
				Email: v.Investor.Email,
			}
		}

		if v.Loan != nil {
			loan = loanDomain.Loan{
				ID:              v.Loan.ID,
				BorrowerID:      v.Loan.BorrowerID,
				PrincipalAmount: v.Loan.PrincipalAmount,
				Rate:            v.Loan.Rate,
				RequestDate:     v.Loan.RequestDate,
				Status:          v.Loan.Status,
				TotalInvested:   v.Loan.TotalInvested,
			}
		}

		investment := domain.Investment{
			ID:         v.ID,
			LoanID:     v.LoanID,
			Loan:       &loan,
			InvestorID: v.InvestorID,
			Investor:   &investor,
			Amount:     v.Amount,
			Roi:        v.Roi,
			Date:       v.Date,
		}

		result = append(result, &investment)
	}

	return result, nil
}

func (r *investmentRepository) ListInvestments(filter *domain.FilterInvestment) ([]domain.Investment, error) {
	var models []Model

	if filter.InvestorID != nil {
		r.DB.Where("investor_id = ?", filter.InvestorID)
	}

	if err := r.DB.Preload("Loan").Preload("Investor").Find(&models).Error; err != nil {
		return nil, err
	}

	var result []domain.Investment
	for _, v := range models {
		var investor userDomain.User
		var loan loanDomain.Loan
		if v.Investor != nil {
			investor = userDomain.User{
				ID:    v.Investor.ID,
				Name:  v.Investor.Name,
				Email: v.Investor.Email,
			}
		}

		if v.Loan != nil {
			loan = loanDomain.Loan{
				ID:              v.Loan.ID,
				BorrowerID:      v.Loan.BorrowerID,
				PrincipalAmount: v.Loan.PrincipalAmount,
				Rate:            v.Loan.Rate,
				RequestDate:     v.Loan.RequestDate,
				Status:          v.Loan.Status,
				TotalInvested:   v.Loan.TotalInvested,
			}
		}

		investment := domain.Investment{
			ID:         v.ID,
			LoanID:     v.LoanID,
			Loan:       &loan,
			InvestorID: v.InvestorID,
			Investor:   &investor,
			Amount:     v.Amount,
			Roi:        v.Roi,
			Date:       v.Date,
		}

		result = append(result, investment)
	}

	return result, nil
}
