package investment

import (
	"time"

	loanRepo "github.com/ahmadfauzi110/loan-service/internal/infrastructure/repository/loan"
	userRepo "github.com/ahmadfauzi110/loan-service/internal/infrastructure/repository/user"
)

type Model struct {
	ID              *int            `gorm:"column:id;primaryKey;not null"`
	LoanID          int             `gorm:"column:loan_id;not null"`
	Loan            *loanRepo.Model `gorm:"foreignKey:LoanID;references:ID"`
	InvestorID      int             `gorm:"column:investor_id;not null"`
	Investor        *userRepo.Model `gorm:"foreignKey:InvestorID;references:ID"`
	Amount          int             `gorm:"column:amount;not null"`
	Roi             int             `gorm:"column:rate;not null"`
	AggrementLetter *string         `gorm:"column:aggrement_letter;size:255;null"`
	Date            *time.Time      `gorm:"column:rate;not null"`
}

func (Model) TableName() string {
	return "investments"
}
