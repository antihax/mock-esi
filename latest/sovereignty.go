package esiLatest

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

var _ time.Time
var _ = mux.NewRouter

func GetSovereigntyCampaigns(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "attackers_score" : 0.4,
  "campaign_id" : 32833,
  "constellation_id" : 20000125,
  "defender_id" : 1695357456,
  "defender_score" : 0.6,
  "event_type" : "station_defense",
  "solar_system_id" : 30000856,
  "start_time" : "2016-10-29T14:34:40Z",
  "structure_id" : 61001096
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

func GetSovereigntyStructures(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "alliance_id" : 498125261,
  "solar_system_id" : 30000570,
  "structure_id" : 1018253388776,
  "structure_type_id" : 32226,
  "vulnerability_occupancy_level" : 2,
  "vulnerable_end_time" : "2016-10-29T05:30:00Z",
  "vulnerable_start_time" : "2016-10-28T20:30:00Z"
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


