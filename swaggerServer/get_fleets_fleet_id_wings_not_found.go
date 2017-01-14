package swaggerServer

import "time"
var _ time.Time

/* 
404 not found object */
type GetFleetsFleetIdWingsNotFound struct {
/*
	 error message */
	_error string `json:"error,omitempty"`
}
