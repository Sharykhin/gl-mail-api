package api

import (
	"os"
)

const (
	SQL_DRIVER = "mysql"
)

var config  = map[string]string{
	"driver": SQL_DRIVER,
	"dns": os.Getenv("MYSQL_DNS"),
}







