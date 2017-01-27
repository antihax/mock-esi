package swS

import "time"
var _ time.Time

/* 
Bad request */
type PostCharactersCharacterIdMailBadRequest struct {
/*
	 Bad request message */
	_error string `json:"error,omitempty"`
}
