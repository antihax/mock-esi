package swaggerServer

import "time"
var _ time.Time

/* 
attacker object */
type GetKillmailsKillmailIdKillmailHashAttacker struct {
/*
	 alliance_id integer */
	alliance_id int32 `json:"alliance_id,omitempty"`
/*
	 character_id integer */
	character_id int32 `json:"character_id,omitempty"`
/*
	 corporation_id integer */
	corporation_id int32 `json:"corporation_id,omitempty"`
/*
	 damage_done integer */
	damage_done int32 `json:"damage_done,omitempty"`
/*
	 faction_id integer */
	faction_id int32 `json:"faction_id,omitempty"`
/*
	 Was the attacker the one to achieve the final blow
 */
	final_blow bool `json:"final_blow,omitempty"`
/*
	 Security status for the attacker
 */
	security_status float32 `json:"security_status,omitempty"`
/*
	 What ship was the attacker flying
 */
	ship_type_id int32 `json:"ship_type_id,omitempty"`
/*
	 What weapon was used by the attacker for the kill
 */
	weapon_type_id int32 `json:"weapon_type_id,omitempty"`
}
