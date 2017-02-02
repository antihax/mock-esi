package esiLegacy

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

var _ time.Time
var _ = mux.NewRouter

func GetInsurancePrices(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		language string
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "levels" : [ {
    "cost" : 10.0,
    "name" : "Basic",
    "payout" : 20.0
  } ],
  "type_id" : 1
} ]`
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("language") != "" {
		localV, err = processParameters(language, r.Form.Get("language"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		language = localV.(string)
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
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


