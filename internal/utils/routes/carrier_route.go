package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers/carrier"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/carrier"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/carrier"
)

func RegisterCarrierRoutes(r chi.Router, database *sql.DB) {
	repository := repository.NewCarrierRepository(database)
	service := service.NewCarrierService(repository)
	handler := handler.NewCarrierHandler(service)

	// - endpoints
	r.Route("/carries", func(r chi.Router) {
		r.Post("/", handler.Create())
	})
}
