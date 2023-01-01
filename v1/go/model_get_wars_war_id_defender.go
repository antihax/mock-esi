package esiv1

/*
The defending corporation or alliance that declared this war, only contains either corporation_id or alliance_id
*/
type GetWarsWarIdDefender struct {
	/*
	 Alliance ID if and only if the defender is an alliance */
	AllianceId int32 `json:"alliance_id,omitempty"`
	/*
	 Corporation ID if and only if the defender is a corporation */
	CorporationId int32 `json:"corporation_id,omitempty"`
	/*
	 ISK value of ships the defender has killed */
	IskDestroyed float32 `json:"isk_destroyed,omitempty"`
	/*
	 The number of ships the defender has killed */
	ShipsKilled int32 `json:"ships_killed,omitempty"`
}
