package swaggerServer

import (
	"net/http"

)

func GetCharactersCharacterIdWallets(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "balance" : 295000,
  "wallet_id" : 1000
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


