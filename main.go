package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Sharykhin/gl-mail-api/handler"
)

func main() {
	address := os.Getenv("HTTP_ADDRESS")
	fmt.Printf("Server is listening on %s\n", address)
	log.Fatal(http.ListenAndServe(address, handler.Handler()))
}
