package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/config"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(cfg *config.Config) *sql.DB {
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("mysql", datasource)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database!")
	return db
}
