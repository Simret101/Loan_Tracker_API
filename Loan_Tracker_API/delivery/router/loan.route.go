package router

import (
	"Loan_Tracker_API/delivery/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(loanController *controller.LoanPackageController) *gin.Engine {
	router := gin.Default()

	// Loan routes
	loans := router.Group("/loans")
	{
		loans.POST("/", loanController.CreateLoan())
		loans.GET("/", loanController.GetAllLoans())
		loans.GET("/:id", loanController.GetOneLoan())
		loans.PUT("/:id", loanController.UpdateLoan())
		loans.DELETE("/:id", loanController.DeleteLoan())
		loans.GET("/filter", loanController.FilterLoans())
		loans.GET("/unique", loanController.GetUniqueLoan())
	}

	// You can remove the user routes if they’re just placeholders for now

	return router
}
