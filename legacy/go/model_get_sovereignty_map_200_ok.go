package esilegacy

/*
200 ok object
*/
type GetSovereigntyMap200Ok struct {
	/*
	 alliance_id integer */
	AllianceId int32 `json:"alliance_id,omitempty"`
	/*
	 corporation_id integer */
	CorporationId int32 `json:"corporation_id,omitempty"`
	/*
	 faction_id integer */
	FactionId int32 `json:"faction_id,omitempty"`
	/*
	 system_id integer */
	SystemId int32 `json:"system_id,omitempty"`
}
