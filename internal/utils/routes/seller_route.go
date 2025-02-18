package routes

import (
	"database/sql"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/seller"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/seller"
	"github.com/go-chi/chi/v5"
)

func RegisterSellerRoutes(r chi.Router, db *sql.DB) {

	//dependency injection
	rp := repository.NewSellerRepositoryDB(db)
	sv := service.NewSellerServiceDefault(rp)
	hd := handler.NewSellerHandler(sv)

	// // - middlewares

	// - endpoints
	r.Route("/sellers", func(rt chi.Router) {
		rt.Get("/", hd.GetAll())
		rt.Get("/{id}", hd.GetByID())
		rt.Post("/", hd.Create())
		rt.Delete("/{id}", hd.Delete())
		rt.Patch("/{id}", hd.Update())
	})
}
