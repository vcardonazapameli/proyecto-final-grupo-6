package handlers

import (
	"net/http"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/employee"
	"github.com/bootcamp-go/web/response"
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
		data, _ := h.sv.GetAll()

		println(data)

		response.JSON(w, http.StatusOK, "hola")
	}
}
