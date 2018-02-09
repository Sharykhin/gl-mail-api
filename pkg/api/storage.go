package api

import (
	"database/sql"
	"github.com/Sharykhin/gl-mail-api/pkg/storage"
)

func getStorage () *sql.DB {
	db, _ := storage.Connect(config["driver"], config["dns"])
	return db
}
