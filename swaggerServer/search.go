package swaggerServer

import (
	"net/http"

)

func GetCharactersCharacterIdSearch(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "solarsystem" : [ 30002510 ],
  "station" : [ 60004588, 60004594, 60005725, 60009106, 60012721, 60012724, 60012727 ]
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetSearch(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "solarsystem" : [ 30002510 ],
  "station" : [ 60004588, 60004594, 60005725, 60009106, 60012721, 60012724, 60012727 ]
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


