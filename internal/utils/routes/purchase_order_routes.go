package routes

import (
	"database/sql"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers/purchase_order"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/buyer"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/carrier"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/purchase_order"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/warehouse"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/purchase_order"

	"github.com/go-chi/chi/v5"
)

func RegisterPurchaseOrderRoutes(r chi.Router, database *sql.DB) {

	// //dependency injection
	purchaseOrderrepo := repository.NewPurchaseOrderRepository(database)
	warehouseRepo := warehouse.NewWarehouseRepository(database)
	buyerRepo := buyer.NewBuyerRepository(database)
	carrierRepo := carrier.NewCarrierRepository(database)
	sv := service.NewPurchaseOrderService(purchaseOrderrepo, warehouseRepo, buyerRepo, carrierRepo)
	hd := handler.NewPurchaseOrderHandler(sv)

	// // - middlewares

	// - endpoints
	r.Route("/purchaseOrders", func(rt chi.Router) {
		// - GET /purchaseOrders
		rt.Post("/", hd.CreatePurchaseOrder())
	})
}
