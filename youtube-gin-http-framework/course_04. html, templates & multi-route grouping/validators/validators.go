package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

/*
Using to check a field has datatype is String which contains word `Cool` or not. */
var ValidateCoolTitle = func(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "Cool")
}