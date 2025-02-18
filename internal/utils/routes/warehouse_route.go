package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/warehouse"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/warehouse"
)

func RegisterWarehouseRoutes(r chi.Router, database *sql.DB) {
	rp := repository.NewWarehouseRepository(database)
	sv := service.NewWarehouseService(rp)
	hd := handler.NewWarehouseHandler(sv)

	r.Route("/warehouses", func(r chi.Router) {
		r.Get("/", hd.GetAll())
		r.Get("/{id}", hd.GetById())
		r.Post("/", hd.CreateWarehouse())
		r.Delete("/{id}", hd.DeleteWarehouse())
		r.Patch("/{id}", hd.UpdateWarehouse())
	})
}
