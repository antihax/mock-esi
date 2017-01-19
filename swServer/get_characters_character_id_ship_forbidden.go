package swServer

import "time"
var _ time.Time

/* 
Forbidden */
type GetCharactersCharacterIdShipForbidden struct {
/*
	 Forbidden message */
	_error string `json:"error,omitempty"`
}
