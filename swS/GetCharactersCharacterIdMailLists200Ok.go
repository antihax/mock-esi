package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdMailLists200Ok struct {
/*
	 Mailing list ID */
	mailing_list_id int32 `json:"mailing_list_id,omitempty"`
/*
	 name string */
	name string `json:"name,omitempty"`
}
