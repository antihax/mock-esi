package esilegacy

import "time"

/*
200 ok object
*/
type GetCorporationsCorporationIdStarbases200Ok struct {
	/*
	 The moon this starbase (POS) is anchored on, unanchored POSes do not have this information */
	MoonId int32 `json:"moon_id,omitempty"`
	/*
	 When the POS onlined, for starbases (POSes) in online state */
	OnlinedSince time.Time `json:"onlined_since,omitempty"`
	/*
	 When the POS will be out of reinforcement, for starbases (POSes) in reinforced state */
	ReinforcedUntil time.Time `json:"reinforced_until,omitempty"`
	/*
	 Unique ID for this starbase (POS) */
	StarbaseId int64 `json:"starbase_id,omitempty"`
	/*
	 state string */
	State string `json:"state,omitempty"`
	/*
	 The solar system this starbase (POS) is in, unanchored POSes have this information */
	SystemId int32 `json:"system_id,omitempty"`
	/*
	 Starbase (POS) type */
	TypeId int32 `json:"type_id,omitempty"`
	/*
	 When the POS started unanchoring, for starbases (POSes) in unanchoring state */
	UnanchorAt time.Time `json:"unanchor_at,omitempty"`
}
