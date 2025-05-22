package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/ahmadfauzi110/loan-service/config"
	userHandler "github.com/ahmadfauzi110/loan-service/internal/adapter/handler/user"
	userRepository "github.com/ahmadfauzi110/loan-service/internal/infrastructure/repository/user"
	userService "github.com/ahmadfauzi110/loan-service/internal/service/user"

	investmentHandler "github.com/ahmadfauzi110/loan-service/internal/adapter/handler/investment"
	investmentRepository "github.com/ahmadfauzi110/loan-service/internal/infrastructure/repository/investment"
	investmentService "github.com/ahmadfauzi110/loan-service/internal/service/investment"

	loanHandler "github.com/ahmadfauzi110/loan-service/internal/adapter/handler/loan"
	loanRepository "github.com/ahmadfauzi110/loan-service/internal/infrastructure/repository/loan"
	loanService "github.com/ahmadfauzi110/loan-service/internal/service/loan"

	emailSender "github.com/ahmadfauzi110/loan-service/internal/infrastructure/email"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {

	// Initialize services
	userRepo := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(userRepo)
	userHandler := userHandler.NewUserHandler(userService)

	email := emailSender.NewBrevoEmailService(config.CurrentConfig.BREVO)

	loanRepo := loanRepository.NewLoanRepository(db)
	loanService := loanService.NewLoanService(loanRepo, userRepo)
	loanHandler := loanHandler.NewLoanHandler(loanService)

	investmentRepo := investmentRepository.NewInvestmentRepository(db)
	investmentService := investmentService.NewInvestmentService(investmentRepo, loanRepo, userRepo, email)
	investmentHandler := investmentHandler.NewInvestmentHandler(investmentService)

	// Initialize routes

	apiV1 := e.Group("api/v1")

	user := apiV1.Group("/users")
	user.GET("", userHandler.ListUsers)
	user.POST("", userHandler.CreateUser)

	loan := apiV1.Group("/loans")
	loan.GET("", loanHandler.ListLoans)
	loan.POST("", loanHandler.CreateLoan)
	loan.POST("/:id/approve", loanHandler.ApproveLoan)
	loan.POST("/:id/disburse", loanHandler.DisburseLoan)

	investment := apiV1.Group("/investments")
	investment.GET("", investmentHandler.ListInvestments)
	investment.POST("", investmentHandler.CreateInvestment)
}
