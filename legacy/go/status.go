package esilegacy

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

var _ time.Time
var _ = mux.NewRouter

func GetStatus(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		datasource string
		userAgent string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "players" : 12345,
  "server_version" : "1132976",
  "start_time" : "2017-01-02T12:34:56Z"
}`
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


