package handlers

import (
	"encoding/json"
	"net/http"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/product_batch"
	defaultErrors "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/response"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type ProductBatchHandler struct {
	sv service.ProductBatchService
}

func NewProductBatchHandler(sv service.ProductBatchService) *ProductBatchHandler {
	return &ProductBatchHandler{sv: sv}
}

func (h *ProductBatchHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Data models.ProductBatchRequest `json:"data"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.Error(w, defaultErrors.ErrorBadRequest)
			return
		}

		newProductBatch, err := h.sv.Create(req.Data)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusCreated, newProductBatch)
	}
}
