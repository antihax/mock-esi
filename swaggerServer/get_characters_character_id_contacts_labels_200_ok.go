package swaggerServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdContactsLabels200Ok struct {
/*
	 label_id integer */
	label_id int64 `json:"label_id,omitempty"`
/*
	 label_name string */
	label_name string `json:"label_name,omitempty"`
}
