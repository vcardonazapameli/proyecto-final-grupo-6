package mappers

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"


func PurchaseOrderResponseToRequest(response models.PurchaseOrderResponse)models.PurchaseOrderRequest{
	return models.PurchaseOrderRequest{
		OrderNumber: response.OrderNumber,
		OrderDate: response.OrderDate,
		TrackingCode: response.TrackingCode,
		BuyerId: response.BuyerId,
		CarrierId: response.CarrierId,
		OrderStatusId: response.OrderStatusId,
		WarehouseId: response.WarehouseId,
	}
}

func PurchaseOrderRequestToResponse(request models.PurchaseOrderRequest)models.PurchaseOrderResponse{
	return models.PurchaseOrderResponse{
		OrderNumber: request.OrderNumber,
		OrderDate: request.OrderDate,
		TrackingCode: request.TrackingCode,
		BuyerId: request.BuyerId,
		CarrierId: request.CarrierId,
		OrderStatusId: request.OrderStatusId,
		WarehouseId: request.WarehouseId,
	}
}