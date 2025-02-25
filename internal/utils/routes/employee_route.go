package routes

import (
	"database/sql"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers/employee"
	rpEmployee "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/employee"
	rpInboundOrders "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/inbound_order"
	rpWarehouse "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/warehouse"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/employee"

	"github.com/go-chi/chi/v5"
)

func RegisterEmployeeRoutes(r chi.Router, db *sql.DB) {

	//dependency injection
	empRepo := rpEmployee.NewEmployeeMap(db)
	iorRepo := rpInboundOrders.NewInboundOrderMap(db)
	whRepo := rpWarehouse.NewWarehouseRepository(db)
	sv := service.NewEmployeeDefault(empRepo, iorRepo, whRepo)
	hd := handler.NewEmployeeHandler(sv)

	// - middlewares

	// - endpoints
	r.Route("/employee", func(rt chi.Router) {
		// - GET /vehicles
		rt.Get("/", hd.GetAll())
		rt.Get("/{id}", hd.GetById())
		rt.Get("/reportInboundOrders", hd.GetReportInboundOrders())
		rt.Post("/", hd.Create())
		rt.Patch("/{id}", hd.Update())
		rt.Delete("/{id}", hd.Delete())
	})
}
