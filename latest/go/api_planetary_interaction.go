package esilatest

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var _ time.Time
var _ = mux.NewRouter

func GetCharactersCharacterIdPlanets(w http.ResponseWriter, r *http.Request) {

	var (
		localV      interface{}
		err         error
		characterId int32
		datasource  string
		token       string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "last_update" : "2016-11-28T16:42:51Z",
  "num_pins" : 1,
  "owner_id" : 90000001,
  "planet_id" : 40023691,
  "planet_type" : "plasma",
  "solar_system_id" : 30000379,
  "upgrade_level" : 0
}, {
  "last_update" : "2016-11-28T16:41:54Z",
  "num_pins" : 1,
  "owner_id" : 90000001,
  "planet_id" : 40023697,
  "planet_type" : "barren",
  "solar_system_id" : 30000379,
  "upgrade_level" : 0
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(characterId, vars["character_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	characterId = localV.(int32)
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

func GetCharactersCharacterIdPlanetsPlanetId(w http.ResponseWriter, r *http.Request) {

	var (
		localV      interface{}
		err         error
		characterId int32
		planetId    int32
		datasource  string
		token       string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "links" : [ {
    "destination_pin_id" : 1000000017022,
    "link_level" : 0,
    "source_pin_id" : 1000000017021
  } ],
  "pins" : [ {
    "latitude" : 1.55087844973,
    "longitude" : 0.717145933308,
    "pin_id" : 1000000017021,
    "type_id" : 2254
  }, {
    "latitude" : 1.53360639935,
    "longitude" : 0.709775584394,
    "pin_id" : 1000000017022,
    "type_id" : 2256
  } ],
  "routes" : [ {
    "content_type_id" : 2393,
    "destination_pin_id" : 1000000017030,
    "quantity" : 20,
    "route_id" : 4,
    "source_pin_id" : 1000000017029
  } ]
}`
	vars := mux.Vars(r)
	localV, err = processParameters(characterId, vars["character_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	characterId = localV.(int32)
	localV, err = processParameters(planetId, vars["planet_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	planetId = localV.(int32)
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

func GetCorporationsCorporationIdCustomsOffices(w http.ResponseWriter, r *http.Request) {

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
  "alliance_tax_rate" : 0.1,
  "allow_access_with_standings" : true,
  "allow_alliance_access" : false,
  "corporation_tax_rate" : 0.02,
  "excellent_standing_tax_rate" : 0.05,
  "good_standing_tax_rate" : 0.2,
  "neutral_standing_tax_rate" : 0.5,
  "office_id" : 1000000014530,
  "reinforce_exit_end" : 1,
  "reinforce_exit_start" : 23,
  "standing_level" : "neutral",
  "system_id" : 30003657
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

func GetUniverseSchematicsSchematicId(w http.ResponseWriter, r *http.Request) {

	var (
		localV      interface{}
		err         error
		schematicId int32
		datasource  string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "cycle_time" : 1800,
  "schematic_name" : "Bacteria"
}`
	vars := mux.Vars(r)
	localV, err = processParameters(schematicId, vars["schematic_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	schematicId = localV.(int32)
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
