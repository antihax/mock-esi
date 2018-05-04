package esiv3



/* 
Full details of a specific event */
type GetCharactersCharacterIdCalendarEventIdOk struct {
/*
	 date string */
	Date time.Time `json:"date,omitempty"`
/*
	 Length in minutes */
	Duration int32 `json:"duration,omitempty"`
/*
	 event_id integer */
	EventId int32 `json:"event_id,omitempty"`
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
	 owner_type string */
	OwnerType string `json:"owner_type,omitempty"`
/*
	 response string */
	Response string `json:"response,omitempty"`
/*
	 text string */
	Text string `json:"text,omitempty"`
/*
	 title string */
	Title string `json:"title,omitempty"`
}
