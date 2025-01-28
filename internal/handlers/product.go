package handlers

import (
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/product"
)

func NewProductHandler(sv service.ProductService) *ProductHandler {
	return &ProductHandler{sv: sv}
}

type ProductHandler struct {
	sv service.ProductService
}
