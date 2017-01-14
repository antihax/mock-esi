package swaggerServer

import (
	"net/http"

)

func GetCharactersCharacterIdAssets(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "is_singleton" : true,
  "item_id" : 1000000016835,
  "location_flag" : "Hangar",
  "location_id" : 60002959,
  "location_type" : "station",
  "type_id" : 3516
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


