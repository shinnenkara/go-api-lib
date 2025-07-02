package api

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

type FieldValidationError struct {
	Field       string `json:"field"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func CommonFieldValidationError(field string) FieldValidationError {
	return FieldValidationError{
		Field:       field,
		Message:     "Field validation failed",
		Description: "This field contains validation error",
	}
}

func NewValidationError(err error) Error[any] {
	validationError := GetError[FieldValidationError](400)

	var merror *json.UnmarshalTypeError
	if errors.As(err, &merror) {
		out := make([]FieldValidationError, 1)
		out[0] = CommonFieldValidationError(merror.Field)
		out[0].Message = merror.Error()

		validationError.Errors = out
	}

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]FieldValidationError, len(ve))
		for i, fe := range ve {
			field := removeStructField(fe.Namespace())
			errorMessage := getNestedErrorMessage(fe.Error())

			out[i] = CommonFieldValidationError(field)
			out[i].Message = errorMessage
		}
		validationError.Errors = out
	}

	return Error[any]{
		Code:        validationError.Code,
		Message:     validationError.Message,
		Description: validationError.Description,
		Errors:      convertSliceToAny(validationError.Errors),
	}
}

func convertSliceToAny[T any](input []T) []any {
	result := make([]any, len(input))
	for i, v := range input {
		result[i] = v
	}
	return result
}

func removeStructField(field string) string {
	fields := strings.Split(field, ".")
	fields = fields[1:]
	structField := strings.Join(fields, ".")
	if len(fields) < 1 {
		structField = field
	}

	return structField
}

func getNestedErrorMessage(nestedErrors string) string {
	errors := strings.Split(nestedErrors, "Error:")
	errors = errors[1:]
	nestedError := strings.Join(errors, ", ")
	if len(errors) < 1 {
		nestedError = nestedErrors
	}

	return nestedError
}
