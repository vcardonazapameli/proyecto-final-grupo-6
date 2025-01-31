package main

import (
	"fmt"

	"github.com/arieleon_meli/proyecto-final-grupo-6/cmd/server"
)

func main() {

	//App config
	cfg := &server.ConfigServerChi{
		ServerAddress: ":8080",
	}

	app := server.NewServerChi(cfg)

	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
