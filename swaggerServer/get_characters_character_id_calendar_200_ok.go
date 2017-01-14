package swaggerServer

import "time"
var _ time.Time

/* 
event */
type GetCharactersCharacterIdCalendar200Ok struct {
/*
	 event_date string */
	event_date time.Time `json:"event_date,omitempty"`
/*
	 event_id integer */
	event_id int32 `json:"event_id,omitempty"`
/*
	 event_response string */
	event_response string `json:"event_response,omitempty"`
/*
	 importance integer */
	importance int32 `json:"importance,omitempty"`
/*
	 title string */
	title string `json:"title,omitempty"`
}
