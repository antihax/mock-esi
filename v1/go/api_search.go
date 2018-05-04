package esiv1

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var _ time.Time
var _ = mux.NewRouter

func GetSearch(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		categories []string
		search     string
		datasource string
		language   string
		strict     bool
		userAgent  string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "solarsystem" : [ 30002510 ],
  "station" : [ 60004588, 60004594, 60005725, 60009106, 60012721, 60012724, 60012727 ]
}`
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	localV, err = processParameters(categories, r.Form.Get("categories"))
	if err != nil {
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
	if r.Form.Get("language") != "" {
		localV, err = processParameters(language, r.Form.Get("language"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		language = localV.(string)
	}
	localV, err = processParameters(search, r.Form.Get("search"))
	if err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("strict") != "" {
		localV, err = processParameters(strict, r.Form.Get("strict"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		strict = localV.(bool)
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
