package esilegacy

/*
recipient object
*/
type GetCharactersCharacterIdMailMailIdRecipient struct {
	/*
	 recipient_id integer */
	RecipientId int32 `json:"recipient_id,omitempty"`
	/*
	 recipient_type string */
	RecipientType string `json:"recipient_type,omitempty"`
}
