package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/employee"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/bootcamp-go/web/response"
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
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}

func (h *EmployeeHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid ID")
			return
		}

		data, err := h.sv.GetById(id)

		if err != nil {
			if errors.Is(err, customErrors.ErrorNotFound) {
				response.Error(w, http.StatusNotFound, err.Error())
				return
			}
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}

func (h *EmployeeHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newEmployee models.RequestEmployee

		if err := json.NewDecoder(r.Body).Decode(&newEmployee); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid JSON for employee")
			return
		}

		data, err := h.sv.Create(newEmployee)

		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"data": data,
		})
	}
}

func (h *EmployeeHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid ID")
			return
		}

		err = h.sv.Delete(id)

		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}
