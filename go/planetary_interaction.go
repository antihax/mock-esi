package swaggerServer

import (
	"net/http"

)

func GetCharactersCharacterIdPlanets(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "last_update" : "2016-11-28T16:42:51Z",
  "num_pins" : 1,
  "owner_id" : 90000001,
  "planet_id" : 40023691,
  "planet_type" : "plasma",
  "solar_system_id" : 30000379,
  "upgrade_level" : 0
}, {
  "last_update" : "2016-11-28T16:41:54Z",
  "num_pins" : 1,
  "owner_id" : 90000001,
  "planet_id" : 40023697,
  "planet_type" : "barren",
  "solar_system_id" : 30000379,
  "upgrade_level" : 0
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCharactersCharacterIdPlanetsPlanetId(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "links" : [ {
    "destination_pin_id" : 1000000017022,
    "link_level" : 0,
    "source_pin_id" : 1000000017021
  } ],
  "pins" : [ {
    "is_running" : true,
    "latitude" : 1.55087844973,
    "longitude" : 0.717145933308,
    "pin_id" : 1000000017021,
    "type_id" : 2254
  }, {
    "is_running" : true,
    "latitude" : 1.53360639935,
    "longitude" : 0.709775584394,
    "pin_id" : 1000000017022,
    "type_id" : 2256
  } ],
  "routes" : [ {
    "content_type_id" : 2393,
    "destination_pin_id" : 1000000017030,
    "quantity" : 20,
    "route_id" : 4,
    "source_pin_id" : 1000000017029
  } ]
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetUniverseSchematicsSchematicId(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "cycle_time" : 1800,
  "schematic_name" : "Bacteria"
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


