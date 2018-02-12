package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sharykhin/gl-mail-api/handler"
	"github.com/joho/godotenv"
)

func init() {
	// TODO: keep in mind the the order of init can't guarantee that all this one will be called before others
	// TODO: I guess this package can be removed (godotenv)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Println("Server is listening on port 8002")
	log.Fatal(http.ListenAndServe(":8002", handler.Handler()))
}
