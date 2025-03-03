package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers/product"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product"
	productTyperepository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_type"
	sellerRepository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/seller"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/product"
)

func RegisterProductRoutes(r chi.Router, database *sql.DB) {
	repository := repository.NewProductRepository(database)
	productTypeRepository := productTyperepository.NewProductTypeRepository(database)
	sellerRepository := sellerRepository.NewSellerRepositoryDB(database)
	service := service.NewProductService(repository, productTypeRepository, sellerRepository)
	handler := handler.NewProductHandler(service)

	r.Route("/products", func(r chi.Router) {
		r.Get("/", handler.GetAll())
		r.Get("/{id}", handler.GetById())
		r.Get("/productRecords", handler.GetProductRecords())
		r.Delete("/{id}", handler.Delete())
		r.Post("/", handler.Create())
		r.Patch("/{id}", handler.Update())
	})
}
