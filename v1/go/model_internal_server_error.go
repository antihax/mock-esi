package esiv1

/*
Internal server error model */
type InternalServerError struct {
	/*
	 Internal server error message */
	Error_ string `json:"error,omitempty"`
}
