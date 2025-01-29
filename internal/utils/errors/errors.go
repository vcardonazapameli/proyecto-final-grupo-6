package errors

import "errors"

var (
	// Generics
	ErrorNotFound            error = errors.New("resource not found")
	ErrorInternalServerError error = errors.New("internal server error")
	ErrorConflict            error = errors.New("conflict occurred")
	ErrorBadRequest          error = errors.New("bad request")
)
