package main

import (
	"log"
	"net/http"
	"fmt"

	"github.com/Sharykhin/gl-mail-api/handler"
)

func main() {
	fmt.Println("Server is listening on port 8002")
	log.Fatal(http.ListenAndServe(":8002", handler.Handler()))
}
