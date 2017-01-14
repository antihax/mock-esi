package swaggerServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdMailMailIdOk struct {
/*
	 Mail's body */
	body string `json:"body,omitempty"`
/*
	 From whom the mail was sent */
	from int32 `json:"from,omitempty"`
/*
	 Labels attached to the mail */
	labels []int64 `json:"labels,omitempty"`
/*
	 Whether the mail is flagged as read */
	read bool `json:"read,omitempty"`
/*
	 Recipients of the mail */
	recipients []GetCharactersCharacterIdMailMailIdRecipient `json:"recipients,omitempty"`
/*
	 Mail subject */
	subject string `json:"subject,omitempty"`
/*
	 When the mail was sent */
	timestamp time.Time `json:"timestamp,omitempty"`
}
