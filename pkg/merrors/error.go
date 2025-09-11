package merrors

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound = errors.New("resource not found")
)

func ToHTTPStatus(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
