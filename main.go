package main

import (
	"log"
	"net/http"

	"github.com/antihax/mock-esi/swaggerServer"
)

func main() {
	log.Printf("Server started")

	router := swaggerServer.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
