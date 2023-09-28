package request_contract

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type IError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var inputValidator = validator.New()

func Validate(body interface{}) ([]IError, error) {
	inputValidator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	var errorsList []IError
	err := inputValidator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Message = fmt.Sprintf("Param '%s' %s", err.Field(), tagMessage(err.Tag(), err.Param()))
			errorsList = append(errorsList, el)
		}
		return errorsList, errors.New("invalid params")
	}

	return errorsList, nil
}

func tagMessage(tag string, param string) string {
	switch tag {
	case "required":
		return "is required"
	case "email":
		return "must be a valid email"
	case "eq":
		return fmt.Sprintf("must be equal to '%s'", param)
	case "ne":
		return fmt.Sprintf("must not be equal to '%s'", param)
	case "gt":
		return fmt.Sprintf("must be greater than '%s'", param)
	case "gte":
		return fmt.Sprintf("must be greater than or equal to '%s'", param)
	case "lt":
		return fmt.Sprintf("must be less than '%s'", param)
	case "lte":
		return fmt.Sprintf("must be less than or equal to '%s'", param)
	case "min":
		return fmt.Sprintf("must have at least '%s' characters", param)
	case "max":
		return fmt.Sprintf("must have a maximum of '%s' characters", param)
	default:
		return ""
	}
}
