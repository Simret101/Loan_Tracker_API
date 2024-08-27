package router

import (
	"Loan_Tracker_API/config"
	"Loan_Tracker_API/database"
	"Loan_Tracker_API/repository"

	"github.com/gin-gonic/gin"
)

func SetUp(router *gin.Engine) {
	var client config.ServerConnection
	client.Connect_could()

	userCollection := &database.MongoCollection{
		Collection: client.Client.Database("LoanTracker").Collection("Users"),
	}
	loanCollection := &database.MongoCollection{
		Collection: client.Client.Database("LoanTracker").Collection("Loans"),
	}
	paymentCollection := &database.MongoCollection{
		Collection: client.Client.Database("LoanTracker").Collection("Payments"),
	}
	interestCollection := &database.MongoCollection{
		Collection: client.Client.Database("LoanTracker").Collection("Interests"),
	}

	loanRoute := router.Group("")
	userRoute := router.Group("")
	authRoute := router.Group("")
	paymentRoute := router.Group("")
	interestRoute := router.Group("")

	NewLoanRoutes(loanRoute, loanCollection, userCollection)
	NewUserRoute(userRoute, userCollection)
	NewAuthRoute(authRoute, userCollection)
	NewPaymentRoutes(paymentRoute, paymentCollection, loanCollection)
	NewInterestRoutes(interestRoute, interestCollection, loanCollection)
}