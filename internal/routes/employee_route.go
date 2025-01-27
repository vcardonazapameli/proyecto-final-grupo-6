package routes

import (
	"fmt"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers"
	loader "github.com/arieleon_meli/proyecto-final-grupo-6/internal/loader/employee"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/employee"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/employee"

	"github.com/go-chi/chi/v5"
)

func RegisterEmployeeRoutes(r chi.Router) {

	ld := loader.NewEmployeeJSONFile("docs/employee.json")
	// datos cargados
	db, err := ld.Load()
	if err != nil {
		fmt.Print("Error: ", err.Error())
		return
	}

	// //dependency injection
	rp := repository.NewEmployeeMap(db)
	sv := service.NewEmployeeDefault(rp)
	hd := handler.NewEmployeeHandler(sv)

	// // - middlewares

	// - endpoints
	r.Route("/employee", func(rt chi.Router) {
		// - GET /vehicles
		rt.Get("/", hd.GetAll())
	})
}
