package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers/product_type"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_type"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/product_type"
)

func RegisterProductTypeRoutes(r chi.Router, database *sql.DB) {
	repository := repository.NewProductTypeRepository(database)
	service := service.NewProductTypeService(repository)
	handler := handler.NewProductTypeHandler(service)

	r.Route("/productTypes", func(r chi.Router) {
		r.Get("/", handler.GetAll())
		r.Get("/{id}", handler.GetById())
		r.Post("/", handler.Create())
	})
}
