package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdContacts200Ok struct {
/*
	 contact_id integer */
	contact_id int32 `json:"contact_id,omitempty"`
/*
	 contact_type string */
	contact_type string `json:"contact_type,omitempty"`
/*
	 Whether this contact is in the blocked list. Note a missing value denotes unknown, not true or false */
	is_blocked bool `json:"is_blocked,omitempty"`
/*
	 Whether this contact is being watched */
	is_watched bool `json:"is_watched,omitempty"`
/*
	 Custom label of the contact */
	label_id int64 `json:"label_id,omitempty"`
/*
	 Standing of the contact */
	standing float32 `json:"standing,omitempty"`
}
