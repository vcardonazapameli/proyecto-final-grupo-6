package handlers

import (
	"net/http"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/product_record"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/response"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/bootcamp-go/web/request"
)

func NewProductRecordHandler(service product_record.ProductRecordService) *ProductRecordHandler {
	return &ProductRecordHandler{service: service}
}

type ProductRecordHandler struct {
	service product_record.ProductRecordService
}

func (productRecordHandler *ProductRecordHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productRecordDoc := models.ProductRecordDocRequest{}
		if err := request.JSON(r, &productRecordDoc); err != nil {
			response.Error(w, err)
			return
		}
		if err := validators.ValidateNoEmptyFields(productRecordDoc); err != nil {
			response.Error(w, err)
			return
		}
		product, err := productRecordHandler.service.Create(productRecordDoc)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusCreated, product)
	}
}
