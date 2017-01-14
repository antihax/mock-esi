package swaggerServer

import (
	"net/http"

)

func GetCharactersCharacterId(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "ancestry_id" : 19,
  "birthday" : "2015-03-24T11:37:00Z",
  "bloodline_id" : 3,
  "corporation_id" : 109299958,
  "description" : "",
  "gender" : "male",
  "name" : "CCP Bartender",
  "race_id" : 2
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCharactersCharacterIdCorporationhistory(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "corporation_id" : 90000001,
  "is_deleted" : false,
  "record_id" : 500,
  "start_date" : "2016-06-26T20:00:00Z"
}, {
  "corporation_id" : 90000002,
  "is_deleted" : false,
  "record_id" : 501,
  "start_date" : "2016-07-26T20:00:00Z"
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCharactersCharacterIdPortrait(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "px128x128" : "https://imageserver.eveonline.com/Character/95465499_128.jpg",
  "px256x256" : "https://imageserver.eveonline.com/Character/95465499_256.jpg",
  "px512x512" : "https://imageserver.eveonline.com/Character/95465499_512.jpg",
  "px64x64" : "https://imageserver.eveonline.com/Character/95465499_64.jpg"
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCharactersNames(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "character_id" : 95465499,
  "character_name" : "CCP Bartender"
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func PostCharactersCharacterIdCspa(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "cost" : 295000
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


