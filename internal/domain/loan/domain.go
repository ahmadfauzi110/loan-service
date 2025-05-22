package loan

import "time"

const (
	PROPOSED  = "proposed"
	APPROVED  = "approved"
	INVESTED  = "invested"
	DISBURSED = "disbursed"

	USER_TYPE_CUSTOMER = "customer"
	USER_TYPE_EMPLOYEE = "employee"
)

var AllowedExts = map[string]bool{
	".pdf":  true,
	".jpg":  true,
	".jpeg": true,
}

type Loan struct {
	ID                       *int
	BorrowerID               int
	PrincipalAmount          int
	Rate                     int
	RequestDate              *time.Time
	Status                   string
	TotalInvested            int
	ApprovedBy               *int
	ApprovedAt               *time.Time
	ApprovedPicture          *string
	DisbursedBy              *int
	DisbursedAt              *time.Time
	DisbursedAggrementLetter *string
}

type User struct {
	ID   *int
	Name string
}

type LoanList struct {
	ID                       *int
	Borrower                 *User
	PrincipalAmount          int
	Rate                     int
	RequestDate              *time.Time
	Status                   string
	TotalInvested            int
	Approver                 *User
	ApprovedAt               *time.Time
	ApprovedPicture          *string
	Disburser                *User
	DisbursedAt              *time.Time
	DisbursedAggrementLetter *string
}

type CreateLoan struct {
	BorrowerID      int `json:"borrower_id" validate:"required"`
	PrincipalAmount int `json:"principal_amount" validate:"required"`
	Rate            int `json:"rate" validate:"required"`
}

type ApproveLoan struct {
	EmployeeID int `form:"employee_id" validate:"required"`
	LoanID     int `param:"id" validate:"required"`
}

type DisburseLoan struct {
	EmployeeID int `form:"employee_id" validate:"required"`
	LoanID     int `param:"id" validate:"required"`
}

type FilterLoan struct {
	Status      string `query:"status" json:"status"`
	BorrowerID  string `query:"borrower_id" json:"borrower_id"`
	ApprovedBy  string `query:"approved_by" json:"approved_by"`
	DisbursedBy string `query:"disbursed_by" json:"disbursed_by"`
}
