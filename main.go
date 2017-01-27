package main

import (
	"github.com/antihax/mock-esi/swS"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	router := swS.NewRouter()
	
	log.Fatal(http.ListenAndServe(":8080", router))
}
