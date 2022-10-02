package entity

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("userType", validUserType)
}

func validUserType(fl validator.FieldLevel) bool {
	return UserType(fl.Field().String()).IsValid()
}
