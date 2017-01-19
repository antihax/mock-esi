package swServer

import "time"
var _ time.Time

/* 
404 not found object */
type GetFleetsFleetIdNotFound struct {
/*
	 error message */
	_error string `json:"error,omitempty"`
}
