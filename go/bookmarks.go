package swaggerServer

import (
	"net/http"

)

func GetCharactersCharacterIdBookmarks(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "bookmark_id" : 32,
  "create_date" : "2016-08-09T11:57:47Z",
  "creator_id" : 90000001,
  "folder_id" : 5,
  "memo" : "aoeu ( Citadel )",
  "note" : "",
  "owner_id" : 90000001,
  "target" : {
    "item" : {
      "item_id" : 1000000012668,
      "type_id" : 35832
    },
    "location_id" : 30000005
  }
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCharactersCharacterIdBookmarksFolders(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "folder_id" : 5,
  "name" : "Icecream",
  "owner_id" : 90000001
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


