package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdClonesOk struct {
/* */
	home_location GetCharactersCharacterIdClonesHomeLocation `json:"home_location,omitempty"`
/*
	 jump_clones array */
	jump_clones []GetCharactersCharacterIdClonesJumpClone `json:"jump_clones,omitempty"`
/*
	 last_jump_date string */
	last_jump_date time.Time `json:"last_jump_date,omitempty"`
}
