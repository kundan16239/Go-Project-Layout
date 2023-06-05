package routes

import (
	"go-folder-sample/app/controllers"
	"go-folder-sample/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

type UserRoute struct {
	UserController *controllers.UserController
}

func NewUserRoute(userController *controllers.UserController) *UserRoute {
	// userController := controllers.NewUserController()

	return &UserRoute{
		UserController: userController,
	}
}

func (route *UserRoute) RegisterRoutes(router *fiber.App) {
	router.Post("/api/auth/local/login", route.UserController.Login)
	router.Post("/api/auth/local/register", route.UserController.Register)
	router.Post("/api/auth/check/userexist", route.UserController.UserExist)
	router.Post("/api/auth/token/refresh", route.UserController.RefreshToken)
	router.Post("/api/auth/sendEmailConfirmation", route.UserController.SendEmail)
	router.Post("/api/auth/forgotPassword", route.UserController.ForgotPassword)
	router.Post("/api/auth/resetPassword", route.UserController.ResetPassword)
	router.Post("/api/auth/changePassword", middlewares.Authentication(), route.UserController.ChangePassword)
}
