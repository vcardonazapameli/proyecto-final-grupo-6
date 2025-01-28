package handlers

import (
	"net/http"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/seller"
)

type SellerHandler struct {
	sv service.SellerService
}

func NewSellerHandler(sv service.SellerService) *SellerHandler {
	return &SellerHandler{sv}
}

func (h *SellerHandler) GetAll() (v http.HandlerFunc) {
	return
}
