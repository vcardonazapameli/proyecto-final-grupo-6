package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/warehouse"
	errorsCustom "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
	mapper "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/bootcamp-go/web/response"
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
			response.JSON(w, http.StatusInternalServerError, errorsCustom.ErrorInternalServerError)
			return
		}

		data := []models.WarehouseDoc{}
		for _, value := range warehouses {
			data = append(data, mapper.WarehouseToWarehouseDoc(value))
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": http.StatusOK,
			"data":    data,
		})
	}
}

func (h *WarehouseHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			response.Error(w, http.StatusBadRequest, errorsCustom.ErrorBadRequest.Error())
			return
		}
		warehouse, err := h.sv.GetById(id)
		if err != nil {
			if err == errorsCustom.ErrorNotFound {
				response.Error(w, http.StatusNotFound, errorsCustom.ErrorNotFound.Error())
			} else {
				response.Error(w, http.StatusInternalServerError, errorsCustom.ErrorInternalServerError.Error())
			}
			return
		}
		warehouseResponse := mapper.WarehouseToWarehouseDoc(warehouse)

		response.JSON(w, http.StatusOK, map[string]any{
			"data": warehouseResponse,
		})
	}
}

func (h *WarehouseHandler) CreateWarehouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var warehouseData models.Warehouse
		err := json.NewDecoder(r.Body).Decode(&warehouseData)
		if err != nil {
			response.Error(w, http.StatusBadRequest, errorsCustom.ErrorDataIncorrect.Error())
			return
		}
		newWarehouse, err := h.sv.CreateWarehouse(warehouseData)
		if err != nil {
			if err.Error() == errorsCustom.ErrorWarehouseCoreRepeat.Error() {
				response.Error(w, http.StatusConflict, errorsCustom.ErrorDataIncorrect.Error())
			} else {
				response.Error(w, http.StatusInternalServerError, errorsCustom.ErrorInternalServerError.Error())
			}
			return
		}
		warehouseResponse := mapper.WarehouseToWarehouseDoc(newWarehouse)
		response.JSON(w, http.StatusCreated, map[string]any{
			"data": warehouseResponse,
		})
	}
}

func (h *WarehouseHandler) DeleteWarehouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			response.Error(w, http.StatusBadRequest, errorsCustom.ErrorBadRequest.Error())
			return
		}
		err = h.sv.DeleteWarehouse(id)
		if err != nil {
			if err == errorsCustom.ErrorNotFound {
				response.Error(w, http.StatusNotFound, errorsCustom.ErrorNotFound.Error())
			} else {
				response.Error(w, http.StatusInternalServerError, errorsCustom.ErrorInternalServerError.Error())
			}
			return
		}
		response.JSON(w, http.StatusNoContent, "Warehouse deleted")
	}
}

func (h *WarehouseHandler) UpdateWarehouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			response.Error(w, http.StatusBadRequest, errorsCustom.ErrorBadRequest.Error())
			return
		}

		var warehouseData models.Warehouse
		err = json.NewDecoder(r.Body).Decode(&warehouseData)
		if err != nil {
			response.Error(w, http.StatusBadRequest, errorsCustom.ErrorDataIncorrect.Error())
			return
		}

		warehouseRepository, err := h.sv.UpdateWarehouse(id, warehouseData)
		if warehouseRepository == (models.Warehouse{}) && err == nil {
			response.Error(w, http.StatusConflict, errorsCustom.ErrorConflict.Error())
			return
		}
		if err != nil {
			response.Error(w, http.StatusInternalServerError, errorsCustom.ErrorInternalServerError.Error())
			return
		}
		warehouseRepository.Warehouse_code = warehouseData.Warehouse_code
		warehouseRepository.Address = warehouseData.Address
		warehouseRepository.Telephone = warehouseData.Telephone
		warehouseRepository.Minimun_capacity = warehouseData.Minimun_capacity
		warehouseRepository.Minimun_temperature = warehouseData.Minimun_temperature
		warehouseRepository.Locality_id = warehouseData.Locality_id

		warehouseResponse := mapper.WarehouseToWarehouseDoc(warehouseRepository)

		response.JSON(w, http.StatusOK, warehouseResponse)
	}
}
