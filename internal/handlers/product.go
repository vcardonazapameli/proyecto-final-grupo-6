package handlers

import (
	"net/http"
	"strconv"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/product"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/response"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/bootcamp-go/web/request"
	"github.com/go-chi/chi/v5"
)

func NewProductHandler(sv product.ProductService) *ProductHandler {
	return &ProductHandler{sv: sv}
}

type ProductHandler struct {
	sv product.ProductService
}

func (h *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productsDoc, err := h.sv.GetAll()
		if err != nil {
			response.Error(w, err)
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
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusOK, productDoc)
	}
}

func (h *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		err := h.sv.Delete(id)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusNoContent, "el registro se elimino correctamente")
	}
}

func (h *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productDoc := models.ProductDocRequest{}
		if err := request.JSON(r, &productDoc); err != nil {
			response.Error(w, err)
			return
		}
		if err := validators.ValidateNoEmptyFields(productDoc); err != nil {
			response.Error(w, err)
			return
		}
		product, err := h.sv.Create(productDoc)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusCreated, product)
	}
}

func (h *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		productDoc := models.ProductUpdateDocRequest{}
		if err := request.JSON(r, &productDoc); err != nil {
			response.Error(w, err)
			return
		}
		product, err := h.sv.Update(id, productDoc)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusOK, product)
	}
}
