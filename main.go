package main

import (
	"github.com/antihax/mock-esi/swaggerServer"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	router := swaggerServer.NewRouter()
	
	log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", router))
}
