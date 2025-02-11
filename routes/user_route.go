package routes

import (
	"auth-sederhana-go-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api fiber.Router, userController controllers.UserController, middlewares map[string]fiber.Handler) {
	userGroup := api.Group("/users")

	userGroup.Post("/register", userController.Register)
	userGroup.Post("/login", userController.Login)
	userGroup.Get("/me", middlewares["authMiddleware"], userController.Get)
	userGroup.Post("/logout", middlewares["authMiddleware"], userController.Logout)
}
