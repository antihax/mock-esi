package esiv1

import "time"

/*
event */
type GetCharactersCharacterIdCalendar200Ok struct {
	/*
	 event_id integer */
	EventId int32 `json:"event_id,omitempty"`
	/*
	 event_date string */
	EventDate time.Time `json:"event_date,omitempty"`
	/*
	 title string */
	Title string `json:"title,omitempty"`
	/*
	 importance integer */
	Importance int32 `json:"importance,omitempty"`
	/*
	 event_response string */
	EventResponse string `json:"event_response,omitempty"`
}
