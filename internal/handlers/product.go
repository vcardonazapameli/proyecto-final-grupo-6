package handlers

import (
	"net/http"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/product"
	errorCustom "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
	"github.com/bootcamp-go/web/response"
)

func NewProductHandler(sv service.ProductService) *ProductHandler {
	return &ProductHandler{sv: sv}
}

type ProductHandler struct {
	sv service.ProductService
}

func (h *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productsDoc, err := h.sv.GetAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, errorCustom.ErrorInternalServerError.Error())
			return
		}
		response.JSON(w, http.StatusOK, productsDoc)
	}
}
