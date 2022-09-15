package common

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type (
	ValidationUtil struct {
		validator *validator.Validate
	}
)

func NewValidationUtil() echo.Validator {
	return &ValidationUtil{validator: validator.New()}
}

func (v *ValidationUtil) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
