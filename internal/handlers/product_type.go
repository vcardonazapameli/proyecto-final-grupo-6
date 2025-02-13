package handlers

import (
	"net/http"
	"strconv"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/product_type"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/response"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/validators"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/bootcamp-go/web/request"
	"github.com/go-chi/chi/v5"
)

func NewProductTypeHandler(service product_type.ProductTypeService) *ProductTypeHandler {
	return &ProductTypeHandler{service: service}
}

type ProductTypeHandler struct {
	service product_type.ProductTypeService
}

func (productTypeHandler *ProductTypeHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productsDoc, err := productTypeHandler.service.GetAll()
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusOK, productsDoc)
	}
}

func (productTypeHandler *ProductTypeHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		productTypeDoc, err := productTypeHandler.service.GetById(id)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusOK, productTypeDoc)
	}
}

func (productTypeHandler *ProductTypeHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productTypeDoc := models.ProductTypeDocRequest{}
		if err := request.JSON(r, &productTypeDoc); err != nil {
			response.Error(w, err)
			return
		}
		if err := validators.ValidateNoEmptyFields(productTypeDoc); err != nil {
			response.Error(w, err)
			return
		}
		product, err := productTypeHandler.service.Create(productTypeDoc)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusCreated, product)
	}
}
