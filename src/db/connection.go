package db

import (
	"database/sql"
	"fmt"
	"os"
)

func ConnectDatabase() (*sql.DB, error) {
	NAME_DATABASE := os.Getenv("NAME_DATABASE")
	connection := fmt.Sprintf("postgresql://root@gabrielortega:26257/%s?sslmode=disable", NAME_DATABASE)
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return db, err
	}
	return db, nil
}
