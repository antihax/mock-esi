package swaggerServer

import (
	"net/http"
	"github.com/gorilla/mux"
)

var _ = mux.NewRouter

func GetCharactersCharacterIdWallets(w http.ResponseWriter, r *http.Request) {

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
  "balance" : 295000,
  "wallet_id" : 1000
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(characterId, vars["characterId"])
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}


