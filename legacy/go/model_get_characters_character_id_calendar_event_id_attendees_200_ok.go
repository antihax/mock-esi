package esilegacy

/*
character_id and response of an attendee
*/
type GetCharactersCharacterIdCalendarEventIdAttendees200Ok struct {
	/*
	 character_id integer */
	CharacterId int32 `json:"character_id,omitempty"`
	/*
	 event_response string */
	EventResponse string `json:"event_response,omitempty"`
}
