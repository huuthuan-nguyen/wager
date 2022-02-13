package utils

import "net/http"

type Error struct {
	StatusCode int
	Message    string
}

// Custom error
func (e *Error) Error() string {
	return e.Message
}

// NewError /**
func NewError(c int, s string) Error {
	return Error{c, s}
}

// PanicInternalServerError /**
func PanicInternalServerError() {
	panic(NewError(http.StatusInternalServerError, "Something went wrong."))
}

// PanicBadRequest /**
func PanicBadRequest() {
	panic(NewError(http.StatusBadRequest, "Bad request."))
}

// PanicTooManyRequests /**
func PanicTooManyRequests() {
	panic(NewError(http.StatusTooManyRequests, "Too many requests."))
}

// PanicNotFound /**
func PanicNotFound() {
	panic(NewError(http.StatusNotFound, "Not found."))
}
