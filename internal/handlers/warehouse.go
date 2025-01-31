package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/warehouse"
	errorsCustom "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
	mapper "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/response"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/go-chi/chi/v5"
)

func NewWarehouseHandler(sv service.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{sv: sv}
}

type WarehouseHandler struct {
	sv service.WarehouseService
}

func (h *WarehouseHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		warehouses, err := h.sv.GetAll()
		if err != nil {
			response.Error(w, err)
			return
		}

		data := make([]models.WarehouseDoc, 0, len(warehouses))
		for _, value := range warehouses {
			data = append(data, mapper.WarehouseToWarehouseDoc(value))
		}

		response.JSON(w, http.StatusOK, data)
	}
}

func (h *WarehouseHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			response.Error(w, errorsCustom.ErrorBadRequest)
			return
		}
		warehouse, err := h.sv.GetById(id)
		if err != nil {
			response.Error(w, err)
			return
		}
		warehouseResponse := mapper.WarehouseToWarehouseDoc(warehouse)

		response.JSON(w, http.StatusOK, warehouseResponse)
	}
}

func (h *WarehouseHandler) CreateWarehouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var warehouseData models.Warehouse
		err := json.NewDecoder(r.Body).Decode(&warehouseData)
		if err != nil {
			response.Error(w, errorsCustom.ErrorBadRequest)
			return
		}
		newWarehouse, err := h.sv.CreateWarehouse(warehouseData)
		if err != nil {
			response.Error(w, err)
			return
		}
		warehouseResponse := mapper.WarehouseToWarehouseDoc(newWarehouse)

		response.JSON(w, http.StatusCreated, warehouseResponse)
	}
}

func (h *WarehouseHandler) DeleteWarehouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			response.Error(w, errorsCustom.ErrorBadRequest)
			return
		}

		err = h.sv.DeleteWarehouse(id)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusNoContent, "")
	}
}

func (h *WarehouseHandler) UpdateWarehouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			response.Error(w, errorsCustom.ErrorBadRequest)
			return
		}

		var warehouseData models.WarehouseDocUpdate
		err = json.NewDecoder(r.Body).Decode(&warehouseData)
		if err != nil {
			response.Error(w, errorsCustom.ErrorBadRequest)
			return
		}

		updatedWarehouse, err := h.sv.UpdateWarehouse(id, warehouseData)
		if err != nil {
			response.Error(w, err)
			return
		}
		warehouseResponse := mapper.WarehouseToWarehouseDoc(updatedWarehouse)

		response.JSON(w, http.StatusOK, warehouseResponse)
	}
}
