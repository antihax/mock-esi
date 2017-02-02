package esiLatest

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

var _ time.Time
var _ = mux.NewRouter

func PostUiAutopilotWaypoint(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		destinationId int64
		clearOtherWaypoints bool
		addToBeginning bool
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := ``
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	localV, err = processParameters(destinationId, r.Form.Get("destination_id"))
	if err != nil {
		errorOut(w, r, err)
		return
	}
	localV, err = processParameters(clearOtherWaypoints, r.Form.Get("clear_other_waypoints"))
	if err != nil {
		errorOut(w, r, err)
		return
	}
	localV, err = processParameters(addToBeginning, r.Form.Get("add_to_beginning"))
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
			w.Header().Set("Content-Type", "")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	} 

	w.Header().Set("Content-Type", "")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func PostUiOpenwindowContract(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		contractId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := ``
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	localV, err = processParameters(contractId, r.Form.Get("contract_id"))
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
			w.Header().Set("Content-Type", "")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	} 

	w.Header().Set("Content-Type", "")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func PostUiOpenwindowInformation(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		targetId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := ``
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	localV, err = processParameters(targetId, r.Form.Get("target_id"))
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
			w.Header().Set("Content-Type", "")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	} 

	w.Header().Set("Content-Type", "")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func PostUiOpenwindowMarketdetails(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		typeId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := ``
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	localV, err = processParameters(typeId, r.Form.Get("type_id"))
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
			w.Header().Set("Content-Type", "")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	} 

	w.Header().Set("Content-Type", "")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func PostUiOpenwindowNewmail(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := ``
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
			w.Header().Set("Content-Type", "")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	} 

	w.Header().Set("Content-Type", "")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}


