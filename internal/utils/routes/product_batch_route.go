package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers"
	product "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_batch"
	section "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/section"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/product_batch"
)

func RegisterProductBatchRoutes(r chi.Router, database *sql.DB) {
	productBatchRepository := repository.NewProductBatchRepository(database)
	sectionRepository := section.NewSectionMap(database)
	productRepository := product.NewProductRepository(database)
	productBatchService := service.NewProductBatchService(productBatchRepository, sectionRepository, productRepository)
	handler := handler.NewProductBatchHandler(productBatchService)

	r.Route("/productBatches", func(r chi.Router) {
		r.Post("/", handler.Create())
	})
}
