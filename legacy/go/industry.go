package esiLegacy

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var _ time.Time
var _ = mux.NewRouter

func GetCharactersCharacterIdIndustryJobs(w http.ResponseWriter, r *http.Request) {

	var (
		localV           interface{}
		err              error
		characterId      int32
		datasource       string
		includeCompleted bool
		token            string
		userAgent        string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "activity_id" : 1,
  "blueprint_id" : 1015116533326,
  "blueprint_location_id" : 60006382,
  "blueprint_type_id" : 2047,
  "cost" : 118.0,
  "duration" : 548,
  "end_date" : "2014-07-19T15:56:14Z",
  "facility_id" : 60006382,
  "installer_id" : 498338451,
  "job_id" : 229136101,
  "licensed_runs" : 200,
  "output_location_id" : 60006382,
  "runs" : 1,
  "start_date" : "2014-07-19T15:47:06Z",
  "station_id" : 60006382,
  "status" : "active"
} ]`
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
	if r.Form.Get("includeCompleted") != "" {
		localV, err = processParameters(includeCompleted, r.Form.Get("include_completed"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		includeCompleted = localV.(bool)
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

func GetIndustryFacilities(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		datasource string
		userAgent  string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "facility_id" : 60012544,
  "owner_id" : 1000126,
  "region_id" : 10000001,
  "solar_system_id" : 30000032,
  "tax" : 0.1,
  "type_id" : 2502
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

func GetIndustrySystems(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		datasource string
		userAgent  string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "cost_indices" : [ {
    "activity" : "invention",
    "cost_index" : 0.00480411064973412
  } ],
  "solar_system_id" : 30011392
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
