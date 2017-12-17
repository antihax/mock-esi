package esilatest

/*
Forbidden model */
type Forbidden struct {
	/*
	 Forbidden message */
	Error_ string `json:"error,omitempty"`
	/*
	 Status code received from SSO */
	SsoStatus int32 `json:"sso_status,omitempty"`
}
