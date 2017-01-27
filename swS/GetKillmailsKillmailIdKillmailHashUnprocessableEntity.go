package swS

import "time"
var _ time.Time

/* 
killmail_id and/or killmail_hash is not valid */
type GetKillmailsKillmailIdKillmailHashUnprocessableEntity struct {
/*
	 error message */
	_error string `json:"error,omitempty"`
}
