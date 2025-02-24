package handlers

import (
	"net/http"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/carrier"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/response"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/bootcamp-go/web/request"
)

func NewCarrierHandler(service carrier.CarrierService) *CarrierHandler {
	return &CarrierHandler{service: service}
}

type CarrierHandler struct {
	service carrier.CarrierService
}

func (carrierHandler *CarrierHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		carrierDocRequest := models.CarrierDocRequest{}
		if err := request.JSON(r, &carrierDocRequest); err != nil {
			response.Error(w, err)
			return
		}
		if err := validators.ValidateNoEmptyFields(carrierDocRequest); err != nil {
			response.Error(w, err)
			return
		}
		carrier, err := carrierHandler.service.Create(carrierDocRequest)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusCreated, carrier)
	}
}
