package swaggerServer

import "time"
var _ time.Time

/* 
Is not a character ID */
type GetCharactersCharacterIdUnprocessableEntity struct {
/*
	 error message */
	_error string `json:"error,omitempty"`
}
