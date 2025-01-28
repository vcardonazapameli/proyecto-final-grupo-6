package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	repositories "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/seller"
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

type SellerRequest struct {
	Cid         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   int    `json:"telephone"`
}

func (h *SellerHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &SellerRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.Error(w, http.StatusBadRequest, "There was a problem decoding JSON")
			return
		}

		new, err := h.sv.Create(req.Cid, req.CompanyName, req.Address, req.Telephone)

		// Error handling
		if err != nil {
			if errors.Is(err, repositories.ExistingCIdError) {
				response.Error(w, http.StatusConflict, "Could not create Seller: Existing CID")
				return
			}
			if errors.As(err, &service.ValidationError{}) {
				response.Error(w, http.StatusBadRequest, err.Error())
				return
			}
			response.Error(w, 0, err.Error())
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    new,
		})
	}
}
