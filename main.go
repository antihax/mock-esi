package main

import (
	"log"
	"net/http"

	_ "github.com/antihax/mock-esi/dev/go"
	_ "github.com/antihax/mock-esi/latest/go"
	_ "github.com/antihax/mock-esi/legacy/go"
	_ "github.com/antihax/mock-esi/v1/go"
	_ "github.com/antihax/mock-esi/v2/go"
	_ "github.com/antihax/mock-esi/v3/go"
	_ "github.com/antihax/mock-esi/v4/go"
	_ "github.com/antihax/mock-esi/v5/go"
	_ "github.com/antihax/mock-esi/v6/go"

	"github.com/antihax/mock-esi/mockesi"
)

func main() {
	log.Printf("Server started")

	router := mockesi.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
