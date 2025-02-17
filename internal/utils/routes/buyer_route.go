package routes

import (
	"database/sql"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/buyer"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/buyer"
	
	"github.com/go-chi/chi/v5"
)

func RegisterBuyerRoutes(r chi.Router, database *sql.DB) {

	
	// //dependency injection
	rp := repository.NewBuyerRepository(database)
	sv := service.NewBuyerDefault(rp)
	hd := handler.NewBuyerHandler(sv)

	// // - middlewares

	// - endpoints
	r.Route("/buyer", func(rt chi.Router) {
		// - GET /vehicles
		rt.Get("/", hd.GetAll())
		rt.Get("/{id}", hd.GetById())
		rt.Get("/reportPurchaseOrders",hd.GetPurchaseOrderReports())
		rt.Post("/", hd.CreateBuyer())
		rt.Delete("/{id}", hd.DeleteBuyer())
		rt.Patch("/{id}",hd.UpdateBuyer())
		
	})
}
