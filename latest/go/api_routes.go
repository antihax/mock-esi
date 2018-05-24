package esilatest

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var _ time.Time
var _ = mux.NewRouter

func GetRouteOriginDestination(w http.ResponseWriter, r *http.Request) {

	var (
		localV      interface{}
		err         error
		destination int32
		origin      int32
		avoid       []int32
		connections [][]int32
		datasource  string
		flag        string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ 30002771, 30002770, 30002769, 30002772 ]`
	vars := mux.Vars(r)
	localV, err = processParameters(destination, vars["destination"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	destination = localV.(int32)
	localV, err = processParameters(origin, vars["origin"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	origin = localV.(int32)
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("avoid") != "" {
		localV, err = processParameters(avoid, r.Form.Get("avoid"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		avoid = localV.([]int32)
	}
	if r.Form.Get("connections") != "" {
		localV, err = processParameters(connections, r.Form.Get("connections"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		connections = localV.([][]int32)
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}
	if r.Form.Get("flag") != "" {
		localV, err = processParameters(flag, r.Form.Get("flag"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		flag = localV.(string)
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
