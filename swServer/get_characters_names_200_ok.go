package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersNames200Ok struct {
/*
	 character_id integer */
	character_id int64 `json:"character_id,omitempty"`
/*
	 character_name string */
	character_name string `json:"character_name,omitempty"`
}
