package esilogin

import (
	"net/http"
)

func Token(w http.ResponseWriter, r *http.Request) {

	j := `{
		"token_type":"bearer",
		"access_token":"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyIjoiVlx1MDAxNcKbwoNUwoonbFPCu8KhwrYiLCJpYXQiOjE0NDQyNjI1NDMsImV4cCI6MTQ0NDI2MjU2M30.MldruS1PvZaRZIJR4legQaauQ3_DYKxxP2rFnD37Ip4",
		"expires_in":60,
		"refresh_token":"fdb8fdbecf1d03ce5e6125c067733c0d51de209c"
	}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}
