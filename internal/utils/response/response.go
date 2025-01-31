package response

import (
	"encoding/json"
	"errors"
	"net/http"

	e "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
)

type customResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func JSON(w http.ResponseWriter, code int, data any) {

	response := customResponse{
		Data:    data,
		Message: "Success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// encode response
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		// default error
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type customErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func Error(w http.ResponseWriter, err error) {
	// response
	var statusCode int
	var message string

	switch {
	case errors.Is(err, e.ErrorNotFound):
		statusCode = http.StatusNotFound
		message = err.Error()

	case errors.Is(err, e.ErrorConflict):
		statusCode = http.StatusConflict
		message = err.Error()

	case errors.Is(err, e.ErrorBadRequest):
		statusCode = http.StatusBadRequest
		message = err.Error()

	case errors.Is(err, e.ErrorUnprocessableContent):
		statusCode = http.StatusUnprocessableEntity
		message = err.Error()

	case errors.As(err, &e.ValidationError{}):
		statusCode = http.StatusBadRequest
		message = err.Error()

	default:
		statusCode = http.StatusInternalServerError
		message = "Internal Server Error"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := customErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
	// encode response
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		// default error
		w.WriteHeader(http.StatusInternalServerError)
	}
}
