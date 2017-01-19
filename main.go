package main

import (
	"github.com/antihax/mock-esi/swServer"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	router := swServer.NewRouter()
	
	log.Fatal(http.ListenAndServe(":8080", router))
}
