package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/section"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/section"
)

func RegisterSectionRoutes(r chi.Router, db *sql.DB) {

	rp := repository.NewSectionMap(db)
	sv := service.NewSectionDefault(rp)
	hd := handler.NewSectionHandler(sv)

	r.Route("/sections", func(rt chi.Router) {
		rt.Get("/", hd.GetAll())
		rt.Get("/{id}", hd.GetByID())
		rt.Post("/", hd.Create())
		rt.Patch("/{id}", hd.Update())
		rt.Delete("/{id}", hd.Delete())
	})
}
