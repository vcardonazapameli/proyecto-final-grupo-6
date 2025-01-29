package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	repositories "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/seller"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/seller"
	defaultErrors "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
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

func (h *SellerHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "ID format not correct. Must be int")
			return
		}

		s, err := h.sv.GetByID(id)

		// Error handling
		if err != nil {
			if errors.Is(err, defaultErrors.ErrorNotFound) {
				response.Error(w, http.StatusNotFound, "Seller not found")
				return
			}
			response.Error(w, http.StatusInternalServerError, "Unknown error ocurred")
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    s,
		})
	}
}

func (h *SellerHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Incorrect ID")
			return
		}

		// Error handling
		if err := h.sv.Delete(id); err != nil {
			if errors.Is(err, repositories.ErrorNotFound) {
				response.Error(w, http.StatusNotFound, "Seller not found")
				return
			}
			response.Error(w, http.StatusInternalServerError, "Internal error")
			return
		}

		// Deleted Successfully
		response.JSON(w, http.StatusNoContent, nil)

	}
}
