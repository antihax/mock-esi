package esilogin

import (
	"net/http"
)

func Token(w http.ResponseWriter, r *http.Request) {

	j := `{
		"token_type":"bearer",
		"access_token":"VERYFAKETOKEN",
		"expires_in":60,
		"refresh_token":"THISISNOTAREALTOKEN"
	}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}
