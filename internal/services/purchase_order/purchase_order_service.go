package purchaseorder

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"


type PurchaseOrderService interface {
	CreatePurchaseOrder(purchaseOrder models.PurchaseOrderRequest)(*models.PurchaseOrderResponse, error)
	
}