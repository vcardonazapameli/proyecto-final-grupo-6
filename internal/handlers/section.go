package handlers

import (
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/section"
)

func NewSectionHandler(sv service.SectionService) *SectionHandler {
	return &SectionHandler{sv: sv}
}

type SectionHandler struct {
	sv service.SectionService
}
