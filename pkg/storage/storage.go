package storage

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)

func Connect(driver string, dns string) (*sql.DB, error) {
	db, err := sql.Open(driver, dns)
	if err != nil {
		// TODO: should we use fatal? Since it will exit the program
		log.Fatalf("Failed to connect to %s with DNS: %s. %s", driver, dns, err)
	}
	return db, nil
}
