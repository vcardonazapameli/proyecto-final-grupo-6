package routes

import (
	"database/sql"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/inbound_order"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/inbound_order"
	"github.com/go-chi/chi/v5"
)

func RegisterInboundOrderRoutes(r chi.Router, db *sql.DB) {

	//dependency injection
	rp := repository.NewInboundOrderMap(db)
	sv := service.NewInboundOrderDefault(rp)
	hd := handler.NewInboundOrderHandler(sv)
	// // - middlewares

	// - endpoints
	r.Route("/inboundOrders", func(rt chi.Router) {
		rt.Post("/", hd.Create())
	})
}
