package swaggerServer

import (
	"net/http"

)

func GetInsurancePrices(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "levels" : [ {
    "cost" : 10.0,
    "name" : "Basic",
    "payout" : 20.0
  } ],
  "type_id" : 1
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


