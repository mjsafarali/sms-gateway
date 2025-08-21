package validator

import (
	validator "github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"
)

type echoValidator struct {
	validator *validator.Validate
}

func New() echo.Validator {
	v := echoValidator{validator: validator.New()}

	return &v
}

func (v *echoValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func (v *echoValidator) MustRegister(tag string, fn validator.Func) {
	if err := v.validator.RegisterValidation(tag, fn); err != nil {
		panic(err)
	}
}
