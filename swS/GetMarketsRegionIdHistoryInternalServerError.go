package swS

import "time"
var _ time.Time

/* 
Internal server error */
type GetMarketsRegionIdHistoryInternalServerError struct {
/*
	 Internal server error message */
	_error string `json:"error,omitempty"`
}
