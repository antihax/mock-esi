package esilatest

import "time"

/*
200 ok object
*/
type GetCharactersCharacterIdFatigueOk struct {
	/*
	 Character's jump fatigue expiry */
	JumpFatigueExpireDate time.Time `json:"jump_fatigue_expire_date,omitempty"`
	/*
	 Character's last jump activation */
	LastJumpDate time.Time `json:"last_jump_date,omitempty"`
	/*
	 Character's last jump update */
	LastUpdateDate time.Time `json:"last_update_date,omitempty"`
}
