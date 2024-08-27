package router

import (
	"Loan_Tracker_API/delivery/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	
	userController := &controller.UserController{}
	authController := &controller.AuthController{}

	authMiddleware := func(c *gin.Context) {}
	roleMiddleware := func(c *gin.Context) {}

	
	RegisterUserRoutes(r, userController, authMiddleware, roleMiddleware)
	RegisterAuthRoutes(r, authController)
}

func RegisterUserRoutes(
	r *gin.Engine,
	userController *controller.UserController,
	authMiddleware gin.HandlerFunc,
	roleMiddleware gin.HandlerFunc,
) {
	userRoute := r.Group("/users").Use(authMiddleware)
	{
		userRoute.GET("/:id", userController.GetOneUser())
		userRoute.GET("/", roleMiddleware, userController.GetUsers())
		userRoute.PUT("/:id", roleMiddleware, userController.UpdateUser())
		userRoute.DELETE("/:id", roleMiddleware, userController.DeleteUser())
		userRoute.GET("/search", roleMiddleware, userController.GetUsers())
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
