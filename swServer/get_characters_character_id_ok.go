package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdOk struct {
/*
	 ancestry_id integer */
	ancestry_id int32 `json:"ancestry_id,omitempty"`
/*
	 Creation date of the character */
	birthday time.Time `json:"birthday,omitempty"`
/*
	 bloodline_id integer */
	bloodline_id int32 `json:"bloodline_id,omitempty"`
/*
	 The character's corporation ID */
	corporation_id int32 `json:"corporation_id,omitempty"`
/*
	 description string */
	description string `json:"description,omitempty"`
/*
	 gender string */
	gender string `json:"gender,omitempty"`
/*
	 The name of the character */
	name string `json:"name,omitempty"`
/*
	 race_id integer */
	race_id int32 `json:"race_id,omitempty"`
/*
	 security_status number */
	security_status float32 `json:"security_status,omitempty"`
}
