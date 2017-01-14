package swaggerServer

import (
	"net/http"

)

func DeleteFleetsFleetIdMembersMemberId(w http.ResponseWriter, r *http.Request) {

	j := (``)
		w.Header().Set("Content-Type", "")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func DeleteFleetsFleetIdSquadsSquadId(w http.ResponseWriter, r *http.Request) {

	j := (``)
		w.Header().Set("Content-Type", "")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func DeleteFleetsFleetIdWingsWingId(w http.ResponseWriter, r *http.Request) {

	j := (``)
		w.Header().Set("Content-Type", "")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetFleetsFleetId(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "is_free_move" : false,
  "is_registered" : false,
  "is_voice_enabled" : false,
  "motd" : "This is an <b>awesome</b> fleet!"
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetFleetsFleetIdMembers(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "character_id" : 93265215,
  "join_time" : "2016-04-29T12:34:56Z",
  "role" : "squad_commander",
  "role_name" : "Squad Commander (Boss)",
  "ship_type_id" : 33328,
  "solar_system_id" : 30003729,
  "squad_id" : 3129411261968,
  "station_id" : 61000180,
  "takes_fleet_warp" : true,
  "wing_id" : 2073711261968
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetFleetsFleetIdWings(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "id" : 2073711261968,
  "name" : "Wing 1",
  "squads" : [ {
    "id" : 3129411261968,
    "name" : "Squad 1"
  } ]
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func PostFleetsFleetIdMembers(w http.ResponseWriter, r *http.Request) {

	j := (``)
		w.Header().Set("Content-Type", "")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func PostFleetsFleetIdWings(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "wing_id" : 123
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func PostFleetsFleetIdWingsWingIdSquads(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "squad_id" : 123
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func PutFleetsFleetId(w http.ResponseWriter, r *http.Request) {

	j := (``)
		w.Header().Set("Content-Type", "")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func PutFleetsFleetIdMembersMemberId(w http.ResponseWriter, r *http.Request) {

	j := (``)
		w.Header().Set("Content-Type", "")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func PutFleetsFleetIdSquadsSquadId(w http.ResponseWriter, r *http.Request) {

	j := (``)
		w.Header().Set("Content-Type", "")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func PutFleetsFleetIdWingsWingId(w http.ResponseWriter, r *http.Request) {

	j := (``)
		w.Header().Set("Content-Type", "")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


