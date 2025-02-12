package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

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

		newLocality := models.Locality{}
		if err := json.NewDecoder(r.Body).Decode(&newLocality); err != nil {
			response.Error(w, defaultErrors.ErrorBadRequest)
			return
		}

		err := h.sv.Create(&newLocality)
		if err != nil {
			fmt.Println(err.Error())
			response.Error(w, err)
			return
		}

		response.JSON(w, 201, newLocality)
		return
	}
}
