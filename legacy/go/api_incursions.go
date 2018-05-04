package esilegacy

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

var _ time.Time
var _ = mux.NewRouter

func GetIncursions(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		datasource string
		userAgent string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "constellation_id" : 20000607,
  "faction_id" : 500019,
  "has_boss" : true,
  "infested_solar_systems" : [ 30004148, 30004149, 30004150, 30004151, 30004152, 30004153, 30004154 ],
  "influence" : 0.9,
  "staging_solar_system_id" : 30004154,
  "state" : "mobilizing",
  "type" : "Incursion"
} ]`
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


