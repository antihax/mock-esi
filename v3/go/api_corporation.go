package esiv3

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var _ time.Time
var _ = mux.NewRouter

func GetCorporationsCorporationIdAlliancehistory(w http.ResponseWriter, r *http.Request) {

	var (
		localV        interface{}
		err           error
		corporationId int32
		datasource    string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "alliance_id" : 99000006,
  "is_deleted" : true,
  "record_id" : 23,
  "start_date" : "2016-10-25T14:46:00Z"
}, {
  "record_id" : 1,
  "start_date" : "2015-07-06T20:56:00Z"
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(corporationId, vars["corporation_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	corporationId = localV.(int32)
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

func GetCorporationsCorporationIdBlueprints(w http.ResponseWriter, r *http.Request) {

	var (
		localV        interface{}
		err           error
		corporationId int32
		datasource    string
		page          int32
		token         string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "item_id" : 1000000010495,
  "location_flag" : "CorpSAG1",
  "location_id" : 60014719,
  "material_efficiency" : 0,
  "quantity" : 1,
  "runs" : -1,
  "time_efficiency" : 0,
  "type_id" : 691
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(corporationId, vars["corporation_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	corporationId = localV.(int32)
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
		localV, err = processParameters(page, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		page = localV.(int32)
	}
	if r.Form.Get("token") != "" {
		localV, err = processParameters(token, r.Form.Get("token"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		token = localV.(string)
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

func GetCorporationsCorporationIdContainersLogs(w http.ResponseWriter, r *http.Request) {

	var (
		localV        interface{}
		err           error
		corporationId int32
		datasource    string
		page          int32
		token         string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "action" : "set_password",
  "character_id" : 2112625428,
  "container_id" : 1000000012279,
  "container_type_id" : 17365,
  "location_flag" : "CorpSAG1",
  "location_id" : 1000000012278,
  "logged_at" : "2017-10-10T14:00:00Z",
  "password_type" : "general"
}, {
  "action" : "lock",
  "character_id" : 2112625428,
  "container_id" : 1000000012279,
  "container_type_id" : 17365,
  "location_flag" : "CorpSAG1",
  "location_id" : 1000000012278,
  "logged_at" : "2017-10-11T12:04:33Z",
  "quantity" : 30,
  "type_id" : 1230
}, {
  "action" : "configure",
  "character_id" : 2112625428,
  "container_id" : 1000000012279,
  "container_type_id" : 17365,
  "location_flag" : "CorpSAG2",
  "location_id" : 1000000012278,
  "logged_at" : "2017-10-11T12:06:29Z",
  "new_config_bitmask" : 31,
  "old_config_bitmask" : 23
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(corporationId, vars["corporation_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	corporationId = localV.(int32)
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
		localV, err = processParameters(page, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		page = localV.(int32)
	}
	if r.Form.Get("token") != "" {
		localV, err = processParameters(token, r.Form.Get("token"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		token = localV.(string)
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

func GetCorporationsCorporationIdMembers(w http.ResponseWriter, r *http.Request) {

	var (
		localV        interface{}
		err           error
		corporationId int32
		datasource    string
		token         string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ 90000001, 90000002 ]`
	vars := mux.Vars(r)
	localV, err = processParameters(corporationId, vars["corporation_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	corporationId = localV.(int32)
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
			w.Header().Set("warning", "299 - This route is deprecated.")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	}

	w.Header().Set("warning", "299 - This route is deprecated.")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}
