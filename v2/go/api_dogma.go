package esiv2

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var _ time.Time
var _ = mux.NewRouter

func GetDogmaEffectsEffectId(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		effectId   int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "description" : "Requires a high power slot.",
  "display_name" : "High power",
  "effect_category" : 0,
  "effect_id" : 12,
  "icon_id" : 293,
  "name" : "hiPower",
  "post_expression" : 131,
  "pre_expression" : 131,
  "published" : true
}`
	vars := mux.Vars(r)
	localV, err = processParameters(effectId, vars["effect_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	effectId = localV.(int32)
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
