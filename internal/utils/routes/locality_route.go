package routes

import (
	"database/sql"

	handler "github.com/arieleon_meli/proyecto-final-grupo-6/internal/handlers"
	repository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/locality"
	service "github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/locality"
	"github.com/go-chi/chi/v5"
)

func RegisterLocalityRoutes(r chi.Router, database *sql.DB) {
	repository := repository.NewLocalityRepositoryDB(database)
	service := service.NewLocalityServiceDefault(repository)
	handler := handler.NewLocalityHandler(service)

	r.Route("/localities", func(r chi.Router) {
		r.Post("/", handler.Create())
	})
}
