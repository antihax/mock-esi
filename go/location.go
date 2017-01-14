package swaggerServer

import (
	"net/http"

)

func GetCharactersCharacterIdLocation(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "solar_system_id" : 30002505,
  "structure_id" : 1000000016989
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCharactersCharacterIdShip(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "ship_item_id" : 1000000016991,
  "ship_name" : "SPACESHIPS!!!",
  "ship_type_id" : 1233
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


