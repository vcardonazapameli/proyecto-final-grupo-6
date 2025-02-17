package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/seller"
	defaultErrors "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	customResponse "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/response"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
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
			customResponse.Error(w, err)
			return
		}

		customResponse.JSON(w, http.StatusOK, s)
	}
}

type SellerRequest struct {
	Cid         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   int    `json:"telephone"`
	LocalityID  int    `json:"locality_id"`
}

func (h *SellerHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &SellerRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			customResponse.Error(w, defaultErrors.ErrorBadRequest)
			return
		}

		sDoc := models.NewSellerDoc(-1, req.Cid, req.CompanyName, req.Address, req.Telephone, req.LocalityID)
		new, err := h.sv.Create(*sDoc)

		// Error handling
		if err != nil {
			customResponse.Error(w, err)
			return
		}

		customResponse.JSON(w, http.StatusCreated, new)
	}
}

func (h *SellerHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			customResponse.Error(w, defaultErrors.ErrorBadRequest) // BadRequest
			return
		}

		s, err := h.sv.GetByID(id)

		// Error handling
		if err != nil {
			customResponse.Error(w, err)
			return
		}

		customResponse.JSON(w, http.StatusOK, s)
	}
}

func (h *SellerHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			customResponse.Error(w, defaultErrors.ErrorBadRequest) // BadRequest
			return
		}

		// Error handling
		if err := h.sv.Delete(id); err != nil {
			customResponse.Error(w, err)
			return
		}

		// Deleted Successfully
		customResponse.JSON(w, http.StatusNoContent, nil)

	}
}

type UpdateSellerRequest struct {
	Cid         *int    `json:"cid"`
	CompanyName *string `json:"company_name"`
	Address     *string `json:"address"`
	Telephone   *int    `json:"telephone"`
	LocalityID  *int    `json:"locality_id"`
}

func (h *SellerHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			customResponse.Error(w, defaultErrors.ErrorBadRequest)
			return
		}

		req := &UpdateSellerRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			customResponse.Error(w, defaultErrors.ErrorBadRequest)
			return
		}

		sellerDoc, err := h.sv.Update(id, req.Cid, req.CompanyName, req.Address, req.Telephone, req.LocalityID)

		// Error handling
		if err != nil {
			customResponse.Error(w, err)
			return
		}

		customResponse.JSON(w, http.StatusOK, sellerDoc)
	}
}
