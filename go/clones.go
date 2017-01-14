package swaggerServer

import (
	"net/http"

)

func GetCharactersCharacterIdClones(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "home_location" : {
    "location_id" : 1021348135816,
    "location_type" : "structure"
  },
  "jump_clones" : [ {
    "implants" : [ 22118 ],
    "location_id" : 60003463,
    "location_type" : "station"
  }, {
    "implants" : [ ],
    "location_id" : 1021348135816,
    "location_type" : "structure"
  } ]
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


