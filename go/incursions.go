package swaggerServer

import (
	"net/http"

)

func GetIncursions(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "constellation_id" : 20000607,
  "faction_id" : 500019,
  "has_boss" : true,
  "infested_solar_systems" : [ 30004148, 30004149, 30004150, 30004151, 30004152, 30004153, 30004154 ],
  "influence" : 1.0,
  "staging_solar_system_id" : 30004154,
  "state" : "mobilizing",
  "type" : "Incursion"
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


