package swaggerServer

import (
	"net/http"
	"github.com/gorilla/mux"
)

var _ = mux.NewRouter

func GetAlliances(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ 99000001, 99000002 ]`
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func GetAlliancesAllianceId(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		allianceId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "alliance_name" : "C C P Alliance",
  "date_founded" : "2016-06-26T21:00:00Z",
  "executor_corp" : 98356193,
  "ticker" : "<C C P>"
}`
	vars := mux.Vars(r)
	localV, err = processParameters(allianceId, vars["allianceId"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	allianceId = localV.(int32)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func GetAlliancesAllianceIdCorporations(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		allianceId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ 98000001 ]`
	vars := mux.Vars(r)
	localV, err = processParameters(allianceId, vars["allianceId"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	allianceId = localV.(int32)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func GetAlliancesAllianceIdIcons(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		allianceId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "px128x128" : "https://imageserver.eveonline.com/Alliance/503818424_128.png",
  "px64x64" : "https://imageserver.eveonline.com/Alliance/503818424_64.png"
}`
	vars := mux.Vars(r)
	localV, err = processParameters(allianceId, vars["allianceId"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	allianceId = localV.(int32)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func GetAlliancesNames(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		allianceIds []int64
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "alliance_id" : 1000171,
  "alliance_name" : "Republic University"
} ]`
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	localV, err = processParameters(allianceIds, r.Form.Get("allianceIds"))
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}


