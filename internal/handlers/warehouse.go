package handlers

import (
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/warehouse"
)

func NewWarehouseHandler(sv service.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{sv: sv}
}

type WarehouseHandler struct {
	sv service.WarehouseService
}
