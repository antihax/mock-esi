package esilatest

import "time"

/*
200 ok object */
type GetCharactersCharacterIdClonesOk struct {
	/*
	 last_clone_jump_date string */
	LastCloneJumpDate time.Time `json:"last_clone_jump_date,omitempty"`
	/* */
	HomeLocation GetCharactersCharacterIdClonesHomeLocation `json:"home_location,omitempty"`
	/*
	 last_station_change_date string */
	LastStationChangeDate time.Time `json:"last_station_change_date,omitempty"`
	/*
	 jump_clones array */
	JumpClones []GetCharactersCharacterIdClonesJumpClone `json:"jump_clones,omitempty"`
}
