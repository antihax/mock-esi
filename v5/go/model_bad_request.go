package esiv5

/*
Bad request model
*/
type BadRequest struct {
	/*
	 Bad request message */
	Error_ string `json:"error,omitempty"`
}
