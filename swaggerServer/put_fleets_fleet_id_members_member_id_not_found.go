package swaggerServer

import "time"
var _ time.Time

/* 
404 not found object */
type PutFleetsFleetIdMembersMemberIdNotFound struct {
/*
	 Error message */
	_error string `json:"error,omitempty"`
}
