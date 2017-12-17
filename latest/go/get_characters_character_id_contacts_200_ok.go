package esilatest

/*
200 ok object */
type GetCharactersCharacterIdContacts200Ok struct {
	/*
	 Standing of the contact */
	Standing float32 `json:"standing,omitempty"`
	/*
	 contact_type string */
	ContactType string `json:"contact_type,omitempty"`
	/*
	 contact_id integer */
	ContactId int32 `json:"contact_id,omitempty"`
	/*
	 Whether this contact is being watched */
	IsWatched bool `json:"is_watched,omitempty"`
	/*
	 Whether this contact is in the blocked list. Note a missing value denotes unknown, not true or false */
	IsBlocked bool `json:"is_blocked,omitempty"`
	/*
	 Custom label of the contact */
	LabelId int64 `json:"label_id,omitempty"`
}
