package routes

import (
	"fmt"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers"
	loader "github.com/arieleon_meli/proyecto-final-grupo-6/internal/loader/seller"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/seller"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/seller"
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
	r.Route("/seller", func(rt chi.Router) {
		rt.Get("/", hd.GetAll())
	})
}
