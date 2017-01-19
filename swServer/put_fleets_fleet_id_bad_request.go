package swServer

import "time"
var _ time.Time

/* 
Bad request */
type PutFleetsFleetIdBadRequest struct {
/*
	 Bad request message */
	_error string `json:"error,omitempty"`
}
