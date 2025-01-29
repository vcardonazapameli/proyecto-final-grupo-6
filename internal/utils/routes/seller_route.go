package routes

import (
	"fmt"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/seller"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/seller"
	loader "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/loader/seller"
	"github.com/go-chi/chi/v5"
)

func RegisterSellerRoutes(r chi.Router) {

	ld := loader.NewSellerJSONFile("docs/seller.json")
	// datos cargados
	db, err := ld.Load()
	if err != nil {
		fmt.Print("Error: ", err.Error())
		return
	}

	// //dependency injection
	rp := repository.NewSellerRepositoryJSON(db, ld)
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
