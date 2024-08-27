package controller

import (
	"net/http"

	"Loan_Tracker_API/domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.User_Usecase_interface
}

func NewUserController(usecase domain.User_Usecase_interface) *UserController {
	return &UserController{UserUsecase: usecase}
}

func (controller *UserController) GetOneUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		user, err := controller.UserUsecase.GetOneUser(id)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		ctx.IndentedJSON(http.StatusOK, gin.H{"data": user})
	}
}

func (controller *UserController) GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := controller.UserUsecase.GetUsers()
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Users not found"})
			return
		}
		ctx.IndentedJSON(http.StatusOK, gin.H{"data": users})
	}
}

func (controller *UserController) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var input domain.UpdateUser
		if err := ctx.BindJSON(&input); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		user, err := controller.UserUsecase.UpdateUser(id, input)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusOK, gin.H{"data": user})
	}
}

func (controller *UserController) DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		err := controller.UserUsecase.DeleteUser(id)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}
