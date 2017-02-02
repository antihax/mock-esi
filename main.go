package main

import (
	"log"
	"net/http"

	_ "github.com/antihax/mock-esi/dev"
	_ "github.com/antihax/mock-esi/latest"
	_ "github.com/antihax/mock-esi/legacy"
	"github.com/antihax/mock-esi/mockesi"
	_ "github.com/antihax/mock-esi/v1"
	_ "github.com/antihax/mock-esi/v2"
	_ "github.com/antihax/mock-esi/v3"
	_ "github.com/antihax/mock-esi/v4"
	//	_ "github.com/antihax/mock-esi/v5"
	//	_ "github.com/antihax/mock-esi/v6"
	//	_ "github.com/antihax/mock-esi/v7"
)

func main() {
	log.Printf("Server started")

	router := mockesi.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
