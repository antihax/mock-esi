package swaggerServer

import (
	"net/http"

)

func GetIndustryFacilities(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "facility_id" : 60012544,
  "owner_id" : 1000126,
  "region_id" : 10000001,
  "solar_system_id" : 30000032,
  "tax" : 0.1,
  "type_id" : 2502
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetIndustrySystems(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "cost_indices" : [ {
    "activity" : "invention",
    "cost_index" : 0.00480411064973412
  } ],
  "solar_system_id" : 30011392
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


