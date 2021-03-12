package common 

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type SerializedError struct {
	Errors []string `json:"errors"`
}

type IApiError interface {
	Error() string
	SerializeErrors() []string
	GetStatusCode() int
}

type ApiError struct {
	IApiError
	StatusCode int
	Reason     string
}

func (e ApiError) Error() string {
	return fmt.Sprintf(e.Reason)
}

func (e *ApiError) SerializeErrors() []string {
	return []string{e.Reason}
}

func (e *ApiError) GetStatusCode() int {
	return e.StatusCode

}

type DatabaseConnectionError struct {
	*ApiError
}

func NewDataBaseConnectionError(err error) ApiError {
	dbConnError := &DatabaseConnectionError{
		&ApiError{
			StatusCode: http.StatusInternalServerError,
			Reason:     err.Error(),
		},
	}
	return *dbConnError.ApiError
}

type RequestValidationError struct {
	*ApiError
}

func NewRequestValidationError(errors []validator.FieldError) ApiError {
	reqValidationErr := &RequestValidationError{
		&ApiError{
			StatusCode: http.StatusBadRequest,
			Reason:     FiledErrorsAsString(errors),
		},
	}
	return *reqValidationErr.ApiError
}

func (e *RequestValidationError) SerializeErrors() []string {
	return strings.Split(e.Reason, Separator)
}

type NotFoundError struct {
	*ApiError
}

func NewNotFoundError() ApiError {
	notFoundErr := &NotFoundError{
		&ApiError{
			StatusCode: http.StatusNotFound,
			Reason:     "Not found",
		},
	}
	return *notFoundErr.ApiError
}

type BadRequestError struct {
	*ApiError
}

func NewBadRequestError(message string) ApiError {
	badRequestError := &BadRequestError{
		&ApiError{
			StatusCode: http.StatusBadRequest,
			Reason:     message,
		},
	}
	return *badRequestError.ApiError
}

type NotAuthorizedError struct {
	*ApiError
}

func NewNotAuthorizedError() ApiError {
	notAuthorizedErr := &NotAuthorizedError{
		&ApiError{
			StatusCode: http.StatusUnauthorized,
			Reason:     "Not Authorized",
		},
	}
	return *notAuthorizedErr.ApiError
}
