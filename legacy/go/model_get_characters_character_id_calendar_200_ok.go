package esilegacy

import "time"

/*
event
*/
type GetCharactersCharacterIdCalendar200Ok struct {
	/*
	 event_date string */
	EventDate time.Time `json:"event_date,omitempty"`
	/*
	 event_id integer */
	EventId int32 `json:"event_id,omitempty"`
	/*
	 event_response string */
	EventResponse string `json:"event_response,omitempty"`
	/*
	 importance integer */
	Importance int32 `json:"importance,omitempty"`
	/*
	 title string */
	Title string `json:"title,omitempty"`
}
