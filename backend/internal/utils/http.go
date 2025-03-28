package utils

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()

	// register function to get tag name from json tags.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

type ValidationError struct {
	Field string `json:"field"`
	error string `json:"error"`
}

func (e ValidationError) Error() string {
	return e.error
}

func ValidateStruct(strct interface{}) []ValidationError {
	err := validate.Struct(strct)
	if err == nil {

	}
}
