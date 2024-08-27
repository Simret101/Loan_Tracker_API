package domain

import (
	"github.com/gin-gonic/gin"
)

type Loan_Controller_interface interface {
	ApplyForLoan() gin.HandlerFunc
	ViewLoanStatus() gin.HandlerFunc
	ViewAllLoans() gin.HandlerFunc
	ApproveRejectLoan() gin.HandlerFunc
	DeleteLoan() gin.HandlerFunc
	ViewSystemLogs() gin.HandlerFunc
}

type Loan_Usecase_interface interface {
	CreateLoan(loan CreateLoan) (Loan, error)
	GetOneLoan(id string) (Loan, error)
	GetAllLoans(page int, limit int, status string, order string) ([]Loan, error)
	UpdateLoan(id string, loan Loan) (Loan, error)
	UpdateLoanStatus(id string, status string) error
	DeleteLoan(id string) error
	FilterLoan(filters map[string]interface{}) ([]Loan, error)
	GetUniqueLoan(filter map[string]interface{}, posts *[]Loan) error
	GetSystemLogs() ([]ResponseLoan, error)
}

type Loan_Repository_interface interface {
	CreateLoanDocument(loan Loan) (Loan, error)
	GetOneLoanDocument(id string) (Loan, error)
	GetLoanDocuments(page int, limit int) ([]Loan, error)
	UpdateLoanDocument(id string, loan Loan) (Loan, error)
	DeleteLoanDocument(id string) error
	FilterLoanDocument(filters map[string]interface{}) ([]Loan, error)
	UpdateLoanStatus(id string, status string) error
	GetSystemLogs() ([]ResponseLoan, error)
}
