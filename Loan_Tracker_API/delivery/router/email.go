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

func NewVerifyEmialRoute(group *gin.RouterGroup, user_collection database.CollectionInterface) {
	repo := repository.NewUserRepository(user_collection)
	user_usecase := usecase.NewUserUseCase(repo)

	email_repo := repository.NewEmailVRepo(*repo)
	email_usecase := usecase.NewEmailVUsecase(user_usecase, email_repo)
	email_ctrl := controller.NewEmailVController(email_usecase, user_usecase)

	//load middlewares
	err := godotenv.Load()
	if err != nil {
		log.Panic(err.Error())
	}
	access_secret := os.Getenv("ACCESSTOKENSECRET")
	if access_secret == "" {
		log.Panic("No accesstoken")
	}

	refresh_secret := os.Getenv("REFRESHTOKENSECRET")
	if refresh_secret == "" {
		log.Panic("No refreshtoken")
	}
	TokenSvc := *tokenservice.NewTokenService(access_secret, refresh_secret)

	LoggedInmiddleWare := middleware.LoggedIn(TokenSvc)

	group.POST("api/verify-email/:id", LoggedInmiddleWare, email_ctrl.SendVerificationEmail())
	group.GET("api/verify-email/:token", LoggedInmiddleWare, email_ctrl.VerifyEmail())

	group.POST("api/forget-password/:id", email_ctrl.SendForgetPasswordEmail())
	group.GET("/api/forget-password/", email_ctrl.ForgetPasswordValidate())
}