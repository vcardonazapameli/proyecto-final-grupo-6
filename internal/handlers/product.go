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

func NewProductHandler(service product.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

type ProductHandler struct {
	service product.ProductService
}

func (productHandler *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productsDoc, err := productHandler.service.GetAll()
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusOK, productsDoc)
	}
}

func (productHandler *ProductHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		productDoc, err := productHandler.service.GetById(id)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusOK, productDoc)
	}
}

func (productHandler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		err := productHandler.service.Delete(id)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusNoContent, "el registro se elimino correctamente")
	}
}

func (productHandler *ProductHandler) Create() http.HandlerFunc {
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
		product, err := productHandler.service.Create(productDoc)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusCreated, product)
	}
}

func (productHandler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		productDoc := models.ProductUpdateDocRequest{}
		if err := request.JSON(r, &productDoc); err != nil {
			response.Error(w, err)
			return
		}
		product, err := productHandler.service.Update(id, productDoc)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusOK, product)
	}
}
