package esiV3

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var _ time.Time
var _ = mux.NewRouter

func GetCorporationsCorporationId(w http.ResponseWriter, r *http.Request) {

	var (
		localV        interface{}
		err           error
		corporationId int32
		datasource    string
		userAgent     string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "alliance_id" : 434243723,
  "ceo_id" : 180548812,
  "corporation_description" : "This is a corporation description, it's basically just a string",
  "corporation_name" : "C C P",
  "creation_date" : "2004-11-28T16:42:51Z",
  "creator_id" : 180548812,
  "member_count" : 656,
  "tax_rate" : 0.256,
  "ticker" : "-CCP-",
  "url" : "http://www.eveonline.com"
}`
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
