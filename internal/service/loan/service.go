package loan

import (
	"errors"
	"time"

	port "github.com/ahmadfauzi110/loan-service/internal/port/loan"

	domain "github.com/ahmadfauzi110/loan-service/internal/domain/loan"

	userPort "github.com/ahmadfauzi110/loan-service/internal/port/user"
)

type loanService struct {
	repo     port.LoanRepository
	userRepo userPort.UserRepository
}

func NewLoanService(repo port.LoanRepository, userRepo userPort.UserRepository) port.LoanService {
	return &loanService{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (s *loanService) CreateLoan(request *domain.CreateLoan) (*int, error) {

	borrower, err := s.userRepo.GetUserByID(request.BorrowerID)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, errors.New("borrower not found")
		}
		return nil, err
	}
	if borrower == nil {
		return nil, errors.New("borrower not found")
	}

	date := time.Now()
	loan := domain.Loan{
		BorrowerID:      request.BorrowerID,
		PrincipalAmount: request.PrincipalAmount,
		Rate:            request.Rate,
		RequestDate:     &date,
		Status:          domain.PROPOSED,
	}

	id, err := s.repo.CreateLoan(&loan)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *loanService) ApproveLoan(request *domain.ApproveLoan, fileUrl string) error {
	loan, err := s.repo.GetLoanByID(request.LoanID)
	if err != nil {
		return err
	}

	if loan == nil {
		return errors.New("loan not found")
	}

	if loan.Status != domain.PROPOSED {
		return errors.New("loan already approved")
	}

	employee, err := s.userRepo.GetUserByID(request.EmployeeID)
	if err != nil {
		if err.Error() == "record not found" {
			return errors.New("employee not found")
		}
		return err
	}
	if employee == nil {
		return errors.New("employee not found")
	}

	if employee.UserType != domain.USER_TYPE_EMPLOYEE {
		return errors.New("only employee can approve loan")
	}

	today := time.Now()

	loan.Status = domain.APPROVED
	loan.ApprovedAt = &today
	loan.ApprovedBy = &request.EmployeeID
	loan.ApprovedPicture = &fileUrl

	_, err = s.repo.ApproveLoan(loan)
	if err != nil {
		return err
	}

	return nil
}

func (s *loanService) DisburseLoan(request *domain.DisburseLoan, fileUrl string) error {
	loan, err := s.repo.GetLoanByID(request.LoanID)
	if err != nil {
		return err
	}

	if loan == nil {
		return errors.New("loan not found")
	}

	if loan.Status != domain.INVESTED {
		return errors.New("only invested loan can be disbursed")
	}

	employee, err := s.userRepo.GetUserByID(request.EmployeeID)
	if err != nil {
		if err.Error() == "record not found" {
			return errors.New("employee not found")
		}
		return err
	}
	if employee == nil {
		return errors.New("employee not found")
	}

	if employee.UserType != domain.USER_TYPE_EMPLOYEE {
		return errors.New("only employee can disburse loan")
	}

	today := time.Now()

	loan.Status = domain.DISBURSED
	loan.DisbursedAt = &today
	loan.DisbursedBy = &request.EmployeeID
	loan.DisbursedAggrementLetter = &fileUrl

	_, err = s.repo.DisburseLoan(loan)
	if err != nil {
		return err
	}

	return nil
}

func (s *loanService) ListLoans(filter *domain.FilterLoan) ([]domain.LoanList, error) {
	return s.repo.ListLoans(*filter)
}
