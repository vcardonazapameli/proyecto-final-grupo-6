package handlers

import (
	"net/http"
	"strconv"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/section"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

func NewSectionHandler(sv service.SectionService) *SectionHandler {
	return &SectionHandler{sv: sv}
}

type SectionHandler struct {
	sv service.SectionService
}

func (h *SectionHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := h.sv.GetAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "An error occurred while getting sections")
			return
		}

		data := mappers.SectionToSectionDoc(s)

		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})

	}
}

func (h *SectionHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" || id == "0" {
			response.Error(w, http.StatusBadRequest, "ID is required")
			return
		}

		idConv, err := strconv.Atoi(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid ID")
			return
		}

		s, err := h.sv.GetByID(idConv)
		if err != nil {
			response.Error(w, http.StatusNotFound, errors.ErrorNotFound.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Section found",
			"data":    mappers.SectionToSectionDoc(map[int]models.Section{idConv: s}),
		})
	}
}
