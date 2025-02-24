package handlers

import (
	"encoding/json"
	"net/http"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/inbound_order"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/response"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewInboundOrderHandler(sv service.InboundOrderService) *InboundOrderHandler {
	return &InboundOrderHandler{sv: sv}
}

// EmployeeDefault is a struct with methods that represent handlers for Employees
type InboundOrderHandler struct {
	// sv is the service that will be used by the handler
	sv service.InboundOrderService
}

func (h *InboundOrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newInboundOrder models.RequestInboundOrder
		if err := json.NewDecoder(r.Body).Decode(&newInboundOrder); err != nil {
			response.Error(w, customErrors.ErrorBadRequest)
			return
		}

		data, err := h.sv.Create(newInboundOrder)
		if err != nil {
			response.Error(w, err)
			return
		}

		response.JSON(w, http.StatusCreated, data)
	}
}
