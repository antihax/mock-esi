package swS

import "time"
var _ time.Time

/* 
Internal server error */
type GetUniverseRacesInternalServerError struct {
/*
	 Internal server error message */
	_error string `json:"error,omitempty"`
}
