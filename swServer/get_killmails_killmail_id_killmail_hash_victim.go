package swServer

import "time"
var _ time.Time

/* 
victim object */
type GetKillmailsKillmailIdKillmailHashVictim struct {
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
	 How much total damage was taken by the victim
 */
	damage_taken int32 `json:"damage_taken,omitempty"`
/*
	 faction_id integer */
	faction_id int32 `json:"faction_id,omitempty"`
/*
	 items array */
	items []GetKillmailsKillmailIdKillmailHashItem1 `json:"items,omitempty"`
/* */
	position GetKillmailsKillmailIdKillmailHashPosition `json:"position,omitempty"`
/*
	 The ship that the victim was piloting and was destroyed
 */
	ship_type_id int32 `json:"ship_type_id,omitempty"`
}
