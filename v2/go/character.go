package esiV2

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var _ time.Time
var _ = mux.NewRouter

func GetCharactersCharacterIdPortrait(w http.ResponseWriter, r *http.Request) {

	var (
		localV      interface{}
		err         error
		characterId int32
		datasource  string
		userAgent   string
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
