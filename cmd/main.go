package main

import (
	"fmt"
	"log"

	"github.com/arieleon_meli/proyecto-final-grupo-6/cmd/server"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Errorf("error loading config: %v", err)
	}

	log.Println("Starting server on :" + cfg.ServerAddress)
	app := server.NewServerChi(cfg)
	// - run
	if err := app.Run(*cfg); err != nil {
		fmt.Println(err)
		return
	}
}
