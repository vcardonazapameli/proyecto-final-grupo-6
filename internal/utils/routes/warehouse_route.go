package routes

import (
	"fmt"

	"github.com/go-chi/chi/v5"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/warehouse"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/warehouse"
	loader "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/loader/warehouse"
)

func RegisterWarehouseRoutes(r chi.Router) {

	ld := loader.NewWarehouseJSONFile("docs/warehouse.json")
	db, err := ld.Load()
	if err != nil {
		fmt.Print("Error: ", err.Error())
		return
	}

	rp := repository.NewWarehouseMap(db)
	sv := service.NewWarehouseDefault(rp)
	hd := handler.NewWarehouseHandler(sv)

	r.Route("/warehouse", func(rt chi.Router) {
		println(hd)
	})
}
