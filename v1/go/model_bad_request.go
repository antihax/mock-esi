package esiv1

/*
Bad request model
*/
type BadRequest struct {
	/*
	 Bad request message */
	Error_ string `json:"error,omitempty"`
}
