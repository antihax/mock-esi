package swaggerServer

import (
	"net/http"
	"github.com/gorilla/mux"
)

var _ = mux.NewRouter

func GetUniverseStationsStationId(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		stationId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "solar_system_id" : 30000142,
  "station_name" : "Jita IV Moon IV - Caldari Navy Assembly Plant"
}`
	vars := mux.Vars(r)
	localV, err = processParameters(stationId, vars["stationId"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	stationId = localV.(int32)
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

func GetUniverseStructures(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ 1000000017013, 1000000025062 ]`
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

func GetUniverseStructuresStructureId(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		structureId int64
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "name" : "V-3YG7 VI - The Capital",
  "solar_system_id" : 30000142
}`
	vars := mux.Vars(r)
	localV, err = processParameters(structureId, vars["structureId"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	structureId = localV.(int64)
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

func GetUniverseSystemsSystemId(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		systemId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "solar_system_name" : "Jita"
}`
	vars := mux.Vars(r)
	localV, err = processParameters(systemId, vars["systemId"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	systemId = localV.(int32)
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

func GetUniverseTypesTypeId(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		typeId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "category_id" : 6,
  "graphic_id" : 46,
  "group_id" : 25,
  "type_description" : "The Rifter is a very powerful combat frigate and can easily tackle the best frigates out there. It has gone through many radical design phases since its inauguration during the Minmatar Rebellion. The Rifter has a wide variety of offensive capabilities, making it an unpredictable and deadly adversary.",
  "type_name" : "Rifter"
}`
	vars := mux.Vars(r)
	localV, err = processParameters(typeId, vars["typeId"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	typeId = localV.(int32)
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

func PostUniverseNames(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "category" : "character",
  "id" : 95465499,
  "name" : "CCP Bartender"
}, {
  "category" : "solar_system",
  "id" : 30000142,
  "name" : "Jita"
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}


