package routes

import (
	"fmt"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/buyer"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/buyer"
	loader "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/loader/buyer"

	"github.com/go-chi/chi/v5"
)

func RegisterBuyerRoutes(r chi.Router) {

	ld := loader.NewBuyerJsonFile("docs/buyer.json")
	// datos cargados
	db, err := ld.Load()
	if err != nil {
		fmt.Print("Error: ", err.Error())
		return
	}

	// //dependency injection
	rp := repository.NewBuyerMap(db)
	sv := service.NewBuyerDefault(rp)
	hd := handler.NewBuyerHandler(sv)

	// // - middlewares

	// - endpoints
	r.Route("/buyer", func(rt chi.Router) {
		// - GET /vehicles
		rt.Get("/", hd.GetAll())
		rt.Get("/{id}", hd.GetById())
		rt.Post("/", hd.CreateBuyer())
		rt.Delete("/{id}", hd.DeleteBuyer())
		rt.Patch("/{id}",hd.UpdateBuyer())
	})
}
