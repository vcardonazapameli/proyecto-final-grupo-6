package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/section"
	defaultErrors "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/response"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
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
			response.Error(w, err)
			return
		}

		sectionDocs := make(map[int]models.SectionDoc)
		for id, section := range sections {
			sectionDocs[id] = mappers.SectionToSectionDoc(section)
		}
		response.JSON(w, http.StatusOK, sectionDocs)

	}
}

func (h *SectionHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" || id == "0" {
			response.Error(w, defaultErrors.ErrorBadRequest)
			return
		}

		idConv, err := strconv.Atoi(id)
		if err != nil {
			response.Error(w, defaultErrors.ErrorBadRequest)
			return
		}

		s, err := h.sv.GetByID(idConv)
		if err != nil {
			response.Error(w, err)
			return
		}

		sectionDoc := mappers.SectionToSectionDoc(s)
		response.JSON(w, http.StatusOK, sectionDoc)
	}
}

func (h *SectionHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var s models.SectionDoc
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			response.Error(w, defaultErrors.ErrorBadRequest)
			return
		}

		section := mappers.SectionDocToSection(s)
		createdSection, err := h.sv.Create(section)
		if err != nil {
			response.Error(w, err)
			return
		}
		createdSectionDoc := mappers.SectionToSectionDoc(createdSection)
		response.JSON(w, http.StatusCreated, createdSectionDoc)
	}
}

func (h *SectionHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idConv, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, defaultErrors.ErrorBadRequest)
			return
		}
		var s models.SectionDoc
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			response.Error(w, defaultErrors.ErrorBadRequest)
			return
		}

		section := mappers.SectionDocToSection(s)
		updatedSection, err := h.sv.Update(idConv, section)
		if err != nil {
			response.Error(w, err)
			return
		}
		updateSectionDoc := mappers.SectionToSectionDoc(updatedSection)
		response.JSON(w, http.StatusOK, updateSectionDoc)
	}
}

func (h *SectionHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idConv, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, defaultErrors.ErrorBadRequest)
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
