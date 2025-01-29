package errors

import "errors"

var (
	// Generics
	ErrorNotFound            error = errors.New("resource not found")
	ErrorInternalServerError error = errors.New("internal server error")
	ErrorConflict            error = errors.New("conflict occurred")
	ErrorBadRequest          error = errors.New("bad request")

	//Warehouse
	ErrorDataIncorrect       error = errors.New("incorrectly formatted or incomplete warehouse data")
	ErrorWarehouseCoreRepeat error = errors.New("warehouse code already exists")
)
