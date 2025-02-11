package controllers

import (
	"auth-sederhana-go-fiber/dtos"
	"auth-sederhana-go-fiber/services"
	"auth-sederhana-go-fiber/utilities"
	errorUtils "auth-sederhana-go-fiber/utilities/error"
	successUtils "auth-sederhana-go-fiber/utilities/success"
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

type (
	UserController interface {
		Register(ctx *fiber.Ctx) error
		Login(ctx *fiber.Ctx) error
		Get(ctx *fiber.Ctx) error
		Logout(ctx *fiber.Ctx) error
	}

	userController struct {
		userService services.UserService
	}
)

func NewUserController(us services.UserService) UserController {
	return &userController{
		userService: us,
	}
}

func (c *userController) Register(ctx *fiber.Ctx) error {
	var req dtos.UserRegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utilities.ResponseError(ctx, errorUtils.MESSAGE_FAILED_TO_GET_DATA_FROM_BODY, errorUtils.ErrBadRequest.Error(), errorUtils.ErrBadRequest.Code())
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println(req)

	resData, err := c.userService.Register(timeoutCtx, req)

	if errors.Is(timeoutCtx.Err(), context.DeadlineExceeded) {
		return utilities.ResponseError(ctx, errorUtils.MESSAGE_FAILED_REGISTER_USER, errorUtils.ErrTimeout.Error(), errorUtils.ErrTimeout.Code())
	}

	if err != nil {
		return utilities.ResponseError(ctx, errorUtils.MESSAGE_FAILED_REGISTER_USER, err.Error(), err.Code())
	}

	return utilities.ResponseSuccess(ctx, successUtils.MESSAGE_SUCCESS_REGISTER_USER, resData, http.StatusCreated)
}

func (c *userController) Login(ctx *fiber.Ctx) error {
	var req dtos.UserLoginRequest

	if err := ctx.BodyParser(&req); err != nil {
		return utilities.ResponseError(ctx, errorUtils.MESSAGE_FAILED_TO_GET_DATA_FROM_BODY, errorUtils.ErrBadRequest.Error(), errorUtils.ErrBadRequest.Code())
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resData, err := c.userService.Login(timeoutCtx, req)

	if errors.Is(timeoutCtx.Err(), context.DeadlineExceeded) {
		return utilities.ResponseError(ctx, errorUtils.MESSAGE_FAILED_LOGIN_USER, errorUtils.ErrTimeout.Error(), errorUtils.ErrTimeout.Code())
	}

	if err != nil {
		return utilities.ResponseError(ctx, errorUtils.MESSAGE_FAILED_LOGIN_USER, err.Error(), err.Code())
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    resData.AccessToken,
		Path:     "/",
		MaxAge:   3600,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
	})

	return utilities.ResponseSuccess(ctx, successUtils.MESSAGE_SUCCESS_LOGIN_USER, nil, http.StatusOK)
}

func (c *userController) Get(ctx *fiber.Ctx) error {
	var req dtos.UserGetByIdRequest

	user := ctx.Locals("user")
	if user == nil {
		return utilities.ResponseError(ctx, errorUtils.MESSAGE_FAILED_GET_USER, errorUtils.ErrUnauthorized.Error(), errorUtils.ErrUnauthorized.Code())
	}

	userID := user.(dtos.AuthPayload).UserID
	req.Id = userID

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resData, err := c.userService.GetById(timeoutCtx, req)

	if errors.Is(timeoutCtx.Err(), context.DeadlineExceeded) {
		return utilities.ResponseError(ctx, errorUtils.MESSAGE_FAILED_GET_USER, errorUtils.ErrTimeout.Error(), errorUtils.ErrTimeout.Code())
	}

	if err != nil {
		return utilities.ResponseError(ctx, errorUtils.MESSAGE_FAILED_GET_USER, err.Error(), err.Code())
	}

	return utilities.ResponseSuccess(ctx, successUtils.MESSAGE_SUCCESS_GET_USER, resData, http.StatusOK)
}

func (c *userController) Logout(ctx *fiber.Ctx) error {
	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
	})

	return utilities.ResponseSuccess(ctx, successUtils.MESSAGE_SUCCESS_LOGOUT_USER, nil, http.StatusOK)
}
