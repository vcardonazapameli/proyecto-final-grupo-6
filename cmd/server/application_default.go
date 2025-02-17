package server

import (
	"net/http"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/config"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/db"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/routes"
	"github.com/go-chi/chi/v5"
)

type ConfigServerChi struct {
	ServerAddress string
}

func NewServerChi(cfg *config.Config) *ServerChi {
	defaultConfig := &config.Config{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.DBHost != "" {
			defaultConfig.DBHost = cfg.DBHost
		}
		if cfg.DBPort != "" {
			defaultConfig.DBPort = cfg.DBPort
		}
		if cfg.DBUser != "" {
			defaultConfig.DBUser = cfg.DBUser
		}
		if cfg.DBPassword != "" {
			defaultConfig.DBPassword = cfg.DBPassword
		}
		if cfg.DBName != "" {
			defaultConfig.DBName = cfg.DBName
		}
	}

	return &ServerChi{
		serverAddress: defaultConfig.ServerAddress,
		config:        defaultConfig,
	}
}

type ServerChi struct {
	serverAddress string
	config        *config.Config
}

func (a *ServerChi) Run(cfg config.Config) (err error) {
	database := db.ConnectDB(&cfg)
	defer database.Close()

	// router
	r := chi.NewRouter()

	// - middlewares

	// - routes

	routes.RegisterWarehouseRoutes(r)
	routes.RegisterEmployeeRoutes(r, database)
	routes.RegisterInboundOrderRoutes(r, database)
	routes.RegisterSellerRoutes(r)
	routes.RegisterSectionRoutes(r)
	routes.RegisterProductRoutes(r, database)
	routes.RegisterBuyerRoutes(r)

	// run server
	err = http.ListenAndServe(a.serverAddress, r)
	return
}
