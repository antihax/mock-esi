package main

import (
	"github.com/antihax/mock-esi/swaggerServer"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	router := swaggerServer.NewRouter()
	
	log.Fatal(http.ListenAndServe(":8080", router))
}
