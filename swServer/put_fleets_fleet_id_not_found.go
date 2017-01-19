package swServer

import "time"
var _ time.Time

/* 
404 not found object */
type PutFleetsFleetIdNotFound struct {
/*
	 Error message */
	_error string `json:"error,omitempty"`
}
