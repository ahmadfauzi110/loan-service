package loan

import (
	port "github.com/ahmadfauzi110/loan-service/internal/port/loan"
	"gorm.io/gorm"

	domain "github.com/ahmadfauzi110/loan-service/internal/domain/loan"
)

type loanRepository struct {
	DB *gorm.DB
}

func NewLoanRepository(db *gorm.DB) port.LoanRepository {
	return &loanRepository{
		DB: db,
	}
}

func (r *loanRepository) CreateLoan(loan *domain.Loan) (*int, error) {
	model := Model{
		BorrowerID:      loan.BorrowerID,
		PrincipalAmount: loan.PrincipalAmount,
		Rate:            loan.Rate,
		RequestDate:     loan.RequestDate,
		Status:          loan.Status,
	}

	if err := r.DB.Create(&model).Error; err != nil {
		return nil, err
	}

	return model.ID, nil
}

func (r *loanRepository) UpdateLoan(loan *domain.Loan) error {
	model := Model{
		ID:              loan.ID,
		BorrowerID:      loan.BorrowerID,
		PrincipalAmount: loan.PrincipalAmount,
		Rate:            loan.Rate,
		RequestDate:     loan.RequestDate,
		Status:          loan.Status,
		TotalInvested:   loan.TotalInvested,
	}

	if err := r.DB.Updates(&model).Error; err != nil {
		return err
	}

	return nil
}

func (r *loanRepository) ApproveLoan(loan *domain.Loan) (*domain.Loan, error) {
	model := Model{
		ID:              loan.ID,
		ApprovedAt:      loan.ApprovedAt,
		ApprovedBy:      loan.ApprovedBy,
		ApprovedPicture: loan.ApprovedPicture,
		Status:          loan.Status,
	}

	if err := r.DB.Updates(&model).Error; err != nil {
		return nil, err
	}

	return loan, nil
}

func (r *loanRepository) DisburseLoan(loan *domain.Loan) (*domain.Loan, error) {
	model := Model{
		ID:                       loan.ID,
		DisbursedAt:              loan.DisbursedAt,
		DisbursedBy:              loan.DisbursedBy,
		DisbursedAggrementLetter: loan.DisbursedAggrementLetter,
		Status:                   loan.Status,
	}

	if err := r.DB.Updates(&model).Error; err != nil {
		return nil, err
	}

	return loan, nil
}

func (r *loanRepository) GetLoanByID(id int) (*domain.Loan, error) {
	loan := Model{}

	if err := r.DB.First(&loan, id).Error; err != nil {
		return nil, err
	}

	return &domain.Loan{
		ID:                       loan.ID,
		BorrowerID:               loan.BorrowerID,
		PrincipalAmount:          loan.PrincipalAmount,
		Rate:                     loan.Rate,
		RequestDate:              loan.RequestDate,
		Status:                   loan.Status,
		TotalInvested:            loan.TotalInvested,
		ApprovedBy:               loan.ApprovedBy,
		ApprovedAt:               loan.ApprovedAt,
		ApprovedPicture:          loan.ApprovedPicture,
		DisbursedBy:              loan.DisbursedBy,
		DisbursedAt:              loan.DisbursedAt,
		DisbursedAggrementLetter: loan.DisbursedAggrementLetter,
	}, nil
}

func (r *loanRepository) ListLoans(filter domain.FilterLoan) ([]domain.LoanList, error) {
	var models []Model

	if filter.Status != "" {
		r.DB.Where("status = ?", filter.Status)
	}

	if filter.DisbursedBy != "" {
		r.DB.Where("disbursed_by = ?", filter.DisbursedBy)
	}

	if filter.ApprovedBy != "" {
		r.DB.Where("approved_by = ?", filter.ApprovedBy)
	}

	if filter.BorrowerID != "" {
		r.DB.Where("borrower_id = ?", filter.BorrowerID)
	}

	if err := r.DB.Preload("Borrower").
		Preload("Approver").
		Preload("Disburser").
		Find(&models).Error; err != nil {
		return nil, err
	}

	var result []domain.LoanList
	for _, v := range models {
		var borrower domain.User
		var approver domain.User
		var disburser domain.User
		if v.Borrower != nil {
			borrower = domain.User{
				ID:   v.Borrower.ID,
				Name: v.Borrower.Name,
			}
		}

		if v.Approver != nil {
			approver = domain.User{
				ID:   v.Approver.ID,
				Name: v.Approver.Name,
			}
		}

		if v.Disburser != nil {
			disburser = domain.User{
				ID:   v.Disburser.ID,
				Name: v.Disburser.Name,
			}
		}

		result = append(result, domain.LoanList{
			ID:                       v.ID,
			Borrower:                 &borrower,
			RequestDate:              v.RequestDate,
			Status:                   v.Status,
			PrincipalAmount:          v.PrincipalAmount,
			Rate:                     v.Rate,
			TotalInvested:            v.TotalInvested,
			Approver:                 &approver,
			ApprovedAt:               v.ApprovedAt,
			ApprovedPicture:          v.ApprovedPicture,
			Disburser:                &disburser,
			DisbursedAt:              v.DisbursedAt,
			DisbursedAggrementLetter: v.DisbursedAggrementLetter,
		})
	}

	return result, nil
}
