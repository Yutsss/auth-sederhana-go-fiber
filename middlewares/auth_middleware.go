package middlewares

import (
	"auth-sederhana-go-fiber/utilities"
	errorUtils "auth-sederhana-go-fiber/utilities/error"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(jwtUtils utilities.JWTUtils) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Cookies("access_token")
		if tokenString == "" {
			return utilities.ResponseError(ctx, errorUtils.MESSAGE_FAILED_TO_AUTHORIZE_USER, errorUtils.ErrUnauthorized.Error(), errorUtils.ErrUnauthorized.Code())
		}

		payload, err := jwtUtils.GetPayload(tokenString)
		if err != nil {
			return utilities.ResponseError(ctx, errorUtils.MESSAGE_FAILED_TO_AUTHORIZE_USER, errorUtils.ErrUnauthorized.Error(), errorUtils.ErrUnauthorized.Code())
		}

		ctx.Locals("user", payload)

		return ctx.Next()
	}
}
