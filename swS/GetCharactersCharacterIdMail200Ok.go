package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdMail200Ok struct {
/*
	 From whom the mail was sent */
	from int32 `json:"from,omitempty"`
/*
	 is_read boolean */
	is_read bool `json:"is_read,omitempty"`
/*
	 labels array */
	labels []int64 `json:"labels,omitempty"`
/*
	 mail_id integer */
	mail_id int64 `json:"mail_id,omitempty"`
/*
	 Recipients of the mail */
	recipients []GetCharactersCharacterIdMailRecipient `json:"recipients,omitempty"`
/*
	 Mail subject */
	subject string `json:"subject,omitempty"`
/*
	 When the mail was sent */
	timestamp time.Time `json:"timestamp,omitempty"`
}
