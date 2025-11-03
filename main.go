package main

import (
	"log"
	"net/http"

	"github.com/antihax/mock-esi/mockesi"
)

func main() {
	log.Printf("Server started")

	router := mockesi.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
