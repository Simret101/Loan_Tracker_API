package router

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"Loan_Tracker_API/delivery/controller"
	"Loan_Tracker_API/infrastructure/middleware"
	
)

func SetupRouter(
	r *gin.Engine,
	loanController *controller.LoanController,
	authController *controller.AuthController,
 // Added AdminController
) {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Panic(err.Error())
	}
	accessSecret := os.Getenv("ACCESSTOKENSECRET")
	if accessSecret == "" {
		log.Panic("No access token secret")
	}
	refreshSecret := os.Getenv("REFRESHTOKENSECRET")
	if refreshSecret == "" {
		log.Panic("No refresh token secret")
	}

	// Initialize token service
	tokenSvc := token_service.NewTokenService(accessSecret, refreshSecret)

	// Middlewares
	authMiddleware := middleware.LoggedIn(tokenSvc)
	roleMiddleware := middleware.RoleBasedAuth(true) // Adjust based on roles needed

	// Register routes
	RegisterLoanRoutes(r, loanController, authMiddleware)
	RegisterAuthRoutes(r, authController)
	RegisterAdminRoutes(r, adminController, roleMiddleware)
}

func RegisterLoanRoutes(
	r *gin.Engine,
	loanController *controller.LoanController,
	authMiddleware gin.HandlerFunc,
) {
	loanRoute := r.Group("/loans").Use(authMiddleware)
	{
		loanRoute.POST("/", loanController.ApplyForLoan())
		loanRoute.GET("/:id", loanController.ViewLoanStatus())
	}
}

func RegisterAdminRoutes(
	r *gin.Engine,
	adminController *controller.AdminController,
	roleMiddleware gin.HandlerFunc,
) {
	adminRoute := r.Group("/admin").Use(roleMiddleware)
	{
		adminRoute.GET("/loans", adminController.ViewAllLoans())
		adminRoute.PATCH("/loans/:id/status", adminController.ApproveRejectLoan())
		adminRoute.DELETE("/loans/:id", adminController.DeleteLoan())
		adminRoute.GET("/logs", adminController.ViewSystemLogs())
	}
}

func RegisterAuthRoutes(r *gin.Engine, authController *controller.AuthController) {
	authRoute := r.Group("/auth")
	{
		authRoute.POST("/signup", authController.SignUp())
		authRoute.POST("/login", authController.LogIn())
		authRoute.POST("/logout", authController.LogOut())
		authRoute.POST("/refresh", authController.Refresh())
		authRoute.GET("/google", authController.GoogleLogIn())
		authRoute.GET("/oauth2/callback/google", authController.GoogleCallBack())
	}
}
