package swaggerServer

import (
	"net/http"

)

func GetSovereigntyCampaigns(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "attackers_score" : 0.4,
  "campaign_id" : 32833,
  "constellation_id" : 20000125,
  "defender_id" : 1695357456,
  "defender_score" : 0.6,
  "event_type" : "station_defense",
  "solar_system_id" : 30000856,
  "start_time" : "2016-10-29T14:34:40Z",
  "structure_id" : 61001096
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetSovereigntyStructures(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "alliance_id" : 498125261,
  "solar_system_id" : 30000570,
  "structure_id" : 1018253388776,
  "structure_type_id" : 32226,
  "vulnerability_occupancy_level" : 2,
  "vulnerable_end_time" : "2016-10-29T05:30:00Z",
  "vulnerable_start_time" : "2016-10-28T20:30:00Z"
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


