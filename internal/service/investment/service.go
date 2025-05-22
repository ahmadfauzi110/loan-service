package investment

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ahmadfauzi110/loan-service/config"
	port "github.com/ahmadfauzi110/loan-service/internal/port/investment"
	loanPort "github.com/ahmadfauzi110/loan-service/internal/port/loan"
	"github.com/jung-kurt/gofpdf"

	domain "github.com/ahmadfauzi110/loan-service/internal/domain/investment"
	loanDomain "github.com/ahmadfauzi110/loan-service/internal/domain/loan"
	emailPort "github.com/ahmadfauzi110/loan-service/internal/port/email"
	userPort "github.com/ahmadfauzi110/loan-service/internal/port/user"
)

type investmentService struct {
	repo     port.InvestmentRepository
	loanRepo loanPort.LoanRepository
	userRepo userPort.UserRepository
	email    emailPort.EmailSender
}

func NewInvestmentService(repo port.InvestmentRepository, loanRepo loanPort.LoanRepository,
	userRepo userPort.UserRepository, email emailPort.EmailSender) port.InvestmentService {
	return &investmentService{
		repo:     repo,
		loanRepo: loanRepo,
		userRepo: userRepo,
		email:    email,
	}
}

func (s *investmentService) CreateInvestment(request *domain.CreateInvestment) (*int, error) {
	loan, err := s.loanRepo.GetLoanByID(request.LoanID)
	if err != nil {
		return nil, err
	}

	if loan == nil {
		return nil, errors.New("loan not found")
	}

	if loan.Status != loanDomain.APPROVED {
		return nil, errors.New("loan not yet approved, cannot create investment")
	}

	if loan.TotalInvested+request.Amount > loan.PrincipalAmount {
		return nil, errors.New("investment amount cannot exceed principal amount")
	}

	investor, err := s.userRepo.GetUserByID(request.InvestorID)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, errors.New("investor not found")
		}
		return nil, err
	}
	if investor == nil {
		return nil, errors.New("investor not found")
	}

	investment := &domain.Investment{
		LoanID:     request.LoanID,
		InvestorID: request.InvestorID,
		Amount:     request.Amount,
		Roi:        loan.Rate,
	}

	id, err := s.repo.CreateInvestment(investment)
	if err != nil {
		return nil, err
	}

	loan.TotalInvested = loan.TotalInvested + request.Amount
	err = s.loanRepo.UpdateLoan(loan)
	if err != nil {
		return nil, err
	}

	if loan.TotalInvested == loan.PrincipalAmount {
		loan.Status = loanDomain.INVESTED
		err = s.loanRepo.UpdateLoan(loan)
		if err != nil {
			return nil, err
		}

		list, err := s.repo.GetAllInvestmentByLoanID(*loan.ID)
		if err != nil {
			return nil, err
		}

		for _, v := range list {
			link, err := s.GenerateInvestmentAggrement(v.LoanID, v.Amount, *v.ID, v.Investor.Name, v.Investor.Email)
			if err != nil {
				return nil, err
			}

			v.AggrementLetter = &link

			err = s.repo.UpdateInvestment(v)
			if err != nil {
				return nil, err
			}

			err = s.email.Send(v.Investor.Email, "Aggrement Letter",
				fmt.Sprintf(domain.AggrementLetterBody, v.Investor.Name, v.LoanID, v.Amount, *v.AggrementLetter))
			if err != nil {
				return nil, err
			}
		}

	}

	return id, nil
}

func (s *investmentService) ListInvestments(filter *domain.FilterInvestment) ([]domain.Investment, error) {
	return s.repo.ListInvestments(filter)
}

func (s *investmentService) GenerateInvestmentAggrement(loanID, amount, investmentID int, name, email string) (string, error) {
	os.MkdirAll("storage/letter", os.ModePerm)
	filename := fmt.Sprintf("investment-aggrement-%d%d.pdf", loanID, investmentID)
	fullPath := filepath.Join("storage/letter", filename)

	loanIDStr := strconv.Itoa(loanID)
	amountStr := strconv.Itoa(amount)
	investmentIDStr := strconv.Itoa(investmentID)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(60, 10, "Aggrement Letter")
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	pdf.Ln(10)
	pdf.Cell(40, 10, "Dear Mr/Ms  "+name+",\nBelow are the details of your investment:")
	pdf.Ln(10)
	pdf.Cell(40, 10, "Loan ID : "+loanIDStr)
	pdf.Ln(10)
	pdf.Cell(40, 10, "Investment ID : "+investmentIDStr)
	pdf.Ln(10)
	pdf.Cell(40, 10, "Amount : "+amountStr)
	pdf.Ln(10)
	pdf.Cell(40, 10, "Thank you")

	err := pdf.OutputFileAndClose(fullPath)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s/%s", config.CurrentConfig.BASE_URL, config.CurrentConfig.STATIC_PATH, filename), nil
}
