package purchaseorder

import "github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"

type PurchaseOrderRepository interface {
	CreatePurchaseOrder(purchaseOrder *models.PurchaseOrderResponse) error
	ValidateIfOrderStatusExist(orderStatusId int) bool
	ValidateIfOrderNumberExist(orderNumber string) bool
}
