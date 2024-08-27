package router

import (
	"Loan_Tracker_API/delivery/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(loanController *controller.LoanController) *gin.Engine {
	r := gin.Default()

	loans := r.Group("/loans")
	{
		loans.POST("/", loanController.ApplyForLoan())
		loans.GET("/:id", loanController.ViewLoanStatus())
		loans.GET("", loanController.ViewAllLoans())
		loans.PATCH("/:id/status", loanController.ApproveRejectLoan())
		loans.DELETE("/:id", loanController.DeleteLoan())
		loans.GET("/logs", loanController.ViewSystemLogs())
	}

	return r
}
