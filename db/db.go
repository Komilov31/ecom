package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewSqlStorage(cfg string) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
