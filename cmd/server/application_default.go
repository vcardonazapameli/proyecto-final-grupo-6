package server

import (
	"net/http"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/routes"
	"github.com/go-chi/chi/v5"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
}

// NewServerChi is a function that returns a new instance of ServerChi
func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	// default values
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
	}

	return &ServerChi{
		serverAddress: defaultConfig.ServerAddress,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
}

// Run is a method that runs the server
func (a *ServerChi) Run() (err error) {

	// router
	r := chi.NewRouter()

	// - middlewares

	// - routes
	routes.RegisterEmployeeRoutes(r)
	routes.RegisterSectionRoutes(r)

	// run server
	err = http.ListenAndServe(a.serverAddress, r)
	return
}
