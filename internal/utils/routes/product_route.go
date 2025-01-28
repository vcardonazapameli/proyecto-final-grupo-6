package routes

import (
	"fmt"

	"github.com/go-chi/chi/v5"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/product"
	loader "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/loader/product"
)

func RegisterProductRoutes(r chi.Router) {

	ld := loader.NewProductJSONFile("docs/product.json")
	db, err := ld.Load()
	if err != nil {
		fmt.Print("Error: ", err.Error())
		return
	}

	rp := repository.NewProductMap(db)
	sv := service.NewProductDefault(rp)
	hd := handler.NewProductHandler(sv)

	r.Route("/products", func(r chi.Router) {
		r.Get("/", hd.GetAll())
		r.Get("/{id}", hd.GetById())
	})
}
