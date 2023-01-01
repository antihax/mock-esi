package esiv1

/*
200 ok object
*/
type GetCorporationCorporationIdMiningObserversObserverId200Ok struct {
	/*
	 The character that did the mining  */
	CharacterId int32 `json:"character_id,omitempty"`
	/*
	 last_updated string */
	LastUpdated string `json:"last_updated,omitempty"`
	/*
	 quantity integer */
	Quantity int64 `json:"quantity,omitempty"`
	/*
	 The corporation id of the character at the time data was recorded.  */
	RecordedCorporationId int32 `json:"recorded_corporation_id,omitempty"`
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
}
