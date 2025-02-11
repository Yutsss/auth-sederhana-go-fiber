package utilities

import "github.com/gofiber/fiber/v2"

type APIResponse struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func ResponseSuccess(ctx *fiber.Ctx, message string, data any, code int) error {
	res := APIResponse{
		Status:  true,
		Code:    code,
		Message: message,
		Data:    data,
	}

	return ctx.Status(code).JSON(res)
}

func ResponseError(ctx *fiber.Ctx, message string, err string, code int) error {
	res := APIResponse{
		Status:  false,
		Code:    code,
		Message: message,
		Error:   err,
	}

	return ctx.Status(code).JSON(res)
}
