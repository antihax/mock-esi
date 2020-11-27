package esiv1

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var _ time.Time
var _ = mux.NewRouter

func GetCorporationsCorporationIdWallets(w http.ResponseWriter, r *http.Request) {

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

	j := `[ {
  "balance" : 123.45,
  "division" : 1
}, {
  "balance" : 123.45,
  "division" : 2
}, {
  "balance" : 123.45,
  "division" : 3
}, {
  "balance" : 123.45,
  "division" : 4
}, {
  "balance" : 123.45,
  "division" : 5
}, {
  "balance" : 123.45,
  "division" : 6
}, {
  "balance" : 123.45,
  "division" : 7
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
