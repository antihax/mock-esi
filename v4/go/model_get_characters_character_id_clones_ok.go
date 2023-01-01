package esiv4

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
	 last_clone_jump_date string */
	LastCloneJumpDate time.Time `json:"last_clone_jump_date,omitempty"`
	/*
	 last_station_change_date string */
	LastStationChangeDate time.Time `json:"last_station_change_date,omitempty"`
}
