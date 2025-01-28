package handlers

import (
	"net/http"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/seller"
	"github.com/bootcamp-go/web/response"
)

type SellerHandler struct {
	sv service.SellerService
}

func NewSellerHandler(sv service.SellerService) *SellerHandler {
	return &SellerHandler{sv}
}

func (h *SellerHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := h.sv.GetAll()

		if err != nil {
			response.Error(w, http.StatusInternalServerError, "There was an error when trying to fetch all sellers")
			return
		}

		response.JSON(w, http.StatusOK, s)
	}
}
