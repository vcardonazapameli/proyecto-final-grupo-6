package handlers

import (
	"net/http"

	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/purchase_order"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/response"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/bootcamp-go/web/request"
)
func NewPurchaseOrderHandler(sv service.PurchaseOrderService )*PurchaseOrderHandler{
	return &PurchaseOrderHandler{sv: sv}
}
type PurchaseOrderHandler struct {
	sv service.PurchaseOrderService
}


func (handler *PurchaseOrderHandler)CreatePurchaseOrder() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		purchaseOrderRequest := models.PurchaseOrderRequest{}
		if err:= request.JSON(r,&purchaseOrderRequest); err != nil{
			response.Error(w, customErrors.ErrorBadRequest)
			return
		}
		createdPurchaseOrder, err := handler.sv.CreatePurchaseOrder(purchaseOrderRequest)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w,http.StatusCreated, createdPurchaseOrder)
	}
}


