package handlers

import (
	"encoding/json"
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
		sections, err := h.sv.GetAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "An error occurred while getting sections")
			return
		}

		sectionDocs := make(map[int]models.SectionDoc)
		for id, section := range sections {
			sectionDocs[id] = mappers.SectionToSectionDoc(section)
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": sectionDocs,
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

		sectionDoc := mappers.SectionToSectionDoc(s)
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Section found",
			"data":    sectionDoc,
		})
	}
}

func (h *SectionHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var s models.SectionDoc
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		section := mappers.SectionDocToSection(s)
		createdSection, err := h.sv.Create(section)
		if err != nil {
			switch err {
			case errors.ErrorUnprocessableContent:
				response.Error(w, http.StatusUnprocessableEntity, err.Error())
			case errors.ErrorConflict:
				response.Error(w, http.StatusConflict, err.Error())
			default:
				response.Error(w, http.StatusInternalServerError, "An error occurred while creating the section")
			}
			return
		}

		createdSectionDoc := mappers.SectionToSectionDoc(createdSection)
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Section created",
			"data":    createdSectionDoc,
		})
	}
}

func (h *SectionHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idConv, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid ID")
			return
		}
		var s models.SectionDoc
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		section := mappers.SectionDocToSection(s)
		updatedSection, err := h.sv.Update(idConv, section)
		if err != nil {
			if err == errors.ErrorNotFound {
				response.Error(w, http.StatusNotFound, err.Error())
			} else {
				response.Error(w, http.StatusInternalServerError, "An error occurred while updating the section")
			}
			return
		}
		updateSectionDoc := mappers.SectionToSectionDoc(updatedSection)
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Section updated",
			"data":    updateSectionDoc,
		})
	}
}

func (h *SectionHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idConv, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid ID")
			return
		}

		err = h.sv.Delete(idConv)
		if err != nil {
			if err == errors.ErrorNotFound {
				response.Error(w, http.StatusNotFound, err.Error())
			} else {
				response.Error(w, http.StatusInternalServerError, "An error occurred while deleting the section")
			}
			return
		}
		response.JSON(w, http.StatusNoContent, map[string]any{
			"message": "Section deleted",
		})
	}
}
