package common 

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

const Separator = "\n"

// FiledErrorsAsString composes custom error mesage from:
// "Key: 'Name' Error:Field validation for 'Name' failed on the 'required' tag"
// to
// "validation failed on field Name, condition: required"
func FiledErrorsAsString(errors []validator.FieldError) string {
	var sb strings.Builder

	for i, ve := range errors {
		sb.WriteString("validation failed on field '" + ve.Field() + "'")
		sb.WriteString(", condition: " + ve.ActualTag())

		// Print condition parameters, e.g. oneof=red blue -> { red blue }
		if ve.Param() != "" {
			sb.WriteString(" { " + ve.Param() + " }")
		}

		if ve.Value() != nil && ve.Value() != "" {
			sb.WriteString(fmt.Sprintf(", actual: %v", ve.Value()))
		}
		if i < len(errors)-1 {
			sb.WriteString(Separator)
		}
	}

	return sb.String()
}
