package esiv3

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var _ time.Time
var _ = mux.NewRouter

func GetCharactersCharacterIdPlanetsPlanetId(w http.ResponseWriter, r *http.Request) {

	var (
		localV      interface{}
		err         error
		characterId int32
		planetId    int32
		datasource  string
		token       string
		userAgent   string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "links" : [ {
    "source_pin_id" : 1000000017021,
    "destination_pin_id" : 1000000017022,
    "link_level" : 0
  } ],
  "pins" : [ {
    "longitude" : 0.717145933308,
    "is_running" : true,
    "type_id" : 2254,
    "pin_id" : 1000000017021,
    "latitude" : 1.55087844973
  }, {
    "longitude" : 0.709775584394,
    "is_running" : true,
    "type_id" : 2256,
    "pin_id" : 1000000017022,
    "latitude" : 1.53360639935
  } ],
  "routes" : [ {
    "source_pin_id" : 1000000017029,
    "quantity" : 20,
    "content_type_id" : 2393,
    "route_id" : 4,
    "destination_pin_id" : 1000000017030
  } ]
}`
	vars := mux.Vars(r)
	localV, err = processParameters(characterId, vars["character_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	characterId = localV.(int32)
	localV, err = processParameters(planetId, vars["planet_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	planetId = localV.(int32)
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
	if r.Form.Get("token") != "" {
		localV, err = processParameters(token, r.Form.Get("token"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		token = localV.(string)
	}
	if r.Form.Get("userAgent") != "" {
		localV, err = processParameters(userAgent, r.Form.Get("user_agent"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		userAgent = localV.(string)
	}

	if r.Form.Get("page") != "" {
		var (
			localPage    int32
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
