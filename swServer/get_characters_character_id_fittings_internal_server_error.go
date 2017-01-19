package swServer

import "time"
var _ time.Time

/* 
Internal server error */
type GetCharactersCharacterIdFittingsInternalServerError struct {
/*
	 Internal server error message */
	_error string `json:"error,omitempty"`
}
