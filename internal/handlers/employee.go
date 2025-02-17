package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/employee"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/response"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/go-chi/chi/v5"
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

		data, err := h.sv.GetAll()
		if err != nil {
			response.Error(w, err)
			return
		}

		response.JSON(w, http.StatusOK, data)
	}
}

func (h *EmployeeHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.Error(w, customErrors.ErrorBadRequest)
			return
		}

		data, err := h.sv.GetById(id)
		if err != nil {
			response.Error(w, err)
			return
		}

		response.JSON(w, http.StatusOK, data)
	}
}

func (h *EmployeeHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newEmployee models.RequestEmployee
		if err := json.NewDecoder(r.Body).Decode(&newEmployee); err != nil {
			response.Error(w, customErrors.ErrorBadRequest)
			return
		}

		data, err := h.sv.Create(newEmployee)
		if err != nil {
			response.Error(w, err)
			return
		}

		response.JSON(w, http.StatusCreated, data)
	}
}

func (h *EmployeeHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.Error(w, customErrors.ErrorBadRequest)
			return
		}

		var updEmployee models.UpdateEmployee
		if err := json.NewDecoder(r.Body).Decode(&updEmployee); err != nil {
			response.Error(w, customErrors.ErrorBadRequest)
			return
		}

		data, err := h.sv.Update(id, updEmployee)
		if err != nil {
			response.Error(w, err)
			return
		}

		response.JSON(w, http.StatusOK, data)
	}
}

func (h *EmployeeHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.Error(w, customErrors.ErrorBadRequest)
			return
		}

		err = h.sv.Delete(id)
		if err != nil {
			response.Error(w, err)
			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}

func (h *EmployeeHandler) GetReportInboundOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employeeIDStr := r.URL.Query().Get("id")
		var employeeID *int
		if employeeIDStr != "" {
			id, err := strconv.Atoi(employeeIDStr)
			if err != nil {
				response.Error(w, customErrors.ErrorBadRequest)
				return
			}
			employeeID = &id
		}

		report, err := h.sv.GetReportInboundOrders(employeeID)
		if err != nil {
			response.Error(w, err)
			return
		}

		response.JSON(w, http.StatusOK, report)
	}
}
