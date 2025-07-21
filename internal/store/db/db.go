package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	connectString := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	connectInput := fmt.Sprintf(
		connectString,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	driver := os.Getenv("DB_DRIVER")
	db, err := sql.Open(driver, connectInput)
	if err != nil {
		return nil, err
	}
	return db, nil
}
