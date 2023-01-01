package esidev

import "time"

/*
200 ok object
*/
type GetCharactersCharacterIdMailMailIdOk struct {
	/*
	 Mail's body */
	Body string `json:"body,omitempty"`
	/*
	 From whom the mail was sent */
	From int32 `json:"from,omitempty"`
	/*
	 Labels attached to the mail */
	Labels []int32 `json:"labels,omitempty"`
	/*
	 Whether the mail is flagged as read */
	Read bool `json:"read,omitempty"`
	/*
	 Recipients of the mail */
	Recipients []GetCharactersCharacterIdMailMailIdRecipient `json:"recipients,omitempty"`
	/*
	 Mail subject */
	Subject string `json:"subject,omitempty"`
	/*
	 When the mail was sent */
	Timestamp time.Time `json:"timestamp,omitempty"`
}
