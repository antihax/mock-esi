package esiv5

/*
Forbidden model */
type Forbidden struct {
	/*
	 Forbidden message */
	Error_ string `json:"error,omitempty"`
	/*
	 status code received from SSO */
	SsoStatus int32 `json:"sso_status,omitempty"`
}
