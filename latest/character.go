package esiLatest

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

var _ time.Time
var _ = mux.NewRouter

func GetCharactersCharacterId(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		characterId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "ancestry_id" : 19,
  "birthday" : "2015-03-24T11:37:00Z",
  "bloodline_id" : 3,
  "corporation_id" : 109299958,
  "description" : "",
  "gender" : "male",
  "name" : "CCP Bartender",
  "race_id" : 2
}`
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

func GetCharactersCharacterIdCorporationhistory(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		characterId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "corporation_id" : 90000001,
  "is_deleted" : false,
  "record_id" : 500,
  "start_date" : "2016-06-26T20:00:00Z"
}, {
  "corporation_id" : 90000002,
  "is_deleted" : false,
  "record_id" : 501,
  "start_date" : "2016-07-26T20:00:00Z"
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

func GetCharactersCharacterIdPortrait(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		characterId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "px128x128" : "https://imageserver.eveonline.com/Character/95465499_128.jpg",
  "px256x256" : "https://imageserver.eveonline.com/Character/95465499_256.jpg",
  "px512x512" : "https://imageserver.eveonline.com/Character/95465499_512.jpg",
  "px64x64" : "https://imageserver.eveonline.com/Character/95465499_64.jpg"
}`
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

func GetCharactersNames(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		characterIds []int64
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "character_id" : 95465499,
  "character_name" : "CCP Bartender"
} ]`
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	localV, err = processParameters(characterIds, r.Form.Get("character_ids"))
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

func PostCharactersCharacterIdCspa(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		characterId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "cost" : 295000
}`
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


