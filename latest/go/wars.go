package esilatest

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

var _ time.Time
var _ = mux.NewRouter

func GetWars(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		datasource string
		maxWarId int32
		userAgent string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ 3, 2, 1 ]`
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
	if r.Form.Get("maxWarId") != "" {
		localV, err = processParameters(maxWarId, r.Form.Get("max_war_id"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		maxWarId = localV.(int32)
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

func GetWarsWarId(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		warId int32
		datasource string
		userAgent string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "aggressor" : {
    "corporation_id" : 986665792,
    "isk_destroyed" : 0,
    "ships_killed" : 0
  },
  "declared" : "2004-05-22T05:20:00Z",
  "defender" : {
    "corporation_id" : 1001562011,
    "isk_destroyed" : 0,
    "ships_killed" : 0
  },
  "id" : 1941,
  "mutual" : false,
  "open_for_allies" : false
}`
	vars := mux.Vars(r)
	localV, err = processParameters(warId, vars["war_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	warId = localV.(int32)
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

func GetWarsWarIdKillmails(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		warId int32
		datasource string
		page int32
		userAgent string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "killmail_hash" : "8eef5e8fb6b88fe3407c489df33822b2e3b57a5e",
  "killmail_id" : 2
}, {
  "killmail_hash" : "b41ccb498ece33d64019f64c0db392aa3aa701fb",
  "killmail_id" : 1
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(warId, vars["war_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	warId = localV.(int32)
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


