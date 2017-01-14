package swaggerServer

import (
	"net/http"
	"github.com/gorilla/mux"
)

var _ = mux.NewRouter

func GetCharactersCharacterIdLocation(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		characterId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "solar_system_id" : 30002505,
  "structure_id" : 1000000016989
}`
	vars := mux.Vars(r)
	localV, err = processParameters(characterId, vars["characterId"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	characterId = localV.(int32)
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func GetCharactersCharacterIdShip(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		characterId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "ship_item_id" : 1000000016991,
  "ship_name" : "SPACESHIPS!!!",
  "ship_type_id" : 1233
}`
	vars := mux.Vars(r)
	localV, err = processParameters(characterId, vars["characterId"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	characterId = localV.(int32)
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}


