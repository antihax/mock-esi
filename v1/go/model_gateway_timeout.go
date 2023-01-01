package esiv1

/*
Gateway timeout model
*/
type GatewayTimeout struct {
	/*
	 Gateway timeout message */
	Error_ string `json:"error,omitempty"`
	/*
	 number of seconds the request was given */
	Timeout int32 `json:"timeout,omitempty"`
}
