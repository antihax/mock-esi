package swaggerServer

import (
	"net/http"

)

func GetCorporationsCorporationId(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "alliance_id" : 434243723,
  "ceo_id" : 180548812,
  "corporation_name" : "C C P",
  "member_count" : 656,
  "ticker" : "-CCP-"
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCorporationsCorporationIdAlliancehistory(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "alliance" : {
    "alliance_id" : 99000006,
    "is_deleted" : false
  },
  "record_id" : 23,
  "start_date" : "2016-10-25T14:46:00Z"
}, {
  "record_id" : 1,
  "start_date" : "2015-07-06T20:56:00Z"
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCorporationsCorporationIdIcons(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "px128x128" : "https://imageserver.eveonline.com/Corporation/1000010_128.png",
  "px256x256" : "https://imageserver.eveonline.com/Corporation/1000010_256.png",
  "px64x64" : "https://imageserver.eveonline.com/Corporation/1000010_64.png"
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCorporationsCorporationIdMembers(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "character_id" : 90000001
}, {
  "character_id" : 90000002
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCorporationsCorporationIdRoles(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "character_id" : 1000171,
  "roles" : [ "Director", "Station_Manager" ]
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCorporationsNames(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "corporation_id" : 1000171,
  "corporation_name" : "Republic University"
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


