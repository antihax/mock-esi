package swaggerServer

import (
	"net/http"

)

func GetMarketsPrices(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "adjusted_price" : 306988.09,
  "average_price" : 306292.67,
  "type_id" : 32772
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetMarketsRegionIdHistory(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "average" : 5.25,
  "date" : "2015-05-01",
  "highest" : 5.27,
  "lowest" : 5.11,
  "order_count" : 2267,
  "volume" : 16276782035
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetMarketsRegionIdOrders(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "duration" : 90,
  "is_buy_order" : false,
  "issued" : "2016-09-03T05:12:25Z",
  "location_id" : 60005599,
  "min_volume" : 1,
  "order_id" : 4623824223,
  "price" : 9.9,
  "range" : "region",
  "type_id" : 34,
  "volume_remain" : 1296000,
  "volume_total" : 2000000
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


