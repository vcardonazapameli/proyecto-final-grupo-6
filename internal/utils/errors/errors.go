package errors

import "errors"

var (
	// Generics
	ErrorNotFound             error = errors.New("not found")
	ErrorInternalServerError  error = errors.New("internal server error")
	ErrorConflict             error = errors.New("conflict")
	ErrorBadRequest           error = errors.New("bad request")
	ErrorUnprocessableContent error = errors.New("unprocessable content")
)
