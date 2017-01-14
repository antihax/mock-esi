package swaggerServer

import "time"
var _ time.Time

/* 
recipient object */
type PostCharactersCharacterIdMailRecipient struct {
/*
	 recipient_id integer */
	recipient_id int32 `json:"recipient_id,omitempty"`
/*
	 recipient_type string */
	recipient_type string `json:"recipient_type,omitempty"`
}
