package config

import (
	"auth-sederhana-go-fiber/controllers"
	middleware "auth-sederhana-go-fiber/middlewares"
	"auth-sederhana-go-fiber/repositories"
	"auth-sederhana-go-fiber/services"
	"auth-sederhana-go-fiber/utilities"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserDependencyInjection(db *gorm.DB) controllers.UserController {
	userRepository := repositories.NewUserRepository(db)
	jwtUtils := utilities.NewJWTUtils()
	userService := services.NewUserService(userRepository, jwtUtils)
	userController := controllers.NewUserController(userService)

	return userController
}

func MiddlewareDependencyInjection() map[string]fiber.Handler {
	jwtUtils := utilities.NewJWTUtils()

	middlewaresMap := make(map[string]fiber.Handler)

	middlewaresMap["authMiddleware"] = middleware.AuthMiddleware(jwtUtils)

	return middlewaresMap
}
