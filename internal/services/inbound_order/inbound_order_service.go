package inbound_order

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type InboundOrderService interface {
	Create(request models.RequestInboundOrder) (*models.InboundOrder, error)
}
