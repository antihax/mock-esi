package swS

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

var _ = mux.NewRouter

func GetUniverseBloodlines(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		language string
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "bloodline_id" : 1,
  "charisma" : 6,
  "corporation_id" : 1000006,
  "description" : "The Deteis are regarded as ...",
  "intelligence" : 7,
  "memory" : 7,
  "name" : "Deteis",
  "perception" : 5,
  "race_id" : 1,
  "ship_type_id" : 601,
  "willpower" : 5
} ]`
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("language") != "" {
		localV, err = processParameters(language, r.Form.Get("language"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		language = localV.(string)
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}

fmt.Printf("%s\n", r.Form.Get("page"))
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

func GetUniverseCategories(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ 1, 2, 3 ]`
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

fmt.Printf("%s\n", r.Form.Get("page"))
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

func GetUniverseCategoriesCategoryId(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		categoryId int32
		language string
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "category_id" : 6,
  "groups" : [ 25, 26, 27 ],
  "name" : "Ship",
  "published" : true
}`
	vars := mux.Vars(r)
	localV, err = processParameters(categoryId, vars["category_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	categoryId = localV.(int32)
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("language") != "" {
		localV, err = processParameters(language, r.Form.Get("language"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		language = localV.(string)
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}

fmt.Printf("%s\n", r.Form.Get("page"))
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

func GetUniverseFactions(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		language string
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "corporation_id" : 456,
  "description" : "blah blah",
  "faction_id" : 1,
  "is_unique" : true,
  "name" : "Faction",
  "size_factor" : 1.0,
  "solar_system_id" : 123,
  "station_count" : 1000,
  "station_system_count" : 100
} ]`
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("language") != "" {
		localV, err = processParameters(language, r.Form.Get("language"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		language = localV.(string)
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}

fmt.Printf("%s\n", r.Form.Get("page"))
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

func GetUniverseGroups(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		page int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ 1, 2, 3 ]`
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("page") != "" {
		localV, err = processParameters(page, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		page = localV.(int32)
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}

fmt.Printf("%s\n", r.Form.Get("page"))
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

func GetUniverseGroupsGroupId(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		groupId int32
		language string
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "category_id" : 6,
  "group_id" : 25,
  "name" : "Frigate",
  "published" : true,
  "types" : [ 587, 586, 585 ]
}`
	vars := mux.Vars(r)
	localV, err = processParameters(groupId, vars["group_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	groupId = localV.(int32)
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("language") != "" {
		localV, err = processParameters(language, r.Form.Get("language"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		language = localV.(string)
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}

fmt.Printf("%s\n", r.Form.Get("page"))
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

func GetUniverseRaces(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		language string
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "alliance_id" : 500001,
  "description" : "Founded on the tenets of patriotism and hard work...",
  "name" : "Caldari",
  "race_id" : 1
} ]`
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("language") != "" {
		localV, err = processParameters(language, r.Form.Get("language"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		language = localV.(string)
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}

fmt.Printf("%s\n", r.Form.Get("page"))
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
	localV, err = processParameters(stationId, vars["station_id"])
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

fmt.Printf("%s\n", r.Form.Get("page"))
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

fmt.Printf("%s\n", r.Form.Get("page"))
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
	localV, err = processParameters(structureId, vars["structure_id"])
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

fmt.Printf("%s\n", r.Form.Get("page"))
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
	localV, err = processParameters(systemId, vars["system_id"])
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

fmt.Printf("%s\n", r.Form.Get("page"))
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

func GetUniverseTypes(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		page int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ 1, 2, 3 ]`
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("page") != "" {
		localV, err = processParameters(page, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		page = localV.(int32)
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}

fmt.Printf("%s\n", r.Form.Get("page"))
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

func GetUniverseTypesTypeId(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		typeId int32
		language string
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "description" : "The Rifter is a...",
  "group_id" : 25,
  "name" : "Rifter",
  "published" : true,
  "type_id" : 587
}`
	vars := mux.Vars(r)
	localV, err = processParameters(typeId, vars["type_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	typeId = localV.(int32)
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("language") != "" {
		localV, err = processParameters(language, r.Form.Get("language"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		language = localV.(string)
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}

fmt.Printf("%s\n", r.Form.Get("page"))
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

fmt.Printf("%s\n", r.Form.Get("page"))
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


