package loan

import (
	"time"

	userRepo "github.com/ahmadfauzi110/loan-service/internal/infrastructure/repository/user"
)

type Model struct {
	ID                       *int            `gorm:"column:id;primaryKey;autoIncrement;not null"`
	BorrowerID               int             `gorm:"column:borrower_id;not null"`
	Borrower                 *userRepo.Model `gorm:"foreignKey:BorrowerID;references:ID"`
	PrincipalAmount          int             `gorm:"column:principal_amount;not null"`
	Rate                     int             `gorm:"column:rate;not null"`
	RequestDate              *time.Time      `gorm:"column:date;not null"`
	Status                   string          `gorm:"column:status;not null"`
	TotalInvested            int             `gorm:"column:total_invested;not null"`
	ApprovedBy               *int            `gorm:"column:approved_by;null"`
	Approver                 *userRepo.Model `gorm:"foreignKey:ApprovedBy;references:ID"`
	ApprovedAt               *time.Time      `gorm:"column:approved_at;null"`
	ApprovedPicture          *string         `gorm:"column:approved_picture;size:255;null"`
	DisbursedBy              *int            `gorm:"column:disbursed_by;null"`
	Disburser                *userRepo.Model `gorm:"foreignKey:DisbursedBy;references:ID"`
	DisbursedAt              *time.Time      `gorm:"column:disbursed_at;null"`
	DisbursedAggrementLetter *string         `gorm:"column:disbursed_aggrement_letter;size:255;null"`
}

func (Model) TableName() string {
	return "loans"
}
