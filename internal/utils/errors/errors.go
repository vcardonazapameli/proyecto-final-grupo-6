package errors

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// Generics
	ErrorNotFound            error = errors.New("resource not found")
	ErrorInternalServerError error = errors.New("internal server error")
	ErrorConflict            error = errors.New("conflict occurred")
	ErrorBadRequest          error = errors.New("bad request")
	ErrorUnprocessableContent error = errors.New("unprocessable content")
  
	//Warehouse
	ErrorDataIncorrect       error = errors.New("incorrectly formatted or incomplete warehouse data")
	ErrorWarehouseCoreRepeat error = errors.New("warehouse code already exists")

)

type ValidationError struct {
	Messages []string
}

func (ve ValidationError) Error() string {
	return fmt.Sprintf("There were some errors validating:  %s", strings.Join(ve.Messages, ", "))
}
