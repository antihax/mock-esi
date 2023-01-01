package esidev

/*
200 ok object
*/
type GetFwWars200Ok struct {
	/*
	 The faction ID of the enemy faction. */
	AgainstId int32 `json:"against_id,omitempty"`
	/*
	 faction_id integer */
	FactionId int32 `json:"faction_id,omitempty"`
}
