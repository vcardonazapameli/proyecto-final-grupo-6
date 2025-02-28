package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers/product_record"
	productRepository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_record"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/product_record"
)

func RegisterProductRecordRoutes(r chi.Router, database *sql.DB) {
	repository := repository.NewProductRecordRepository(database)
	productRepository := productRepository.NewProductRepository(database)
	service := service.NewProductRecordService(repository, productRepository)
	handler := handler.NewProductRecordHandler(service)

	r.Route("/productRecords", func(r chi.Router) {
		r.Post("/", handler.Create())
	})
}
