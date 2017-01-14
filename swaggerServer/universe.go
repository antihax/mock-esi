package swaggerServer

import (
	"net/http"

)

func GetUniverseStationsStationId(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "solar_system_id" : 30000142,
  "station_name" : "Jita IV Moon IV - Caldari Navy Assembly Plant"
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetUniverseStructures(w http.ResponseWriter, r *http.Request) {

	j := (`[ 1000000017013, 1000000025062 ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetUniverseStructuresStructureId(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "name" : "V-3YG7 VI - The Capital",
  "solar_system_id" : 30000142
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetUniverseSystemsSystemId(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "solar_system_name" : "Jita"
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetUniverseTypesTypeId(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "category_id" : 6,
  "graphic_id" : 46,
  "group_id" : 25,
  "type_description" : "The Rifter is a very powerful combat frigate and can easily tackle the best frigates out there. It has gone through many radical design phases since its inauguration during the Minmatar Rebellion. The Rifter has a wide variety of offensive capabilities, making it an unpredictable and deadly adversary.",
  "type_name" : "Rifter"
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func PostUniverseNames(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "category" : "character",
  "id" : 95465499,
  "name" : "CCP Bartender"
}, {
  "category" : "solar_system",
  "id" : 30000142,
  "name" : "Jita"
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


