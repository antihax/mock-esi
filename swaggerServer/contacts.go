package swaggerServer

import (
	"net/http"

)

func DeleteCharactersCharacterIdContacts(w http.ResponseWriter, r *http.Request) {

	j := (``)
		w.Header().Set("Content-Type", "")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCharactersCharacterIdContacts(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "contact_id" : 123,
  "contact_type" : "character",
  "is_blocked" : false,
  "is_watched" : true,
  "standing" : 10.0
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCharactersCharacterIdContactsLabels(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "label_id" : 123,
  "label_name" : "Friends"
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func PostCharactersCharacterIdContacts(w http.ResponseWriter, r *http.Request) {

	j := (`[ 123, 456 ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func PutCharactersCharacterIdContacts(w http.ResponseWriter, r *http.Request) {

	j := (``)
		w.Header().Set("Content-Type", "")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


