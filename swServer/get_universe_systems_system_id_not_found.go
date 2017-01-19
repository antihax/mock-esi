package swServer

import "time"
var _ time.Time

/* 
System not found */
type GetUniverseSystemsSystemIdNotFound struct {
/*
	 error message */
	_error string `json:"error,omitempty"`
}
