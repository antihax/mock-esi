package swServer

import "time"
var _ time.Time

/* 
response schema */
type PutCharactersCharacterIdCalendarEventIdResponse struct {
/*
	 response string */
	response string `json:"response,omitempty"`
}
