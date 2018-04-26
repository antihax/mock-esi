package esilatest

import "time"

/*
200 ok object */
type GetCharactersCharacterIdMail200Ok struct {
	/*
	 From whom the mail was sent */
	From int32 `json:"from,omitempty"`
	/*
	 is_read boolean */
	IsRead bool `json:"is_read,omitempty"`
	/*
	 labels array */
	Labels []int32 `json:"labels,omitempty"`
	/*
	 mail_id integer */
	MailId int32 `json:"mail_id,omitempty"`
	/*
	 Recipients of the mail */
	Recipients []GetCharactersCharacterIdMailRecipient `json:"recipients,omitempty"`
	/*
	 Mail subject */
	Subject string `json:"subject,omitempty"`
	/*
	 When the mail was sent */
	Timestamp time.Time `json:"timestamp,omitempty"`
}
