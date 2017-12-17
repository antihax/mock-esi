package esilatest

import "time"

/*
Full details of a specific event */
type GetCharactersCharacterIdCalendarEventIdOk struct {
	/*
	 event_id integer */
	EventId int32 `json:"event_id,omitempty"`
	/*
	 owner_id integer */
	OwnerId int32 `json:"owner_id,omitempty"`
	/*
	 owner_name string */
	OwnerName string `json:"owner_name,omitempty"`
	/*
	 date string */
	Date time.Time `json:"date,omitempty"`
	/*
	 title string */
	Title string `json:"title,omitempty"`
	/*
	 Length in minutes */
	Duration int32 `json:"duration,omitempty"`
	/*
	 importance integer */
	Importance int32 `json:"importance,omitempty"`
	/*
	 response string */
	Response string `json:"response,omitempty"`
	/*
	 text string */
	Text string `json:"text,omitempty"`
	/*
	 owner_type string */
	OwnerType string `json:"owner_type,omitempty"`
}
