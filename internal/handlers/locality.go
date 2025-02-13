package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/locality"
	defaultErrors "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/response"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type LocalityHandler struct {
	sv service.LocalityService
}

func NewLocalityHandler(sv service.LocalityService) *LocalityHandler {
	return &LocalityHandler{sv}
}

func (h *LocalityHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		newLocality := models.LocalityDoc{}
		if err := json.NewDecoder(r.Body).Decode(&newLocality); err != nil {
			response.Error(w, defaultErrors.ErrorBadRequest)
			return
		}

		err := h.sv.Create(&newLocality)
		if err != nil {
			response.Error(w, err)
			return
		}

		response.JSON(w, 201, newLocality)
		return
	}
}

func (h *LocalityHandler) GetSellerByLocalityID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")

		// If no ID Query param is found, return every locality count
		if idStr == "" {
			body, err := h.sv.GetAllSellerCountByLocalityID()
			if err != nil {
				response.Error(w, err)
				return
			}
			response.JSON(w, 200, body)
			return

		}

		// If ID Query param is found, return that locality count
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.Error(w, defaultErrors.ErrorBadRequest)
			return
		}

		body, err := h.sv.GetSellerCountByLocalityID(id)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, 200, body)
	}
}
