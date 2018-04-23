package esidev

import "time"

/*
200 ok object */
type GetCharactersCharacterIdMail200Ok struct {
	/*
	 mail_id integer */
	MailId int32 `json:"mail_id,omitempty"`
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
	 labels array */
	Labels []int32 `json:"labels,omitempty"`
	/*
	 Recipients of the mail */
	Recipients []GetCharactersCharacterIdMailRecipient `json:"recipients,omitempty"`
	/*
	 is_read boolean */
	IsRead bool `json:"is_read,omitempty"`
}
