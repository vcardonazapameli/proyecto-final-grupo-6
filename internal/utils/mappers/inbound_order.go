package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

func RequestInboundOrderToInboundOrder(request models.RequestInboundOrder) models.InboundOrder {
	return models.InboundOrder{
		OrderDate:      request.OrderDate,
		OrderNumber:    request.OrderNumber,
		EmployeeID:     request.EmployeeID,
		ProductBatchID: request.ProductBatchID,
		WarehouseID:    request.WarehouseID,
	}
}
