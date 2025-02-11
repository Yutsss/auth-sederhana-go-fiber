package routes

import (
	"auth-sederhana-go-fiber/config"
	"auth-sederhana-go-fiber/utilities"
	successUtils "auth-sederhana-go-fiber/utilities/success"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	middlewares := config.MiddlewareDependencyInjection()
	userController := config.UserDependencyInjection(db)

	api := app.Group("/api")

	api.Get("/", func(ctx *fiber.Ctx) error {
		data := "Selamat datang di auth sederhana go fiber"
		return utilities.ResponseSuccess(ctx, successUtils.MESSAGE_SUCCESS_OK, data, http.StatusOK)
	})

	UserRoutes(api, userController, middlewares)
}
