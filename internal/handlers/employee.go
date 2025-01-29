package handlers

import (
	"net/http"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/employee"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/response"
)

func NewEmployeeHandler(sv service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{sv: sv}
}

// EmployeeDefault is a struct with methods that represent handlers for Employees
type EmployeeHandler struct {
	// sv is the service that will be used by the handler
	sv service.EmployeeService
}

func (h *EmployeeHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all
		data, err := h.sv.GetAll()

		if err != nil {
			response.Error(w, err)
			return
		}

		response.JSON(w, http.StatusOK, data)
	}
}
