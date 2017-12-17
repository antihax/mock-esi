package esilegacy

import "time"

/*
200 ok object */
type GetCharactersCharacterIdMailMailIdOk struct {
	/*
	 Mail subject */
	Subject string `json:"subject,omitempty"`
	/*
	 From whom the mail was sent */
	From int32 `json:"from,omitempty"`
	/*
	 When the mail was sent */
	Timestamp time.Time `json:"timestamp,omitempty"`
	/*
	 Recipients of the mail */
	Recipients []GetCharactersCharacterIdMailMailIdRecipient `json:"recipients,omitempty"`
	/*
	 Mail's body */
	Body string `json:"body,omitempty"`
	/*
	 Labels attached to the mail */
	Labels []int64 `json:"labels,omitempty"`
	/*
	 Whether the mail is flagged as read */
	Read bool `json:"read,omitempty"`
}
