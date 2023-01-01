package esiv2

import "time"

/*
200 ok object
*/
type GetCharactersCharacterIdClonesOk struct {
	/* */
	HomeLocation GetCharactersCharacterIdClonesHomeLocation `json:"home_location,omitempty"`
	/*
	 jump_clones array */
	JumpClones []GetCharactersCharacterIdClonesJumpClone `json:"jump_clones,omitempty"`
	/*
	 last_jump_date string */
	LastJumpDate time.Time `json:"last_jump_date,omitempty"`
}
