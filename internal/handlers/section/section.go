package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/section"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/response"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/bootcamp-go/web/request"
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
		sections, err := h.sv.GetAll()
		if err != nil {
			response.Error(w, customErrors.ErrorNotFound)
			return
		}
		response.JSON(w, http.StatusOK, sections)

	}
}

func (h *SectionHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, customErrors.ErrorBadRequest)
			return
		}
		section, err := h.sv.GetByID(id)
		if err != nil {
			response.Error(w, customErrors.ErrorNotFound)
			return
		}
		response.JSON(w, http.StatusOK, section)
	}
}

func (h *SectionHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		section := models.SectionDocRequest{}
		if err := json.NewDecoder(r.Body).Decode(&section); err != nil {
			response.Error(w, customErrors.ErrorBadRequest)
			return
		}
		createdSection, err := h.sv.Create(section)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusCreated, createdSection)
	}
}

func (h *SectionHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idConv, err := strconv.Atoi(chi.URLParam(r, "id"))
		var sectionDoc models.UpdateSectionDto
		if err != nil {
			response.Error(w, customErrors.ErrorBadRequest)
			return
		}
		if err := request.JSON(r, &sectionDoc); err != nil {
			response.Error(w, customErrors.ErrorBadRequest)
			return
		}
		updatedSection, err := h.sv.Update(idConv, sectionDoc)

		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusOK, updatedSection)
	}
}

func (h *SectionHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idConv, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, customErrors.ErrorBadRequest)
			return
		}

		err = h.sv.Delete(idConv)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusNoContent, nil)
	}
}

func (handler *SectionHandler) GetSectionReports() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		var sectionId = 0
		var err error
		if idStr != "" {
			sectionId, err = strconv.Atoi(idStr)
			if err != nil {
				response.Error(w, customErrors.ErrorBadRequest)
				return
			}
		}
		sectionReports, err := handler.sv.GetSectionReports(sectionId)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, http.StatusOK, sectionReports)
	}
}
