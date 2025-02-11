package validation

import (
	errorUtils "auth-sederhana-go-fiber/utilities/error"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Validate(data interface{}) errorUtils.CustomError {
	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			return CustomErrorMessage(validationErrors[0])
		} else {
			return errorUtils.ErrInternalServer
		}
	}

	return nil
}

func CustomErrorMessage(err validator.FieldError) errorUtils.CustomError {
	switch err.Tag() {
	case "required":
		return errorUtils.NewCustomError(fmt.Errorf("%s is required", err.Field()), http.StatusBadRequest)
	case "min":
		return errorUtils.NewCustomError(fmt.Errorf("%s must be at least %s", err.Field(), err.Param()), http.StatusBadRequest)
	case "max":
		return errorUtils.NewCustomError(fmt.Errorf("%s must be at most %s", err.Field(), err.Param()), http.StatusBadRequest)
	case "email":
		return errorUtils.NewCustomError(fmt.Errorf("%s is not a valid email", err.Field()), http.StatusBadRequest)
	default:
		return errorUtils.ErrInternalServer
	}
}
