package handlers

import (
	"net/http"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/section"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/bootcamp-go/web/response"
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
