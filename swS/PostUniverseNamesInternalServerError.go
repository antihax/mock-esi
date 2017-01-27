package swS

import "time"
var _ time.Time

/* 
Internal server error */
type PostUniverseNamesInternalServerError struct {
/*
	 Internal server error message */
	_error string `json:"error,omitempty"`
}
