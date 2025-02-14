package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_batch"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/product_batch"
)

func RegisterProductBatchRoutes(r chi.Router, database *sql.DB) {
	repository := repository.NewProductBatchRepositoryDB(database)
	service := service.NewProductBatchService(repository)
	handler := handler.NewProductBatchHandler(service)

	r.Route("/productBatches", func(r chi.Router) {
		r.Post("/", handler.Create())
	})
}
