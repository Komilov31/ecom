package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Komilov31/ecom/cmd/api"
	"github.com/Komilov31/ecom/config"
	"github.com/Komilov31/ecom/db"
)

func main() {

	DBConfig := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Envs.DBHost,
		config.Envs.DBPort,
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBName,
	)

	db, err := db.NewSqlStorage(DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")
}
