package swServer

import "time"
var _ time.Time

/* 
Forbidden */
type GetCharactersCharacterIdMailForbidden struct {
/*
	 Forbidden message */
	_error string `json:"error,omitempty"`
}
