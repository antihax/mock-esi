package swServer

import "time"
var _ time.Time

/* 
bad region_id */
type GetMarketsRegionIdOrdersUnprocessableEntity struct {
/*
	 error message */
	_error string `json:"error,omitempty"`
}
