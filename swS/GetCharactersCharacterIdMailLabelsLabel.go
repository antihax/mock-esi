package swS

import "time"
var _ time.Time

/* 
label object */
type GetCharactersCharacterIdMailLabelsLabel struct {
/*
	 color string */
	color string `json:"color,omitempty"`
/*
	 label_id integer */
	label_id int32 `json:"label_id,omitempty"`
/*
	 name string */
	name string `json:"name,omitempty"`
/*
	 unread_count integer */
	unread_count int32 `json:"unread_count,omitempty"`
}
