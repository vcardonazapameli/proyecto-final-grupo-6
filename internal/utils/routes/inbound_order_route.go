package routes

import (
	"database/sql"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers"
	empRepository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/employee"
	iorRepository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/inbound_order"
	pbRepository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_batch"
	whRepository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/warehouse"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/inbound_order"
	"github.com/go-chi/chi/v5"
)

func RegisterInboundOrderRoutes(r chi.Router, db *sql.DB) {

	//dependency injection
	iorRepo := iorRepository.NewInboundOrderMap(db)
	empRepo := empRepository.NewEmployeeMap(db)
	pbRepo := pbRepository.NewProductBatchMap(db)
	whRepo := whRepository.NewWarehouseRepository(db)

	sv := service.NewInboundOrderDefault(iorRepo, empRepo, pbRepo, whRepo)
	hd := handler.NewInboundOrderHandler(sv)
	// // - middlewares

	// - endpoints
	r.Route("/inboundOrders", func(rt chi.Router) {
		rt.Post("/", hd.Create())
	})
}
