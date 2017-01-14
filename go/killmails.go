package swaggerServer

import (
	"net/http"

)

func GetCharactersCharacterIdKillmailsRecent(w http.ResponseWriter, r *http.Request) {

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

func GetKillmailsKillmailIdKillmailHash(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "attackers" : [ {
    "character_id" : 95810944,
    "corporation_id" : 1000179,
    "damage_done" : 5745,
    "faction_id" : 500003,
    "final_blow" : true,
    "security_status" : -0.3,
    "ship_type_id" : 17841,
    "weapon_type_id" : 3074
  } ],
  "killmail_id" : 56733821,
  "killmail_time" : "2016-10-22T17:13:36Z",
  "solar_system_id" : 30002976,
  "victim" : {
    "alliance_id" : 621338554,
    "character_id" : 92796241,
    "corporation_id" : 841363671,
    "damage_taken" : 5745,
    "items" : [ {
      "flag" : 20,
      "item_type_id" : 5973,
      "quantity_dropped" : 1,
      "singleton" : 0
    } ],
    "position" : {
      "x" : 4.521866005694748E11,
      "y" : 1.4670496149090222E11,
      "z" : 1.0951459653254477E11
    },
    "ship_type_id" : 17812
  }
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


