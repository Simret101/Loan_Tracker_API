package router

import (
	"log"
	"os"

	"Loan_Tracker_API/database"
	"Loan_Tracker_API/delivery/controller"
	"Loan_Tracker_API/infrastructure/middleware"
	tokenservice "Loan_Tracker_API/infrastructure/token_service"
	"Loan_Tracker_API/repository"
	"Loan_Tracker_API/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func NewUserRoute(group *gin.RouterGroup, userCollection database.CollectionInterface) {
	repo := repository.NewUserRepository(userCollection)
	usecase := usecase.NewUserUseCase(repo)
	ctrl := controller.NewUserController(usecase)

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Panic(err.Error())
	}
	accessSecret := os.Getenv("ACCESSTOKENSECRET")
	if accessSecret == "" {
		log.Panic("No access token")
	}

	refreshSecret := os.Getenv("REFRESHTOKENSECRET")
	if refreshSecret == "" {
		log.Panic("No refresh token")
	}
	tokenSvc := tokenservice.NewTokenService(accessSecret, refreshSecret)

	loggedInMiddleware := middleware.LoggedIn(*tokenSvc)
	mustOwn := middleware.RoleBasedAuth(false)
	mustBeAdmin := middleware.RoleBasedAuth(true)

	group.GET("/user/:id", loggedInMiddleware, ctrl.GetOneUser())         // Added leading slash
	group.GET("/users", loggedInMiddleware, mustBeAdmin, ctrl.GetUsers()) // Added leading slash

	group.PUT("/user/:id", loggedInMiddleware, mustOwn, ctrl.UpdateUser())    // Added leading slash
	group.DELETE("/user/:id", loggedInMiddleware, mustOwn, ctrl.DeleteUser()) // Added leading slash
}
