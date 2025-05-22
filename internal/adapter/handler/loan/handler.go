package loan

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ahmadfauzi110/loan-service/config"
	port "github.com/ahmadfauzi110/loan-service/internal/port/loan"
	"github.com/ahmadfauzi110/loan-service/util"
	"github.com/labstack/echo/v4"

	domain "github.com/ahmadfauzi110/loan-service/internal/domain/loan"
)

type loanHandler struct {
	loanService port.LoanService
}

func NewLoanHandler(loanService port.LoanService) port.LoanHandler {
	return &loanHandler{
		loanService: loanService,
	}
}

func (h *loanHandler) CreateLoan(c echo.Context) error {
	var loan domain.CreateLoan
	if err := c.Bind(&loan); err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}

	if err := c.Validate(loan); err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "validation error",
		})
	}

	result, err := h.loanService.CreateLoan(&loan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}

	return c.JSON(http.StatusCreated, util.Response{
		Data:    result,
		Message: "loan created successfully",
	})
}

func (h *loanHandler) ApproveLoan(c echo.Context) error {

	var request domain.ApproveLoan
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "validation error",
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   "File not provided",
			Message: "bad request",
		})
	}

	// Validate extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !domain.AllowedExts[ext] {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   "Only PDF and JPG files allowed",
			Message: "bad request",
		})
	}

	// Open and save file
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}
	defer src.Close()

	saveDir := "storage/letter"
	os.MkdirAll(saveDir, os.ModePerm)

	filename := fmt.Sprintf("approval-proof-%d%s", request.LoanID, ext)

	dstPath := filepath.Join(saveDir, filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}

	fileURL := fmt.Sprintf("%s/%s/%s", config.CurrentConfig.BASE_URL, config.CurrentConfig.STATIC_PATH, filename)

	err = h.loanService.ApproveLoan(&request, fileURL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}

	return c.JSON(http.StatusOK, util.Response{
		Data:    nil,
		Message: "success",
	})
}

func (h *loanHandler) DisburseLoan(c echo.Context) error {
	var request domain.DisburseLoan
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "validation error",
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   "File not provided",
			Message: "bad request",
		})
	}

	// Validate extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !domain.AllowedExts[ext] {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   "Only PDF and JPG files allowed",
			Message: "bad request",
		})
	}

	// Open and save file
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}
	defer src.Close()

	saveDir := "storage/letter"
	os.MkdirAll(saveDir, os.ModePerm)

	filename := fmt.Sprintf("disburse-aggrement-%d%s", request.LoanID, ext)

	dstPath := filepath.Join(saveDir, filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}

	fileURL := fmt.Sprintf("%s/%s/%s", config.CurrentConfig.BASE_URL, config.CurrentConfig.STATIC_PATH, filename)

	err = h.loanService.DisburseLoan(&request, fileURL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}

	return c.JSON(http.StatusOK, util.Response{
		Data:    nil,
		Message: "success",
	})
}

func (h *loanHandler) ListLoans(c echo.Context) error {
	var filter domain.FilterLoan
	if err := c.Bind(filter); err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}

	result, err := h.loanService.ListLoans(&filter)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Error:   err.Error(),
			Message: "bad request",
		})
	}

	return c.JSON(http.StatusOK, util.Response{
		Data:    result,
		Message: "success",
	})
}
