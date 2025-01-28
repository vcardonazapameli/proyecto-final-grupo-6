package routes

import (
	"fmt"

	"github.com/go-chi/chi/v5"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers"
	loader "github.com/arieleon_meli/proyecto-final-grupo-6/internal/loader/section"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/section"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/section"
)

func RegisterSectionRoutes(r chi.Router) {

	ld := loader.NewSectionJSONFile("docs/section.json")
	db, err := ld.Load()
	if err != nil {
		fmt.Print("Error: ", err.Error())
		return
	}

	rp := repository.NewSectionMap(db)
	sv := service.NewSectionDefault(rp)
	hd := handler.NewSectionHandler(sv)

	r.Route("/section", func(rt chi.Router) {
		println(hd)
	})
}
