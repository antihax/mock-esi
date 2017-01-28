package swS

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

var _ = mux.NewRouter

func GetCharactersCharacterIdClones(w http.ResponseWriter, r *http.Request) {

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
  "home_location" : {
    "location_id" : 1021348135816,
    "location_type" : "structure"
  },
  "jump_clones" : [ {
    "implants" : [ 22118 ],
    "location_id" : 60003463,
    "location_type" : "station"
  }, {
    "implants" : [ ],
    "location_id" : 1021348135816,
    "location_type" : "structure"
  } ]
}`
	vars := mux.Vars(r)
	localV, err = processParameters(characterId, vars["character_id"])
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

fmt.Printf("%s\n", r.Form.Get("page"))
	if r.Form.Get("page") != "" {
		var (
			localPage int32 
			localIntPage interface{}
		)
		localIntPage, err := processParameters(localPage, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		localPage = localIntPage.(int32)
		if localPage > 1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	} 

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}


