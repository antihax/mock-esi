package swServer

import "time"
var _ time.Time

/* 
Full details of a specific event */
type GetCharactersCharacterIdCalendarEventIdOk struct {
/*
	 date string */
	date time.Time `json:"date,omitempty"`
/*
	 Length in minutes */
	duration int32 `json:"duration,omitempty"`
/*
	 event_id integer */
	event_id int32 `json:"event_id,omitempty"`
/*
	 importance integer */
	importance int32 `json:"importance,omitempty"`
/*
	 owner_id integer */
	owner_id int32 `json:"owner_id,omitempty"`
/*
	 owner_name string */
	owner_name string `json:"owner_name,omitempty"`
/*
	 owner_type string */
	owner_type string `json:"owner_type,omitempty"`
/*
	 response string */
	response string `json:"response,omitempty"`
/*
	 text string */
	text string `json:"text,omitempty"`
/*
	 title string */
	title string `json:"title,omitempty"`
}
