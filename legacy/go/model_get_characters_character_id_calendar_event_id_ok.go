package esilegacy



/* 
Full details of a specific event */
type GetCharactersCharacterIdCalendarEventIdOk struct {
/*
	 duration_in_minutes integer */
	DurationInMinutes int32 `json:"duration_in_minutes,omitempty"`
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
	 event_text string */
	EventText string `json:"event_text,omitempty"`
/*
	 importance integer */
	Importance int32 `json:"importance,omitempty"`
/*
	 owner_id integer */
	OwnerId int32 `json:"owner_id,omitempty"`
/*
	 owner_name string */
	OwnerName string `json:"owner_name,omitempty"`
/*
	 owner_type_id integer */
	OwnerTypeId int32 `json:"owner_type_id,omitempty"`
/*
	 title string */
	Title string `json:"title,omitempty"`
}
