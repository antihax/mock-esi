package swaggerServer

import (
	"net/http"

)

func GetAlliances(w http.ResponseWriter, r *http.Request) {

	j := (`[ 99000001, 99000002 ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetAlliancesAllianceId(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "alliance_name" : "C C P Alliance",
  "date_founded" : "2016-06-26T21:00:00Z",
  "executor_corp" : 98356193,
  "ticker" : "<C C P>"
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetAlliancesAllianceIdCorporations(w http.ResponseWriter, r *http.Request) {

	j := (`[ 98000001 ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetAlliancesAllianceIdIcons(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "px128x128" : "https://imageserver.eveonline.com/Alliance/503818424_128.png",
  "px64x64" : "https://imageserver.eveonline.com/Alliance/503818424_64.png"
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetAlliancesNames(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "alliance_id" : 1000171,
  "alliance_name" : "Republic University"
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


