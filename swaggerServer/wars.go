package swaggerServer

import (
	"net/http"

)

func GetWars(w http.ResponseWriter, r *http.Request) {

	j := (`[ 3, 2, 1 ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetWarsWarId(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "aggressor" : {
    "corporation_id" : 986665792,
    "isk_destroyed" : 0,
    "ships_killed" : 0
  },
  "declared" : "2004-05-22T05:20:00Z",
  "defender" : {
    "corporation_id" : 1001562011,
    "isk_destroyed" : 0,
    "ships_killed" : 0
  },
  "id" : 1941,
  "mutual" : false,
  "open_for_allies" : false
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetWarsWarIdKillmails(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "killmail_hash" : "8eef5e8fb6b88fe3407c489df33822b2e3b57a5e",
  "killmail_id" : 2
}, {
  "killmail_hash" : "b41ccb498ece33d64019f64c0db392aa3aa701fb",
  "killmail_id" : 1
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


