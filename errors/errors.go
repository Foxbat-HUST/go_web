package errors

import (
	"errors"
	"net/http"
)

type Error struct {
	StatusCode int
	error      error
}

func newError(statusCode int, e error) *Error {
	return &Error{
		StatusCode: statusCode,
		error:      e,
	}
}

func (a *Error) Error() string {
	return a.error.Error()
}

func BadRequest(e error) error {
	return newError(http.StatusBadRequest, e)
}

func BadRequestFromStr(reason string) error {
	return newError(http.StatusBadRequest, errors.New(reason))
}

func NotFound(e error) error {
	return newError(http.StatusNotFound, e)
}

func InternalServerErr(e error) error {
	return newError(http.StatusInternalServerError, e)
}

func InternalServerErrFromStr(reason string) error {
	return newError(http.StatusInternalServerError, errors.New(reason))
}

func Unauthorized(e error) error {
	return newError(http.StatusUnauthorized, e)
}

func Forbidden(e error) error {
	return newError(http.StatusForbidden, e)
}
