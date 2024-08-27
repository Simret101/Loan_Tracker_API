package controller

import (
	"Loan_Tracker_API/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanController struct {
	LoanUsecase domain.Loan_Usecase_interface
	UserUsecase domain.User_Usecase_interface
}

func NewLoanController(loanUsecase domain.Loan_Usecase_interface, userUsecase domain.User_Usecase_interface) *LoanController {
	return &LoanController{
		LoanUsecase: loanUsecase,
		UserUsecase: userUsecase,
	}
}

func (lc *LoanController) ApplyForLoan() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loan domain.CreateLoan
		if err := c.BindJSON(&loan); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request: " + err.Error()})
			return
		}

		createdLoan, err := lc.LoanUsecase.CreateLoan(loan)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to apply for loan: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Loan application submitted successfully!", "loan": createdLoan})
	}
}

func (lc *LoanController) ViewLoanStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format."})
			return
		}

		loan, err := lc.LoanUsecase.GetOneLoan(id.Hex())
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"loan": loan})
	}
}

func (lc *LoanController) ViewAllLoans() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil || limit < 1 {
			limit = 5
		}

		status := c.Query("status")
		order := c.Query("order")

		loans, err := lc.LoanUsecase.GetAllLoans(page, limit, status, order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": loans,
			"meta": gin.H{
				"limit": limit,
				"page":  page,
			},
		})
	}
}

func (lc *LoanController) ApproveRejectLoan() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		status := c.Query("status")

		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format."})
			return
		}

		err = lc.LoanUsecase.UpdateLoanStatus(id.Hex(), status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update loan status: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Loan status updated successfully!"})
	}
}

func (lc *LoanController) DeleteLoan() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format."})
			return
		}

		err = lc.LoanUsecase.DeleteLoan(id.Hex())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete loan: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Loan deleted successfully!"})
	}
}

func (lc *LoanController) ViewSystemLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Fetch logs from the use case layer
		logs, err := lc.LoanUsecase.GetSystemLogs()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve logs: " + err.Error()})
			return
		}

		// Respond with the logs
		c.JSON(http.StatusOK, gin.H{"logs": logs})
	}
}
