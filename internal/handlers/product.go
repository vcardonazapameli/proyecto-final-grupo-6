package handlers

import (
	"errors"
	"net/http"
	"strconv"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/product"
	errorCustom "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
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

func (h *ProductHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		productDoc, err := h.sv.GetById(id)
		if err != nil {
			if errors.Is(err, errorCustom.ErrorNotFound) {
				response.Error(w, http.StatusNotFound, err.Error())
				return
			}
			response.Error(w, http.StatusInternalServerError, errorCustom.ErrorInternalServerError.Error())
			return
		}
		response.JSON(w, http.StatusOK, productDoc)
	}
}
