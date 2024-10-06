package member

import (
	"github.com/go-playground/validator/v10"
)

func NameValid(fl validator.FieldLevel) bool {
	return fl.Field().String() != "admin"
}
