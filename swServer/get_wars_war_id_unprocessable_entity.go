package swServer

import "time"
var _ time.Time

/* 
war_id is not valid */
type GetWarsWarIdUnprocessableEntity struct {
/*
	 error message */
	_error string `json:"error,omitempty"`
}
