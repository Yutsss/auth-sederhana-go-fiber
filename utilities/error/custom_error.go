package error

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type CustomError interface {
	Error() string
	Code() int
	UnWrap() error
}

type customError struct {
	err  error
	code int
}

func NewCustomError(err error, code int) CustomError {
	return &customError{
		err:  err,
		code: code,
	}
}

func (e *customError) Error() string {
	return e.err.Error()
}

func (e *customError) Code() int {
	return e.code
}

func (e *customError) UnWrap() error {
	return e.err
}

const (
	MESSAGE_FAILED_TO_GET_DATA_FROM_BODY = "Failed to get data from body"
	MESSAGE_FAILED_TO_AUTHORIZE_USER     = "Failed to authorize user"
	MESSAGE_FAILED_REGISTER_USER         = "Failed to register user"
	MESSAGE_FAILED_LOGIN_USER            = "Failed to login user"
	MESSAGE_FAILED_GET_USER              = "Failed to get user"
	MESSAGE_FAILED_LOGOUT_USER           = "Failed to logout user"
)

var (
	ErrInternalServer       = NewCustomError(errors.New("Internal server error"), fiber.StatusInternalServerError)
	ErrBadRequest           = NewCustomError(errors.New("Bad request"), fiber.StatusBadRequest)
	ErrNotAllowed           = NewCustomError(errors.New("Not allowed"), fiber.StatusMethodNotAllowed)
	ErrEmailAlreadyExist    = NewCustomError(errors.New("Email already exist"), fiber.StatusConflict)
	ErrWrongEmailOrPassword = NewCustomError(errors.New("Wrong email or password"), fiber.StatusBadRequest)
	ErrUnauthorized         = NewCustomError(errors.New("Unauthorized"), fiber.StatusUnauthorized)
	ErrUserNotFound         = NewCustomError(errors.New("User not found"), fiber.StatusNotFound)
	ErrTimeout              = NewCustomError(errors.New("Timeout"), fiber.StatusRequestTimeout)
)
