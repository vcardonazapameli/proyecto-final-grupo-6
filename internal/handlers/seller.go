package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/seller"
	defaultErrors "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/response"
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
			response.Error(w, err)
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
			response.Error(w, defaultErrors.ErrorBadRequest)
			return
		}

		new, err := h.sv.Create(req.Cid, req.CompanyName, req.Address, req.Telephone)

		// Error handling
		if err != nil {
			response.Error(w, err)
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
			response.Error(w, err) // BadRequest
			return
		}

		s, err := h.sv.GetByID(id)

		// Error handling
		if err != nil {
			response.Error(w, err)
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
			response.Error(w, defaultErrors.ErrorBadRequest) // BadRequest
			return
		}

		// Error handling
		if err := h.sv.Delete(id); err != nil {
			response.Error(w, err)
			return
		}

		// Deleted Successfully
		response.JSON(w, http.StatusNoContent, nil)

	}
}

type UpdateSellerRequest struct {
	Cid         *int    `json:"cid"`
	CompanyName *string `json:"company_name"`
	Address     *string `json:"address"`
	Telephone   *int    `json:"telephone"`
}

func (h *SellerHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.Error(w, defaultErrors.ErrorBadRequest)
			return
		}

		req := &UpdateSellerRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.Error(w, defaultErrors.ErrorBadRequest)
			return
		}

		sellerDoc, err := h.sv.Update(id, req.Cid, req.CompanyName, req.Address, req.Telephone)

		// Error handling
		// TODO: ValidationError
		if err != nil {
			response.Error(w, err)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "updated successfully",
			"data":    sellerDoc,
		})
	}
}
