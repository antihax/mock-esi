package swServer

import "time"
var _ time.Time

/* 
Alliance not found */
type GetAlliancesAllianceIdNotFound struct {
/*
	 error message */
	_error string `json:"error,omitempty"`
}
